package prepare

import (
	"github.com/hyroge/pluginbot/config"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

func GetWorkspaceReady() error {
	LogInfo("try to get workspace ready")
	_ = config.FetchBuildConfig()
	return nil
}
