package config

import (
	"os"

	"github.com/hyroge/pluginbot/utils/json"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

var config *BuildConfig

type BuildConfig struct {
	EnableRemote   bool     `json:"enableRemote"`
	IgnoreRemote   bool     `json:"ignoreRemote"`
	TasksPath      string   `json:"dirTask"`
	WorkspacePath  string   `json:"dirWorkshop"`
	DistPath       string   `json:"dirBuilds"`
	BuildInfoPath  string   `json:"pathDatabase"`
	MaxBuildsNum   uint     `json:"maxBuildsNum"`
	RemoteName     string   `json:"remoteName"`
	RemoteRoot     string   `json:"remoteRoot"`
	Aria2Port      uint     `json:"aria2Port"`
	Aria2Host      string   `json:"aria2Host"`
	Aria2Secret    string   `json:"aria2Secret"`
	Aria2SpawnArgs []string `json:"aria2SpawnArgs"`
}

func FetchBuildConfig() *BuildConfig {
	if config == nil {
		config = resolveBuildConfig()
	}
	return config
}

func resolveBuildConfig() *BuildConfig {
	args := FetchCmdArgs()

	LogDebug("[config/build] try to resolve build config")
	LogDebug("[config/build] check file")
	_, err := os.Stat(args.BuildConfig)
	Must(err)

	LogDebug("[config/build] resolve...")
	f, err := os.Open(args.BuildConfig)
	Must(err)
	defer f.Close()

	var config *BuildConfig
	Must(json.UnmarshalJsonc(f, &config))
	return config
}
