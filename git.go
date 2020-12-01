package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func ConfigureGitConfig(config *Config) {
	gitConfigFile := config.Cwd + "/.git/config"
	fmt.Println(fmt.Sprintf("Configuring git config using token: %s", config.Token))

	fmt.Println("Creating git config backup")
	input, err := ioutil.ReadFile(gitConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(gitConfigFile+".original", input, 644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Setting committer email address")
	output := string(RunCmd("git", "config", "user.email", gitEmail))
	if output != "" {
		fmt.Println(output)
	}

	fmt.Println("Setting committer name")
	output = string(RunCmd("git", "config", "user.name", gitUser))
	if output != "" {
		fmt.Println(output)
	}

	fmt.Println("Updating origin url")
	output = string(RunCmd("git", "remote", "get-url", "origin"))
	url := strings.TrimSpace(output)
	re := regexp.MustCompile("^(git@|https://)([^:/]+)[:/](.+)")
	origin := re.ReplaceAllString(url, fmt.Sprintf("https://x-access-token:%v@$2/$3", config.Token))
	output = string(RunCmd("git", "remote", "set-url", "origin", origin))
	if output != "" {
		fmt.Println(output)
	}
}

func RestoreGitConfig(config *Config) {
	fmt.Println("Restoring git config")
	gitConfigFile := config.Cwd + "/.git/config"
	err := os.Remove(configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename(gitConfigFile+".original", gitConfigFile)
	if err != nil {
		log.Fatal(err)
	}
}

func CurrentBranch() string {
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = GetCwd()
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(output))
}

func BranchExists(branch string) bool {
	cmd := exec.Command("git", "ls-remote", "--exit-code", "--heads", "origin", branch)
	cmd.Dir = GetCwd()
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return true
}
