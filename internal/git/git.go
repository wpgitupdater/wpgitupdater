package git

import (
	"fmt"
	"github.com/wpgitupdater/wpgitupdater/internal/config"
	"github.com/wpgitupdater/wpgitupdater/internal/constants"
	"github.com/wpgitupdater/wpgitupdater/internal/utils"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func ConfigureGitConfig(cnf *config.Config) {
	gitConfigFile := cnf.Cwd + "/.git/config"
	fmt.Println(fmt.Sprintf("Configuring git config using token: %s", cnf.Token))

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
	output := string(utils.RunCmd("git", "config", "user.email", constants.GitEmail))
	if output != "" {
		fmt.Println(output)
	}

	fmt.Println("Setting committer name")
	output = string(utils.RunCmd("git", "config", "user.name", constants.GitUser))
	if output != "" {
		fmt.Println(output)
	}

	fmt.Println("Updating origin url")
	output = string(utils.RunCmd("git", "remote", "get-url", "origin"))
	url := strings.TrimSpace(output)
	re := regexp.MustCompile("^(git@|https://)([^:/]+)[:/](.+)")
	origin := re.ReplaceAllString(url, fmt.Sprintf("https://x-access-token:%v@$2/$3", cnf.Token))
	output = string(utils.RunCmd("git", "remote", "set-url", "origin", origin))
	if output != "" {
		fmt.Println(output)
	}
}

func RestoreGitConfig(cnf *config.Config) {
	fmt.Println("Restoring git config")
	gitConfigFile := cnf.Cwd + "/.git/config"
	err := os.Remove(gitConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename(gitConfigFile+".original", gitConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	utils.RunCmd("chmod", "644", gitConfigFile)
}

func CurrentBranch() string {
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = utils.GetCwd()
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(output))
}

func BranchExists(branch string) bool {
	cmd := exec.Command("git", "ls-remote", "--exit-code", "--heads", "origin", branch)
	cmd.Dir = utils.GetCwd()
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return true
}
