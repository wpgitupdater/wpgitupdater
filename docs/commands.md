# Commands

[Home](./README.md) | [Previous: Configuration File](./configuration.md) | [Next: Authentication](./authentication.md)

Below are the commands available from the `wpgitupdater` cli tool.

Note the commands marked as **(authenticated)** will require the environment variable `WP_GIT_UPDATER_GIT_TOKEN`. Please refer to the [Authentication](./authentication.md) documentation for further details.

## Init Command

```shell
wpgitupdater init
```

The `init` command will generate a `.wpgitupdater.yml` in the current working directory.

You should run this command from the root of your WordPress repository. The command will include only the minimum information needed to provide a valid configuration.

Please review the file after creation and make any amendments required. Pay special attention to the `plugins.path` and `themes.path` variables to ensure it is a valid path to your plugin and theme folders.

Please review the [Configuration File](./configuration.md) documentation for all of the possible options and there behaviour.

## Init CI Github Command

```shell
wpgitupdater init -ci -actions
```

The `init -ci -actions` command will generate a `.github/workflows/wpgitupdater.yml` Github actions CI configuration file.

The Action will allow for manual invocation and also configure a daily invocation using the cron `0 0 * * *`.

Please refer to the [Github Actions](./github-actions.md) documentation for further details.

## List Command (authenticated)

```shell
wpgitupdater list
```

The `list` command will load your configuration and display the status of your installed themes and plugins, noting any that are outdated.

If you prefer to just see either theme or plugins, you can disable one or the other using the following commands

```shell
wpgitupdater list -plugins=false
wpgitupdater list -themes=false
```

## Update Command (authenticated)

```shell
wpgitupdater update
```

The `update` command will load your configuration then inspect available updates.

When an update is found the tool will create a new branch, update the plugin or theme, commit and push the branch, and finally open a pull request with the update details.

When complete the tool will restore the base branch and continue onto the next update until all possible updates have been performed.

#### Dry Run

The `update` command has an optional flag which will prevent any updates being performed, no branches, changes or pull requests will be created when the `-dry-run` flag is present, you can use this flag to inspect the output and review what may be updated.

```shell
wpgitupdater update -dry-run
```

#### Stats

The old `update -stats` flag is accepted for compatibility with earlier versions and has no effect. The tool does not send usage statistics.
