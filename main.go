package main

import (
	"flag"
	"fmt"
	"github.com/wpgitupdater/wpgitupdater/internal/config"
	"github.com/wpgitupdater/wpgitupdater/internal/constants"
	"github.com/wpgitupdater/wpgitupdater/internal/git"
	"github.com/wpgitupdater/wpgitupdater/internal/github"
	"github.com/wpgitupdater/wpgitupdater/internal/plugin"
	"github.com/wpgitupdater/wpgitupdater/internal/theme"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Println("WordPress Git Updater v" + constants.Version)
	fmt.Println("Build:", constants.Build)
	fmt.Println("Build Date:", constants.BuildDate)

	var commands map[string]func()
	commands = make(map[string]func())
	commands["init"] = InitCommand()
	commands["list"] = ListCommand()
	commands["update"] = UpdateCommand()

	keys := make([]string, 0, len(commands))
	for k := range commands {
		keys = append(keys, k)
	}
	commandNames := strings.Join(keys, ", ")

	if len(os.Args) < 2 {
		log.Fatal("Expected one of ", commandNames, " command")
	}

	commandName := os.Args[1]
	if _, exists := commands[commandName]; !exists {
		log.Fatal("Expected one of ", commandNames, " command")
	}

	commands[commandName]()
}

func InitCommand() func() {
	return func() {
		cmd := flag.NewFlagSet("init", flag.ExitOnError)
		var ci bool
		cmd.BoolVar(&ci, "ci", false, "Create a CI config file")
		var actions bool
		cmd.BoolVar(&actions, "actions", false, "Create Github Actions workflow file")
		cmd.Parse(os.Args[2:])

		if ci && actions {
			fmt.Println("Creating workflow file")
			github.CreateWorkflowTemplate()
			fmt.Println("Workflow file created!")
		} else {
			fmt.Println("Creating config file")
			config.CreateConfigTemplate()
			fmt.Println("Config file created!")
		}
	}
}

func ListCommand() func() {
	return func() {
		cmd := flag.NewFlagSet("list", flag.ExitOnError)
		var plugins bool
		cmd.BoolVar(&plugins, "plugins", true, "List plugin updates")
		var themes bool
		cmd.BoolVar(&themes, "themes", true, "List theme updates")
		cmd.Parse(os.Args[2:])
		fmt.Println("List update statuses")

		cnf := config.LoadConfig()
		if plugins {
			plugin.ListPlugins(&cnf)
		} else {
			fmt.Println("Skipping plugins")
		}
		if themes {
			theme.ListThemes(&cnf)
		} else {
			fmt.Println("Skipping themes")
		}
	}
}

func UpdateCommand() func() {
	return func() {
		cmd := flag.NewFlagSet("update", flag.ExitOnError)
		var dryRun bool
		var stats bool
		cmd.BoolVar(&dryRun, "dry-run", false, "Perform an update dry run, this stops short of creating an update branches")
		cmd.BoolVar(&stats, "stats", true, "Login plugin, provider and repository names in your usage statistics")
		cmd.Parse(os.Args[2:])
		fmt.Println("Performing updates")

		cnf := config.LoadConfig()

		if dryRun == false {
			git.ConfigureGitConfig(&cnf)
			defer git.RestoreGitConfig(&cnf)
		}

		if cnf.Plugins.Enabled {
			fmt.Println("Performing plugin updates")
			plugin.UpdatePlugins(&cnf, dryRun, stats)
		} else {
			fmt.Println("Plugin updates disabled")
		}

		if cnf.Themes.Enabled {
			fmt.Println("Performing theme updates")
			theme.UpdateThemes(&cnf, dryRun, stats)
		} else {
			fmt.Println("Theme updates disabled")
		}
	}
}
