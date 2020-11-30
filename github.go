package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// https://developer.github.com/v3/pulls/#create-a-pull-request
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
	req.Header.Add("User-Agent", "wpgitupdater")
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
