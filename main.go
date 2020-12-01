package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const version = "1.0"
const configFile = ".wpgitupdater.yml"
const gitUser = "WordPress Git Updater Bot"
const gitEmail = "bot@wpgitupdater.dev"
const userAgent = "wpgitupdater"
const workflowFile = ".github/workflows/wpgitupdater.yml"
const installerUrl = "https://wpgitupdater.github.io/installer/install.sh"

func main() {

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
		var workflow bool
		cmd.BoolVar(&workflow, "workflow", false, "Create Github Actions workflow file")
		cmd.Parse(os.Args[2:])

		if workflow {
			fmt.Println("Creating workflow file")
			CreateWorkflowTemplate()
			fmt.Println("Workflow file created!")
		} else {
			fmt.Println("Creating config file")
			CreateConfigTemplate()
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

		config := LoadConfig()
		if plugins {
			ListPlugins(&config)
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

		config := LoadConfig()
		if config.Plugins.Enabled {
			UpdatePlugins(&config, dryRun)
		} else {
			fmt.Println("Plugin updates disabled")
		}
	}
}
