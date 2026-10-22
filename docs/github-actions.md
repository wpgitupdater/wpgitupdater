# Github Actions

[Home](./README.md) | [Previous: Github](./github.md) | [Next: Configuration Reference](./configuration-reference.md)

## Workflow Configuration File

Using the [Init Command](./commands.md#init-ci-github-command) you can generate a Github Actions Workflow file to automate the cli tool updates.

```shell
wpgitupdater init -ci -actions
```

The workflow file will be called `wpgitupdater.yml` and will be located in the directory `.github/workflows`.

The contents of the file will be as follows:

```yml
name: wpgitupdater
on:
  schedule:
  - cron: 0 0 * * *
  workflow_dispatch:
jobs:
  update:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v1
    - run: git checkout develop
    - run: curl https://install.wpgitupdater.dev/install.sh | bash -s -- -b $HOME/bin
    - run: $HOME/bin/wpgitupdater update
      env:
        WP_GIT_UPDATER_GIT_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

Once generated you are free to alter the workflow or even generate your own workflow using the above as a reference.

## Schedule

You can see in the above configuration the template uses the [Workflow Schedule](https://docs.github.com/en/free-pro-team@latest/actions/reference/events-that-trigger-workflows#schedule) feature with the cron set to `0 0 * * *` (every day).

You can alter the schedule or remove entirely and rely on other triggers to automate the workflow.

## Workflow Dispatch

In addition to the cron schedule the template makes use of the [Workflow Dispatch](https://docs.github.com/en/free-pro-team@latest/actions/reference/events-that-trigger-workflows#workflow_dispatch) feature which will allow for manual invocation of the workflow.

## Authentication

Note in the above workflow file we have used `${{ secrets.GITHUB_TOKEN }}` as the `WP_GIT_UPDATER_GIT_TOKEN`.

Github Actions generate auth tokens for the user `github-actions` and these can be used in place of a Personal Access Token. With the above configuration the commits, and pull request generated would be created as the `github-actions` user.

As detailed in the [Authentication](./authentication.md) and [Github Authentication](./github.md#authentication) documentation if this is undesirable you are free to add your own Personal Access Token as an [Actions Secret](https://docs.github.com/en/free-pro-team@latest/actions/reference/encrypted-secrets) and reference this instead.
