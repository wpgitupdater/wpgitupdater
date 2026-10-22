# Introduction

[Home](./README.md) | [Next: Getting Started](./getting-started.md)

WP Git Updater is the combination of a cli tool and popular Continuous Integration service configuration files.

It is designed to be run on a schedule and will update your Git source controlled WordPress websites third party themes and plugins.

The cli tool can be used manually or with one of the providers set out in the documentation.

## How It Works

During an update the cli tool will collect any detected themes and plugins. And in line with your configuration selectively update them one by one, creating a new branch for each update and submitting a pull request for you should an update be found.

The tool updates any themes and plugins with new versions that can be found in the [https://wordpress.org](https://wordpress.org) repositories. The cli tool is built in go lang for optimal distribution across CI providers, as a result only plugins or themes found on the official WordPress repositories can be updated. Premium and custom themes or plugins are skipped during the discovery phase of the update.

## Why?

WordPress already has a long history of providing seamless updates, and more recently you can even enable automatic updates from within the admin area.

Additionaly using the `wpcli` tool you can perform updates without using the administration area.

This works great for most WordPress sites.

However there are many cases where this can actually be a problem.

Specifically when a WordPress website is built and maintained using Git source control. Not only are the first party update features not required, they can actually be a problem causing deployments and source to detach.

There are many reasons to maintain a WordPress in Git source control, from hosting integration to management and development of more complex sites requiring staging/develop and production environments.

In these cases there is additional workload in managing theme and plugin versions and staying up to date.

It is these cases where Wp Git Updater shines.

When combined with a scheduled CI system you can be notified of new updates, and have the changes ready to review automatically.

## Example

So you can see WP Git Updater in action head over to our [example repository](https://github.com/wpgitupdater/example) which has been setup to perform updates. You can see the commits and Pull Requests created by the tool.
