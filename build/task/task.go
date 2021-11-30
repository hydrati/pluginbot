package task

import (
	"errors"
	"path"

	"github.com/hyroge/pluginbot/config"
	"github.com/hyroge/pluginbot/utils/fs"
	// . "github.com/hyroge/pluginbot/utils/prelude"
)

var (
	ERR_CANNOT_ACCESS       = errors.New("cannot access the path")
	ERR_TASK_NAME_NOT_EQUAL = errors.New("the task name in path is not equal to the task name in config")
)

func ResolveTaskSpace(name string) (*config.Task, error) {
	cfg := config.FetchBuildConfig()
	task_path := path.Join(cfg.TasksPath, name)
	task_cfg := path.Join(task_path, "config.json")

	if !(fs.IsAccessible(task_path) || fs.IsAccessible(task_cfg)) {
		return nil, ERR_CANNOT_ACCESS
	}

	task, err := config.ResolveTaskFromPath(task_cfg)
	if err != nil {
		return nil, err
	}

	if task.Name != name {
		return nil, ERR_TASK_NAME_NOT_EQUAL
	}

	if task.ExternalScraper != nil && *task.ExternalScraper {
		if err = CheckExternalScraperOptions(task); err != nil {
			return nil, err
		}
	}

	// TODO: check category

	return task, nil
}

var (
	ERR_ES_CFG_NEED_POLICY    = errors.New("need provide externalScraperOptions.policy as external scraper task")
	ERR_ES_CFG_UNKNOWN_POLICY = errors.New("unknown policy")
	ERR_ES_CFG_NEED_BUILD_REQ = errors.New("need provide buildRequirement as external scraper task with make.cmd")
	ERR_ES_CFG_NEED_REL_REQ   = errors.New("need provide releaseRequirement as external scraper task with make.cmd")
)

func CheckExternalScraperOptions(task *config.Task) error {
	if task.AutoMake {
		if task.ExternalScraperOptions.Policy == nil {
			return ERR_ES_CFG_NEED_POLICY
		}
		if *task.ExternalScraperOptions.Policy != config.ES_POLICY_MANUAL || *task.ExternalScraperOptions.Policy != config.ES_POLICY_SILENT {
			return ERR_ES_CFG_UNKNOWN_POLICY
		}
	} else {
		if task.BuildRequirement == nil {
			return ERR_ES_CFG_NEED_BUILD_REQ
		}
	}

	if task.ExternalScraperOptions.ReleaseInstaller != nil {
		if *task.ExternalScraperOptions.ReleaseInstaller && task.ReleaseRequirement == nil {
			return ERR_ES_CFG_NEED_REL_REQ
		}
	}

	return nil
}
