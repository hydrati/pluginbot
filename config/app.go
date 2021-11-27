package config

type AppConfig struct {
	EnableRemote   string   `json:"enableRemote"`
	IgnoreRemote   string   `json:"ignoreRemote"`
	DirTasks       string   `json:"dirTasks"`
	DirWorkshop    string   `json:"dirWorkshop"`
	DirBuilds      string   `json:"dirBuilds"`
	PathDatabase   string   `json:"pathDatabase"`
	MaxBuildsNum   uint     `json:"maxBuildsNum"`
	RemoteName     string   `json:"remoteName"`
	RemoteRoot     string   `json:"remoteRoot"`
	Aria2Port      uint16   `json:"aria2Port"`
	Aria2Host      string   `json:"aria2Host"`
	Aria2Secret    string   `json:"aria2Secret"`
	Aria2SpawnArgs []string `json:"aria2SpawnArgs"`
}
