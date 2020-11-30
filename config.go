package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

type PluginConfig struct {
	Enabled bool
	Path    string
	Commit  string
	Title   string
	Include []string
	Exclude []string
}

type Config struct {
	Cwd     string
	Branch  string
	Version string
	Token   string
	Plugins PluginConfig
}

func LoadConfig(cwd string) Config {
	plugins := PluginConfig{Path: "plugins"}
	config := Config{Cwd: cwd, Branch: CurrentBranch(), Token: GetToken(), Plugins: plugins}
	input, err := ioutil.ReadFile(cwd + "/.wpgitupdates.yml")
	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal(input, &config); err != nil {
		log.Fatal(err)
	}

	return config
}

func (config Config) GetPluginsPath(append string) string {
	path := config.Cwd + "/" + strings.Trim(config.Plugins.Path, "/")
	if append != "" {
		path = path + "/" + strings.Trim(append, "/")
	}
	return path
}

func (config Config) GetPluginsCommit() string {
	if config.Plugins.Commit != "" {
		return config.Plugins.Commit
	}
	return "chore(plugins): Update :plugin from :oldversion to :newversion"
}

func (config Config) GetPluginsPRTitle() string {
	if config.Plugins.Title != "" {
		return config.Plugins.Title
	}
	return "Update Plugin :plugin from :oldversion to :newversion"
}

func (config Config) PluginCanBeUpdated(slug string) bool {
	if len(config.Plugins.Include) > 0 {
		_, found := InSlice(config.Plugins.Include, slug)
		return found
	} else if len(config.Plugins.Exclude) > 0 {
		_, found := InSlice(config.Plugins.Exclude, slug)
		return !found
	} else {
		return true
	}
}
