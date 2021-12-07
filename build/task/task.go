package task

import (
	"os"

	"github.com/hyroge/pluginbot/utils/json"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

var (
	TASK_CATEGORIES = []string{"实用工具", "开发辅助", "配置检测", "资源管理", "办公编辑", "输入法", "集成开发", "录屏看图", "磁盘数据", "安全急救", "网课会议", "即时通讯", "安装备份", "游戏娱乐", "运行环境", "压缩镜像", "美化增强", "驱动管理", "下载上传", "浏览器", "影音播放", "远程连接"}
)

type Task struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Author   string `json:"author"`

	AutoMake        bool    `json:"autoMake"`
	LaunchArgs      *string `json:"launchArgs"`
	ExternalScraper *bool   `json:"externalScraper"`
	PAUrl           *string `json:"paUrl"`
	PreProcess      *bool   `json:"preprocess"`

	BuildRequirement       *[]string               `json:"buildRequirement"`
	ReleaseRequirement     *[]string               `json:"releaseRequirement"`
	ExternalScraperOptions *ExternalScraperOptions `json:"externalScraperOptions"`
}

const (
	ES_POLICY_SILENT = "silent"
	ES_POLICY_MANUAL = "manual"
)

type ExternalScraperOptions struct {
	Policy             *string `json:"policy"`
	ReleaseInstaller   *bool   `json:"releaseInstaller"`
	SilentArg          *string `json:"silentArg"`
	CompressLevel      *uint8  `json:"compressLevel"`
	SlientDelete       *bool   `json:"slientDelete"`
	ManualShortcutName *string `json:"manualShortcutName"`
}

func ResolveTaskFromPath(path string) (*Task, error) {
	LogDebug("[config/task] try to resolve task config")
	LogDebug("[config/task] check file")
	_, err := os.Stat(path)
	if err != nil {
		LogError("[config/task] check file error")
		return nil, err
	}
	LogDebug("[config/build] resolve...")
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
	LogDebug("[config/task] resolved, name = %s", config.Name)
	return config, nil
}
