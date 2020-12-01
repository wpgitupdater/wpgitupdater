package main

import (
	"fmt"
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

func CreateConfigTemplate() {
	template := `version: "` + version + `"
plugins:
  enabled: true
  path: plugins`
	if err := ioutil.WriteFile(configFile, []byte(template), 644); err != nil {
		log.Fatal(err)
	}
	output := string(RunCmd("chmod", "644", configFile))
	fmt.Println(output)
}

func LoadConfig() Config {
	plugins := PluginConfig{Path: "plugins"}
	config := Config{Cwd: GetCwd(), Branch: CurrentBranch(), Token: GetToken(), Plugins: plugins}
	input, err := ioutil.ReadFile(config.Cwd + "/" + configFile)
	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal(input, &config); err != nil {
		log.Fatal(err)
	}

	if config.Version != version {
		log.Println("Configuration file version unsupported! Please ensure you match the config and updater versions.")
		log.Println("Configuration version [" + config.Version + "]")
		log.Fatal("Updater version [" + version + "]")
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
	return "Update plugin :plugin from :oldversion to :newversion"
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
