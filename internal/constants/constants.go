package constants

var Build = "@dev-build"
var BuildDate = ""
var Version = "@dev-version"
var SupportedConfigVersions = [1]string{"1.0"}

const ConfigFile = ".wpgitupdater.yml"
const ConfigVersion = "1.0"
const GitUser = "WP Git Updater Bot"
const GitEmail = "bot@wpgitupdater.dev"
const UserAgent = "wpgitupdater"
const WorkflowFile = ".github/workflows/wpgitupdater.yml"
const InstallerUrl = "https://install.wpgitupdater.dev/install.sh"
const ApiUrl = "https://wpgitupdater.dev/api/v1"

const WordPressPluginApiInfo = "https://api.wordpress.org/plugins/info/1.2/?action=plugin_information&request[slug]="
