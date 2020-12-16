package theme

import (
	"fmt"
	"github.com/wpgitupdater/wpgitupdater/internal/api"
	"github.com/wpgitupdater/wpgitupdater/internal/config"
	"github.com/wpgitupdater/wpgitupdater/internal/constants"
	"github.com/wpgitupdater/wpgitupdater/internal/git"
	"github.com/wpgitupdater/wpgitupdater/internal/github"
	"github.com/wpgitupdater/wpgitupdater/internal/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ThemeInfo struct {
	Version     string `json:"version"`
	Download    string `json:"download_link"`
	LastUpdated string `json:"last_updated"`
	Homepage    string `json:"homepage"`
	Sections    struct {
		Description string `json:"description"`
	} `json:"sections"`
}

type Theme struct {
	Slug    string
	Path    string
	Name    string
	Version string
	Info    ThemeInfo
}

func GetThemes(cnf *config.Config) map[string]Theme {
	themes := map[string]Theme{}

	fmt.Println("Collecting theme information")
	matches, _ := filepath.Glob(cnf.GetThemesPath("/**/style.css"))
	for _, file := range matches {
		path := filepath.Dir(file)
		slug := filepath.Base(path)

		if !cnf.ThemeCanBeUpdated(slug) {
			continue
		}

		_, exists := themes[slug]
		if exists {
			continue
		}

		name, version, err := utils.GetWordPressHeaderInfo(file, "Theme Name", "Version")
		if err != nil {
			continue
		}

		fmt.Println(fmt.Sprintf("[%s] theme found", slug))
		theme := Theme{Slug: slug, Path: path, Name: name, Version: version, Info: ThemeInfo{}}

		fmt.Println(fmt.Sprintf("[%s] loading external theme info", slug))
		utils.LoadWordPressApiInfo(constants.WordPressThemeApiInfo+theme.Slug, &theme.Info)
		themes[slug] = theme
	}

	return themes
}

func ListThemes(cnf *config.Config) {
	themes := GetThemes(cnf)
	for slug, theme := range themes {
		status := ""
		if theme.HasPendingUpdate() {
			status = "outdated"
		} else {
			status = "uptodate"
		}
		fmt.Printf("%-60v[%v]\n", slug, status)
	}
}

func UpdateThemes(cnf *config.Config, dryRun bool) {
	themes := GetThemes(cnf)
	for _, theme := range themes {
		theme.PerformThemeUpdate(cnf, dryRun)
	}
}

func (theme Theme) HasPendingUpdate() bool {
	if theme.Info.Version == "" {
		return false
	}
	return utils.VersionCompare(theme.Version, theme.Info.Version, "<")
}

func (theme Theme) GetBranchName() string {
	return "wpgitupdates-theme-" + theme.Slug + "-" + theme.Version + "-" + theme.Info.Version
}

func (theme Theme) GetCommitMessage(cnf *config.Config) string {
	msg := strings.ReplaceAll(cnf.GetThemesCommit(), ":theme", theme.Slug)
	msg = strings.ReplaceAll(msg, ":oldversion", theme.Version)
	msg = strings.ReplaceAll(msg, ":newversion", theme.Info.Version)
	return msg
}

func (theme Theme) GetPRTitle(cnf *config.Config) string {
	msg := strings.ReplaceAll(cnf.GetThemesPRTitle(), ":theme", theme.Slug)
	msg = strings.ReplaceAll(msg, ":oldversion", theme.Version)
	msg = strings.ReplaceAll(msg, ":newversion", theme.Info.Version)
	return msg
}

func (theme Theme) GetHomePage() string {
	return theme.Info.Homepage
}

func (theme Theme) GetLastUpdated() string {
	return theme.Info.LastUpdated
}

func (theme Theme) GetChangelog() string {
	return "Changelog information for themes is unavailable, please review the theme homepage for further info."
}

func (theme Theme) UpdateBranchExists() bool {
	return git.BranchExists(theme.GetBranchName())
}

func (theme Theme) PerformThemeUpdate(cnf *config.Config, dryRun bool) {
	if !theme.HasPendingUpdate() {
		fmt.Printf("[%s] Already up to date, skipping\n", theme.Slug)
		return
	}

	if theme.UpdateBranchExists() {
		fmt.Printf("[%s] Update branch exists, skipping\n", theme.Slug)
		return
	}

	if dryRun {
		fmt.Printf("[%s] Skipping actual update process...\n", theme.Slug)
		return
	}

	if err := api.UpdateUsage("theme", theme.Slug); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[%s] Usage updated...\n", theme.Slug)

	branchName := theme.GetBranchName()
	downloadPath := cnf.GetThemesPath(filepath.Base(theme.Info.Download))

	sourceBranch := git.CurrentBranch()

	fmt.Printf("Creating Branch [%v]\n", branchName)
	output := utils.RunCmd("git", "checkout", "-b", branchName)
	fmt.Println(output)

	fmt.Printf("Downloading new theme version for [%v]\n", theme.Slug)
	if err := utils.DownloadUrl(theme.Info.Download, downloadPath); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Removing old theme version for [%v]\n", theme.Slug)
	if err := os.RemoveAll(cnf.GetThemesPath(theme.Slug)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Extracting new theme version for [%v]\n", theme.Slug)
	if _, err := utils.Unzip(downloadPath, cnf.GetThemesPath("")); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Removing theme download for [%v]\n", theme.Slug)
	if err := os.Remove(downloadPath); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Commiting theme update for [%v]\n", theme.Slug)
	output = utils.RunCmd("git", "add", "-A", ".")
	fmt.Println(output)

	output = utils.RunCmd("git", "commit", "-a", "-m", theme.GetCommitMessage(cnf))
	fmt.Println(output)

	fmt.Printf("Pushing theme update for [%v]\n", theme.Slug)
	output = utils.RunCmd("git", "push", "-u", "origin", branchName)
	fmt.Println(output)

	fmt.Println("Restoring local branch")
	output = utils.RunCmd("git", "checkout", sourceBranch)
	fmt.Println(output)

	theme.CreatePullRequest(cnf)
}

func (theme Theme) CreatePullRequest(cnf *config.Config) {
	fmt.Println("Creating pull request")
	if err := github.CreatePullRequest(cnf, theme); err != nil {
		log.Fatal(err)
	}
}
