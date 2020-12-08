package config

import (
	"fmt"
	"github.com/wpgitupdater/wpgitupdater/internal/constants"
	"github.com/wpgitupdater/wpgitupdater/internal/utils"
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
	Cwd          string
	Branch       string
	Version      string
	Token        string
	UpdaterToken string
	Plugins      PluginConfig
}

func CreateConfigTemplate() {
	template := `version: "` + constants.ConfigVersion + `"
plugins:
  enabled: true
  path: plugins`
	if err := ioutil.WriteFile(constants.ConfigFile, []byte(template), 644); err != nil {
		log.Fatal(err)
	}
	output := string(utils.RunCmd("chmod", "644", constants.ConfigFile))
	fmt.Println(output)
}

func LoadConfig() Config {
	plugins := PluginConfig{Path: "plugins"}
	config := Config{Cwd: utils.GetCwd(), Token: utils.GetToken(), UpdaterToken: utils.GetUpdaterToken(), Plugins: plugins}
	input, err := ioutil.ReadFile(config.Cwd + "/" + constants.ConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal(input, &config); err != nil {
		log.Fatal(err)
	}

	if _, exists := utils.InSlice(constants.SupportedConfigVersions[:], config.Version); !exists {
		log.Println("Configuration file version unsupported! Please ensure you match the config with the updaters supported versions.")
		log.Println("Configuration version [" + config.Version + "]")
		log.Println("Updater version [" + constants.Version + "]")
		log.Fatal("Supported configuration versions [" + strings.Join(constants.SupportedConfigVersions[:], ",") + "]")
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
		_, found := utils.InSlice(config.Plugins.Include, slug)
		return found
	} else if len(config.Plugins.Exclude) > 0 {
		_, found := utils.InSlice(config.Plugins.Exclude, slug)
		return !found
	} else {
		return true
	}
}
