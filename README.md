# wpgitupdater
CI tool to automate WordPress asset updates for Git controlled websites

Generating install script (https://github.com/goreleaser/godownloader)

You need to create a master branch and run:

`godownloader --repo=wpgitupdater/wpgitupdater > install.sh`

Then delete master branch.

# Usage

```
# Optional flags in []

# Generates a .wpgitupdater.yml file with defaults

$ wpgitupdater init

# Generates a .github/workflows/wpgitupdater.yml workflow file

$ wpgitupdater init -workflow

# You will need the ENV var WP_GIT_UPDATER_GIT_TOKEN set as a personal access token for the following commands
$ export WP_GIT_UPDATER_GIT_TOKEN="***"

# Lists plugin version stats

$ wpgitupdater list [-plugins|-plugins=false]

# Performs updates

$ wpgitupdater update [-dry-run]
```
