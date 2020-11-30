package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const version = "1.0"
const configFile = ".wpgitupdater.yml"

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Expected 'init', 'list' or 'update' subcommands")
	}

	switch os.Args[1] {
	case "init":
		fmt.Println("Creating config file")
		CreateConfigTemplate()
		fmt.Println("Config file created!")
	case "list":
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
	case "update":
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
