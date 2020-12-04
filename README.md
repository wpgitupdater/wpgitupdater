# WP Git Updater

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/wpgitupdater/wpgitupdater/Go%20Build)](https://github.com/wpgitupdater/wpgitupdater/actions)
[![GitHub Release (latest by date)](https://img.shields.io/github/v/release/wpgitupdater/wpgitupdater)](https://github.com/wpgitupdater/wpgitupdater/releases)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wpgitupdater/wpgitupdater)](https://github.com/wpgitupdater/wpgitupdater)
[![GitHub Issues](https://img.shields.io/github/issues/wpgitupdater/wpgitupdater)](https://github.com/wpgitupdater/wpgitupdater/issues)
[![GitHub Stars](https://img.shields.io/github/stars/wpgitupdater/wpgitupdater)](https://github.com/wpgitupdater/wpgitupdater/stargazers)
[![GitHub License](https://img.shields.io/github/license/wpgitupdater/wpgitupdater)](https://github.com/wpgitupdater/wpgitupdater)

Automated Source Controlled WordPress Updates

[![WP Git Updater](https://blog.wpgitupdater.dev/wp-content/uploads/2020/12/contributions-2-fb-post.jpg)](https://wpgitupdater.dev)

# Installation

[![curl https://install.wpgitupdater.dev/install.sh | bash -s -- -b $HOME/bin](./install.svg)](#usage)

# Usage

```shell
# Install locally
curl https://install.wpgitupdater.dev/install.sh | bash -s -- -b $HOME/bin

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
