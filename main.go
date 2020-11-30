package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "1.0"

func main() {

	EnsureValidCommand()
	config := LoadConfig()
	EnsureValidVersion(&config)

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	var listPlugins bool
	listCmd.BoolVar(&listPlugins, "plugins", true, "List plugin updates")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	var dryRun bool
	updateCmd.BoolVar(&dryRun, "dry-run", false, "Perform an update dry run, this stops short of creating an update branch")

	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
		fmt.Println("List update statuses")

		if listPlugins {
			ListPlugins(&config)
		} else {
			fmt.Println("Skipping plugins")
		}
	case "update":
		updateCmd.Parse(os.Args[2:])
		fmt.Println("Performing updates")

		if config.Plugins.Enabled {
			UpdatePlugins(&config, dryRun)
		} else {
			fmt.Println("Plugin updates disabled")
		}
	}
}

func EnsureValidCommand() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'list' or 'update' subcommands")
		os.Exit(1)
	}
}

func EnsureValidVersion(config *Config) {
	if config.Version != version {
		fmt.Println("Configuration file version unsupported! Please ensure your match the config and updater versions.")
		fmt.Println("Configuration version [" + config.Version + "]")
		fmt.Println("Updater version [" + version + "]")
		os.Exit(1)
	}
}
