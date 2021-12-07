package task

import (
	"container/list"
	"errors"
	"os"
	"path"

	"github.com/hyroge/pluginbot/config"
	"github.com/hyroge/pluginbot/utils/fs"
	. "github.com/hyroge/pluginbot/utils/prelude"
	"github.com/hyroge/pluginbot/utils/slices"
)

var (
	ERR_CANNOT_ACCESS       = errors.New("cannot access the path")
	ERR_INVALID_CATE        = errors.New("invalid category")
	ERR_TASK_NAME_NOT_EQUAL = errors.New("the task name in path is not equal to the task name in config")
)

func CheckResolveAllTaskInTasksFolder() (*list.List /* List[*Task] */, error) {
	cfg := config.FetchBuildConfig()
	return CheckResolveAllTaskInFolder(cfg.TasksPath)
}

func CheckResolveAllTaskInFolder(p string) (*list.List /* List[*Task] */, error) {
	if e, err := fs.IsDirectory(p); err != nil || !e {
		return nil, fs.ERR_INVALID_DIR
	}

	entries, err := os.ReadDir(p)
	if err != nil {
		return nil, err
	}

	tasks := list.New()

	for _, entry := range entries {
		if entry.IsDir() {
			task, err := CheckResolveTaskFromTasksFolder(entry.Name())
			if err != nil {
				return nil, err
			}

			tasks.PushBack(task)
		} else {
			continue
		}
	}

	return tasks, nil
}

func CheckResolveTaskFromTasksFolder(name string) (*Task, error) {
	cfg := config.FetchBuildConfig()
	task_path := path.Join(cfg.TasksPath, name)
	task_cfg := path.Join(task_path, "config.json")

	if !(fs.IsAccessible(task_path) || fs.IsAccessible(task_cfg)) {
		return nil, ERR_CANNOT_ACCESS
	}

	task, err := CheckResolveTaskFromPath(task_cfg)
	if err != nil {
		return nil, err
	}

	if task.Name != name {
		return nil, ERR_TASK_NAME_NOT_EQUAL
	}

	return task, nil
}

func CheckResolveTaskFromPath(p string) (*Task, error) {
	task, err := ResolveTaskFromPath(p)
	if err != nil {
		return nil, err
	}

	if IsNil(task.ExternalScraper) && *task.ExternalScraper {
		if err = CheckExternalScraperOptions(task); err != nil {
			return nil, err
		}
	}

	if !slices.IncludeInSliceString(TASK_CATEGORIES, task.Category) {
		return nil, ERR_INVALID_CATE
	}

	return task, nil
}

var (
	ERR_ES_CFG_NEED_POLICY    = errors.New("need provide externalScraperOptions.policy as external scraper task")
	ERR_ES_CFG_UNKNOWN_POLICY = errors.New("unknown policy")
	ERR_ES_CFG_NEED_BUILD_REQ = errors.New("need provide buildRequirement as external scraper task with make.cmd")
	ERR_ES_CFG_NEED_REL_REQ   = errors.New("need provide releaseRequirement as external scraper task with make.cmd")
)

func CheckExternalScraperOptions(task *Task) error {
	if task.AutoMake {
		if task.ExternalScraperOptions.Policy == nil {
			return ERR_ES_CFG_NEED_POLICY
		}
		x := *(task.ExternalScraperOptions.Policy) != ES_POLICY_MANUAL
		if x || *(task.ExternalScraperOptions.Policy) != ES_POLICY_SILENT {
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
