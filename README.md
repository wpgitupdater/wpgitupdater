<p align="center">
    <img src="https://docs.wpgitupdater.dev/wp-git-updater-hero.png" alt="WP Git Updater" width="300" />
</p>

# WP Git Updater

![Workflow](https://github.com/wpgitupdater/wpgitupdater/workflows/Go%20Build/badge.svg)

Automated Source Controlled WordPress Updates

# Installation

```shell
curl https://install.wpgitupdater.dev/install.sh | bash -s -- -b $HOME/bin
```

# Usage

```shell
# Optional flags in []

# Generates a .wpgitupdater.yml file with defaults

$ wpgitupdater init

# Generates a .github/workflows/wpgitupdater.yml workflow file

$ wpgitupdater init -ci -actions

# You will need the ENV var WP_GIT_UPDATER_GIT_TOKEN set as a personal access token for the following commands
$ export WP_GIT_UPDATER_GIT_TOKEN="***"

# Lists plugin version stats

$ wpgitupdater list [-plugins]

# Performs updates

$ wpgitupdater update [-dry-run]
```

For more detailed documentation visit the [Documentation](https://docs.wpgitupdater.dev).
