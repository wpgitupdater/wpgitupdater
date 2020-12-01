package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func CreateWorkflowTemplate() {
	template := `name: wpgitupdater
on:
  schedule:
  - cron: 0 0 * * *
  workflow_dispatch:
jobs:
  update:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v1
    - run: git checkout develop
    - run: curl ` + installerUrl + ` | bash -s -- -b $HOME/bin
    - run: $HOME/bin/wpgitupdater update
      env:
        WP_GIT_UPDATER_TOKEN: ${{ secrets.WP_GIT_UPDATER_TOKEN }}
        WP_GIT_UPDATER_GIT_TOKEN: ${{ secrets.GITHUB_TOKEN }}`

	if err := os.MkdirAll(filepath.Dir(workflowFile), os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(workflowFile, []byte(template), 644); err != nil {
		log.Fatal(err)
	}
	output := string(RunCmd("chmod", "644", workflowFile))
	fmt.Println(output)
}

func CreatePullRequest(config *Config, plugin Plugin) error {
	body := map[string]string{
		"title": plugin.GetPRTitle(config),
		"head":  plugin.GetBranchName(),
		"base":  config.Branch,
		"body":  "**Homepage:** " + plugin.Info.Homepage + "\n**Plugin Updated:** " + plugin.Info.LastUpdated + "\n\n**Changelog:**\n\n" + plugin.Info.Sections.Changelog,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	output := string(RunCmd("git", "remote", "get-url", "origin"))
	parts := strings.Split(strings.TrimSpace(output), "github.com/")
	url := "https://api.github.com/repos/" + strings.Replace(parts[1], ".git", "", 1) + "/pulls"
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "token "+config.Token)
	req.Header.Add("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(responseBody))

	return nil
}
