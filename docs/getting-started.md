# Getting Started

[Home](./README.md) | [Previous: Introduction](./introduction.md) | [Next: Configuration File](./configuration.md)

## Quick Start

First you may want to install the cli tool locally.

```shell
curl https://install.wpgitupdater.dev/install.sh | bash -s -- -b $HOME/bin
```

In your Git repository root run the command to generate the [Configuration File](./configuration.md):

```shell
wpgitupdater init
```

And then to generate a CI configuration (in this case for [Github Actions](https://github.com/features/actions)) run the following command:

```shell
wpgitupdater init -ci -actions
```

You can browse the files generated and commit them to your repository.

Now every day (assuming the standard configuration is used) an automated workflow will be run, checking your repositories plugins for updates and creating pull requests when updates are available.

## Installing Versions

While we would advise you to maintain compatibility with the "latest" release the installer script does support installing historic versions, in the example below we are specially requesting the `1.0.0` version:

```shell
curl https://install.wpgitupdater.dev/install.sh | bash -s -- -b $HOME/bin 1.0.0
```

Additionally you can install historic versions using the [manual install](#manual-installation) process.

## Troubleshooting

### Manual Installation

In addition to the installer script you can [manually download a release](https://github.com/wpgitupdater/wpgitupdater/releases) from our releases page.

### Installation Path

The install command will place the `wpgitupdater` binary in your `$HOME/bin` folder.

You can test if this is accessible by running:

```shell
which wpgitupdater
```

If you are presented with a path to the cli tool the installation was successful.

If `$HOME/bin` is not in your `$PATH` you can choose to add it:

```shell
echo "export PATH=\"\$PATH:\$HOME/bin\"" >> ~/.profile && source ~/.profile
```

Or update the installer script to place it in a directory that is in your `$PATH` (using `/usr/local/bin` as an example below):

```shell
curl https://install.wpgitupdater.dev/install.sh | bash -s -- -b /usr/local/bin
```
