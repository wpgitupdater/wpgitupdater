package interfaces

import "github.com/wpgitupdater/wpgitupdater/internal/config"

type Plugin interface {
	GetBranchName() string
	GetPRTitle(*config.Config) string
	GetHomePage() string
	GetLastUpdated() string
	GetChangelog() string
}
