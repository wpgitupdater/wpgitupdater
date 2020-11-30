# wpgitupdater
CI tool to automate WordPress asset updates for Git controlled websites

# Generating install script (https://github.com/goreleaser/godownloader)

You need to set repo to public, create a master branch and run:

`godownloader --repo=wpgitupdater/wpgitupdater > install.sh`

Switch repo back and delete master branch.

# Usage

```
# You will need the ENV var WP_GIT_UPDATER_GIT_TOKEN set as a personal access token
$ export WP_GIT_UPDATER_GIT_TOKEN="***"

$ wpgitupdater [dir = cwd] [--dry-run]
```
