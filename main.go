package main

import (
	"fmt"
)

func main() {
	cwd := GetCwd()
	config := LoadConfig(cwd)

	if config.Plugins.Enabled == false {
		fmt.Println("Plugin updates disabled")
		return
	}

	plugins := GetPlugins(&config)

	// List update statuses
	if false {
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

	// Perform updates
	if true {
		ConfigureGitConfig(&config)

		for _, plugin := range plugins {
			plugin.PerformPluginUpdate(&config)
		}

		RestoreGitConfig(&config)
	}
}
