package config

type Task struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Author   string `json:"author"`

	AutoMake        bool   `json:"autoMake"`
	LaunchArgs      string `json:"launchArgs"`
	ExternalScraper bool   `json:"externalScraper"`
	PAUrl           string `json:"paUrl"`
	PreProcess      bool   `json:"preprocess"`

	BuildRequirement       []string               `json:"buildRequirement"`
	ReleaseRequirement     []string               `json:"releaseRequirement"`
	ExternalScraperOptions ExternalScraperOptions `json:"externalScraperOptions"`
}

const (
	ES_POLICY_SILENT = "silent"
	ES_POLICY_MANUAL = "manual"
)

type ExternalScraperOptions struct {
	Policy             string `json:"policy"`
	ReleaseInstaller   bool   `json:"releaseInstaller"`
	SilentArg          string `json:"silentArg"`
	CompressLevel      uint8  `json:"compressLevel"`
	SlientDelete       bool   `json:"slientDelete"`
	ManualShortcutName string `json:"manualShortcutName"`
}
