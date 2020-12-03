package plugin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wpgitupdater/wpgitupdater/internal/config"
	"github.com/wpgitupdater/wpgitupdater/internal/git"
	"github.com/wpgitupdater/wpgitupdater/internal/github"
	"github.com/wpgitupdater/wpgitupdater/internal/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type PluginInfo struct {
	Version     string `json:"version"`
	Download    string `json:"download_link"`
	LastUpdated string `json:"last_updated"`
	Homepage    string `json:"homepage"`
	Sections    struct {
		Changelog string `json:"changelog"`
	} `json:"sections"`
}

type Plugin struct {
	Slug    string
	Path    string
	Name    string
	Version string
	Info    PluginInfo
}

func GetPlugins(cnf *config.Config) map[string]Plugin {
	plugins := map[string]Plugin{}

	fmt.Println("Collecting plugin information")
	matches, _ := filepath.Glob(cnf.GetPluginsPath("/**/*.php"))
	for _, file := range matches {
		path := filepath.Dir(file)
		slug := filepath.Base(path)

		if !cnf.PluginCanBeUpdated(slug) {
			continue
		}

		_, exists := plugins[slug]
		if exists {
			continue
		}

		name, version, err := GetPluginInfo(file)
		if err != nil {
			continue
		}

		fmt.Println(fmt.Sprintf("[%s] plugin found", slug))
		plugin := Plugin{Slug: slug, Path: path, Name: name, Version: version}

		fmt.Println(fmt.Sprintf("[%s] loading external plugin info", slug))
		plugin.LoadExternalInfo()

		plugins[slug] = plugin
	}

	return plugins
}

func GetPluginInfo(file string) (string, string, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", "", err
	}

	text := string(content)

	nameR, _ := regexp.Compile("[ \t/*#@]*Plugin Name:(?P<Name>.*)")
	name := nameR.FindStringSubmatch(text)
	if len(name) < 1 {
		return "", "", errors.New("")
	}

	versionR, _ := regexp.Compile("[ \t/*#@]*Version:(?P<Version>.*)")
	version := versionR.FindStringSubmatch(text)
	if len(version) < 1 {
		return "", "", errors.New("")
	}

	return strings.TrimSpace(name[1]), strings.TrimSpace(version[1]), nil
}

func ListPlugins(cnf *config.Config) {
	plugins := GetPlugins(cnf)
	for slug, plugin := range plugins {
		status := ""
		if plugin.HasPendingUpdate() {
			status = "outdated"
		} else {
			status = "uptodate"
		}
		fmt.Printf("%-60v[%v]\n", slug, status)
	}
}

func UpdatePlugins(cnf *config.Config, dryRun bool) {
	plugins := GetPlugins(cnf)
	git.ConfigureGitConfig(cnf)
	for _, plugin := range plugins {
		plugin.PerformPluginUpdate(cnf, dryRun)
	}
	git.RestoreGitConfig(cnf)
}

func (plugin *Plugin) LoadExternalInfo() {
	info := PluginInfo{}
	resp, err := http.Get("https://api.wordpress.org/plugins/info/1.2/?action=plugin_information&request[slug]=" + plugin.Slug)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&info)
	if err != nil {
		log.Fatal(err)
	}

	plugin.Info = info
}

func (plugin Plugin) HasPendingUpdate() bool {
	if plugin.Info.Version == "" {
		return false
	}
	return utils.VersionCompare(plugin.Version, plugin.Info.Version, "<")
}

func (plugin Plugin) GetBranchName() string {
	return "wpgitupdates-plugin-" + plugin.Slug + "-" + plugin.Version + "-" + plugin.Info.Version
}

func (plugin Plugin) GetCommitMessage(cnf *config.Config) string {
	msg := strings.ReplaceAll(cnf.GetPluginsCommit(), ":plugin", plugin.Slug)
	msg = strings.ReplaceAll(msg, ":oldversion", plugin.Version)
	msg = strings.ReplaceAll(msg, ":newversion", plugin.Info.Version)
	return msg
}

func (plugin Plugin) GetPRTitle(cnf *config.Config) string {
	msg := strings.ReplaceAll(cnf.GetPluginsPRTitle(), ":plugin", plugin.Slug)
	msg = strings.ReplaceAll(msg, ":oldversion", plugin.Version)
	msg = strings.ReplaceAll(msg, ":newversion", plugin.Info.Version)
	return msg
}

func (plugin Plugin) GetHomePage() string {
	return plugin.Info.Homepage
}

func (plugin Plugin) GetLastUpdated() string {
	return plugin.Info.LastUpdated
}

func (plugin Plugin) GetChangelog() string {
	return plugin.Info.Sections.Changelog
}

func (plugin Plugin) UpdateBranchExists() bool {
	return git.BranchExists(plugin.GetBranchName())
}

func (plugin Plugin) PerformPluginUpdate(cnf *config.Config, dryRun bool) {
	if !plugin.HasPendingUpdate() {
		fmt.Printf("[%s] Already up to date, skipping\n", plugin.Slug)
		return
	}

	if plugin.UpdateBranchExists() {
		fmt.Printf("[%s] Update branch exists, skipping\n", plugin.Slug)
		return
	}

	if dryRun {
		fmt.Printf("[%s] Skipping actual update process...\n", plugin.Slug)
		return
	}

	branchName := plugin.GetBranchName()
	downloadPath := cnf.GetPluginsPath(filepath.Base(plugin.Info.Download))

	sourceBranch := git.CurrentBranch()

	fmt.Printf("Creating Branch [%v]\n", branchName)
	output := utils.RunCmd("git", "checkout", "-b", branchName)
	fmt.Println(output)

	fmt.Printf("Downloading new plugin version for [%v]\n", plugin.Slug)
	if err := utils.DownloadUrl(plugin.Info.Download, downloadPath); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Removing old plugin version for [%v]\n", plugin.Slug)
	if err := os.RemoveAll(cnf.GetPluginsPath(plugin.Slug)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Extracting new plugin version for [%v]\n", plugin.Slug)
	if _, err := utils.Unzip(downloadPath, cnf.GetPluginsPath("")); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Removing plugin download for [%v]\n", plugin.Slug)
	if err := os.Remove(downloadPath); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Commiting plugin update for [%v]\n", plugin.Slug)
	output = utils.RunCmd("git", "add", "-A", ".")
	fmt.Println(output)

	output = utils.RunCmd("git", "commit", "-a", "-m", plugin.GetCommitMessage(cnf))
	fmt.Println(output)

	fmt.Printf("Pushing plugin update for [%v]\n", plugin.Slug)
	output = utils.RunCmd("git", "push", "-u", "origin", branchName)
	fmt.Println(output)

	fmt.Println("Restoring local branch")
	output = utils.RunCmd("git", "checkout", sourceBranch)
	fmt.Println(output)

	plugin.CreatePullRequest(cnf)
}

func (plugin Plugin) CreatePullRequest(cnf *config.Config) {
	fmt.Println("Creating pull request")
	if err := github.CreatePullRequest(cnf, plugin); err != nil {
		log.Fatal(err)
	}
}