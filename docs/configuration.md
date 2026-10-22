# Configuration File

[Home](./README.md) | [Previous: Getting Started](./getting-started.md) | [Next: Commands](./commands.md)

The `.wpgitupdater.yml` configuration file determines how the cli tool operates across all environments.

For a full configuration file reference please visit the page [here](./configuration-reference.md).

The configuration file should be included in source control and located at the repository root. It can be created using the [Init Command](./commands.md#init-command).

Note below is an example of the minimal configuration needed which should cover the majority of use cases:

```yaml
version: "1.0"
plugins:
  enabled: true
  path: plugins
themes:
  enabled: true
  path: themes
```

The version of the configuration file must match with the major and minor version of the cli tool. You should consider updating your configuration file format when new versions become available, however the installer script does allow for installing historic versions of the tool. Please refer to the [Installing Versions](./getting-started.md#installing-versions) documentation to require historic versions.

Below is an example file, using the excluded option to exclude certain plugins from the update, provide a custom path to plugins `wp-content/plugins` and a custom pull request title template.

```yaml
version: "1.0"
plugins:
  enabled: true
  path: wp-content/plugins
  title: "Custom PR Title: Updated :plugin from :oldversion to :newversion"
  exclude:
    - woocommerce
    - wordpress-seo
themes:
  enabled: true
  path: wp-content/themes
```
