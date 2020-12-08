package interfaces

type Api interface {
	UpdateUsage(usageType string, slug string) error
}

type UpdateUsage struct {
	Provider   string          `json:"provider"`
	Repository string          `json:"repository"`
	Type       string          `json:"type"`
	Slug       string          `json:"slug"`
	Meta       UpdateUsageMeta `json:"meta"`
}

type UpdateUsageMeta struct {
	Build         string `json:"build"`
	Version       string `json:"version"`
	ConfigVersion string `json:"configVersion"`
}
