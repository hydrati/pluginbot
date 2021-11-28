package config

import (
	"os"

	"github.com/hyroge/pluginbot/utils/json"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

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

func ResolveFromPath(path string) (*Task, error) {
	LogInfo("[config/task] try to resolve task config")
	LogInfo("[config/task] check file")
	_, err := os.Stat(path)
	if err != nil {
		LogError("[config/task] check file error")
		return nil, err
	}
	LogInfo("[config/build] resolve...")
	f, err := os.Open(path)
	if err != nil {
		LogError("[config/task] open file error")
		return nil, err
	}
	defer f.Close()
	var config *Task
	err = json.UnmarshalJsonc(f, &config)
	if err != nil {
		LogError("[config/task] resolve task error")
		return nil, err
	}
	return config, nil
}
