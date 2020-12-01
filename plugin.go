package main

import (
	"encoding/json"
	"errors"
	"fmt"
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

func GetPlugins(config *Config) map[string]Plugin {
	plugins := map[string]Plugin{}

	fmt.Println("Collecting plugin information")
	matches, _ := filepath.Glob(config.GetPluginsPath("/**/*.php"))
	for _, file := range matches {
		path := filepath.Dir(file)
		slug := filepath.Base(path)

		if !config.PluginCanBeUpdated(slug) {
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

func ListPlugins(config *Config) {
	plugins := GetPlugins(config)
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

func UpdatePlugins(config *Config, dryRun bool) {
	plugins := GetPlugins(config)
	ConfigureGitConfig(config)
	for _, plugin := range plugins {
		plugin.PerformPluginUpdate(config, dryRun)
	}
	RestoreGitConfig(config)
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
	return VersionCompare(plugin.Version, plugin.Info.Version, "<")
}

func (plugin Plugin) GetBranchName() string {
	return "wpgitupdates-plugin-" + plugin.Slug + "-" + plugin.Version + "-" + plugin.Info.Version
}

func (plugin Plugin) GetCommitMessage(config *Config) string {
	msg := strings.ReplaceAll(config.GetPluginsCommit(), ":plugin", plugin.Slug)
	msg = strings.ReplaceAll(msg, ":oldversion", plugin.Version)
	msg = strings.ReplaceAll(msg, ":newversion", plugin.Info.Version)
	return msg
}

func (plugin Plugin) GetPRTitle(config *Config) string {
	msg := strings.ReplaceAll(config.GetPluginsPRTitle(), ":plugin", plugin.Slug)
	msg = strings.ReplaceAll(msg, ":oldversion", plugin.Version)
	msg = strings.ReplaceAll(msg, ":newversion", plugin.Info.Version)
	return msg
}

func (plugin Plugin) UpdateBranchExists() bool {
	return BranchExists(plugin.GetBranchName())
}

func (plugin Plugin) PerformPluginUpdate(config *Config, dryRun bool) {
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
	downloadPath := config.GetPluginsPath(filepath.Base(plugin.Info.Download))

	fmt.Printf("Creating Branch [%v]\n", branchName)
	output := RunCmd("git", "checkout", "-b", branchName)
	fmt.Println(output)

	fmt.Printf("Downloading new plugin version for [%v]\n", plugin.Slug)
	if err := DownloadUrl(plugin.Info.Download, downloadPath); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Removing old plugin version for [%v]\n", plugin.Slug)
	if err := os.RemoveAll(config.GetPluginsPath(plugin.Slug)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Extracting new plugin version for [%v]\n", plugin.Slug)
	if _, err := Unzip(downloadPath, config.GetPluginsPath("")); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Removing plugin download for [%v]\n", plugin.Slug)
	if err := os.Remove(downloadPath); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Commiting plugin update for [%v]\n", plugin.Slug)
	output = RunCmd("git", "add", "-A", ".")
	fmt.Println(output)

	output = RunCmd("git", "commit", "-a", "-m", plugin.GetCommitMessage(config))
	fmt.Println(output)

	fmt.Printf("Pushing plugin update for [%v]\n", plugin.Slug)
	output = RunCmd("git", "push", "-u", "origin", branchName)
	fmt.Println(output)

	fmt.Println("Restoring local branch")
	output = RunCmd("git", "checkout", config.Branch)
	fmt.Println(output)

	plugin.CreatePullRequest(config)
}

func (plugin Plugin) CreatePullRequest(config *Config) {
	fmt.Println("Creating pull request")
	if err := CreatePullRequest(config, plugin); err != nil {
		log.Fatal(err)
	}
}
