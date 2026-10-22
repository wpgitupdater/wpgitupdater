# Github

[Home](./README.md) | [Previous: Authentication](./authentication.md) | [Next: Github Actions](./github-actions.md)

## Authentication

With Github as a Git provider for your repository your `WP_GIT_UPDATER_GIT_TOKEN` will use what's known as a "Personal Access Token".

Full details on how to generate a Personal Access Token can be found here: [Personal Access Tokens](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token).

The token needs only the `repo` scope when generating. When created please make a note of the token and follow the instructions on the [authentication](./authentication.md) page to expose the token as an environment variable.

Use of the Personal Access Token will result in the cli tool "acting" as the user who provided the token. This may be desirable, or alternatively you could create a "bot account" specifically for the purpose of identifying automated processes.

## Pull Requests

As part of the update process, when a branch is created the cli tool will then use the same authentication token to generate a pull request. The pull request will be created requesting the update branch be merged into the base branch (usually develop).

The pull request will contain information about the cli build tool that generated it, the update details and crucially any changelog information provided for the update itself (unfortunatley themes don't provide changelog information, so this will be ommitted for theme pull requests). Along with the pull request file change diff this should provide the reviewer with all the information needed to make an informed decision on merging the request.
