package main

import (
	"flag"
	"fmt"
	"github.com/wpgitupdater/wpgitupdater/internal/config"
	"github.com/wpgitupdater/wpgitupdater/internal/constants"
	"github.com/wpgitupdater/wpgitupdater/internal/github"
	"github.com/wpgitupdater/wpgitupdater/internal/plugin"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Println("WordPress Git Updater V" + constants.Version)
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
		cmd.Parse(os.Args[2:])
		fmt.Println("List update statuses")

		cnf := config.LoadConfig()
		if plugins {
			plugin.ListPlugins(&cnf)
		} else {
			fmt.Println("Skipping plugins")
		}
	}
}

func UpdateCommand() func() {
	return func() {
		cmd := flag.NewFlagSet("update", flag.ExitOnError)
		var dryRun bool
		cmd.BoolVar(&dryRun, "dry-run", false, "Perform an update dry run, this stops short of creating an update branch")
		cmd.Parse(os.Args[2:])
		fmt.Println("Performing updates")

		cnf := config.LoadConfig()
		if cnf.Plugins.Enabled {
			plugin.UpdatePlugins(&cnf, dryRun)
		} else {
			fmt.Println("Plugin updates disabled")
		}
	}
}
