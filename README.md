<p align="center">
    <img src="https://docs.wpgitupdater.dev/wp-git-updater-hero.png" alt="WP Git Updater" width="300" />
</p>

# WP Git Updater

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/wpgitupdater/wpgitupdater/Go%20Build)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/wpgitupdater/wpgitupdater)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wpgitupdater/wpgitupdater)
[![GitHub issues](https://img.shields.io/github/issues/wpgitupdater/wpgitupdater)](https://github.com/wpgitupdater/wpgitupdater/issues)
[![GitHub stars](https://img.shields.io/github/stars/wpgitupdater/wpgitupdater)](https://github.com/wpgitupdater/wpgitupdater/stargazers)
[![GitHub license](https://img.shields.io/github/license/wpgitupdater/wpgitupdater)](https://github.com/wpgitupdater/wpgitupdater)

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
