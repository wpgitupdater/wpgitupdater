# Configuration Reference

[Home](./README.md) | [Previous: Github Actions](./github-actions.md) | [Next: Plugins](./plugins.md)

Below are the config values and their defaults for the `.wpgitupdater.yml` file.

## version

- Type: `string`
- Default: `1.0`

This is the version of the config file.

This should match the version of the `wpgitupdater` binary being used to perform updates.

## branch

- Type: `string`
- Default: `git branch --show-current`

The branch used as a base for the updates.

This is usually `develop` allowing updates and pull requests to be targeted against a none production environment.

## plugins

### plugins.enabled

- Type: `boolean`
- Default: `false`

To enable plugin updates you must set this to `true`.

### plugins.commit

- Type: `string`
- Default: `chore(plugins): Update :plugin from :oldversion to :newversion`

This is the generated commit message, the values `:plugin`, `:oldversion`, and `:newversion` will be replaced with the respective information on commit.

### plugins.title

- Type: `string`
- Default: `Update plugin :plugin from :oldversion to :newversion`

This is the generated pull request title, the values `:plugin`, `:oldversion`, and `:newversion` will be replaced with the respective information on creation.

### plugins.path

- Type: `string`
- Default: `plugins`

Where to look for plugins, this should be exclusive of any prepended or trailing slashes.

### plugins.include

- Type: `string[]`
- Default: `[]`

To enable fine grain control over which plugins can be updated you can use the `include` value to whitelist plugins.

When enabled **no other plugin updates will be attempted**.

### plugins.exclude

- Type: `string[]`
- Default: `[]`

This is the inverse of `include` and will only be used if a whitelist doesn't exist in the `include` value.

When enabled any plugin not in this list will be considered suitable for updates.

## themes

### themes.enabled

- Type: `boolean`
- Default: `false`

To enable theme updates you must set this to `true`.

### themes.commit

- Type: `string`
- Default: `chore(themes): Update :theme from :oldversion to :newversion`

This is the generated commit message, the values `:theme`, `:oldversion`, and `:newversion` will be replaced with the respective information on commit.

### themes.title

- Type: `string`
- Default: `Update theme :theme from :oldversion to :newversion`

This is the generated pull request title, the values `:theme`, `:oldversion`, and `:newversion` will be replaced with the respective information on creation.

### themes.path

- Type: `string`
- Default: `themes`

Where to look for themes, this should be exclusive of any prepended or trailing slashes.

### themes.include

- Type: `string[]`
- Default: `[]`

To enable fine grain control over which themes can be updated you can use the `include` value to whitelist themes.

When enabled **no other theme updates will be attempted**.

### themes.exclude

- Type: `string[]`
- Default: `[]`

This is the inverse of `include` and will only be used if a whitelist doesn't exist in the `include` value.

When enabled any theme not in this list will be considered suitable for updates.
