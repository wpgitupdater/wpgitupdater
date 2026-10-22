# Authentication

[Home](./README.md) | [Previous: Commands](./commands.md) | [Next: Github](./github.md)

Note that some commands are marked as **(authenticated)**.

This means the operations performed communicate with your Git repository. They require the environment variable `WP_GIT_UPDATER_GIT_TOKEN` to be available which is used for authentication with the Git provider.

No separate WP Git Updater account, service token, or license is required.

Each CI Provider may require or provide ways to consume a token generated as part of the CI process. These will be detailed in their respective sections.

However to run these operations locally if you wish to you will need to provide these tokens within your own environment.

You can set the environment variables in a few ways.

## Current Cli Session

You can export variables for your current session using the following command:

```shell
export WP_GIT_UPDATER_GIT_TOKEN="***"
```

Note that once your session has ended (terminal closed) these variables will be lost.

## Persistent

To persist the environment variables:

```shell
echo "export WP_GIT_UPDATER_GIT_TOKEN=\"***\"" >> ~/.profile
source ~/.profile
```

You may also to use `.bash_profile`, `.zsh`, or `.bashrc` among others depending on your preference or system.
