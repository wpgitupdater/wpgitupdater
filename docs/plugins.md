# Plugins

[Home](./README.md) | [Previous: Configuration Reference](./configuration-reference.md)

## Display Git Status

[Display Git Status](https://wordpress.org/plugins/display-git-status/) is a pretty simple plugin, all it needs is access to the `shell_exec` function and to be pointed at a git repository.

The plugin will not perform any state altering operations, it accesses the repository using read only methods to fetch information like branch name, last commit and status.

When the git repository is located (defaulting to `wp-content`, or overwritten via a setting) the plugin will add a new admin bar item with the git icon and the current branch name.

When your site has untracked changes the menu items background will turn red.

This is ideal as it gives you a small marker directly in the admin area that some changes have been made which you probably need to look into.

Additionally when accessing the plugins admin page you can see the result of the `git status` command, and the last commit applied to your site.

Usage is pretty simple, install, set the repository location and it will display your repository information. Each element has detailed documentation in the plugins dedicated help guide found on the plugins admin page.
