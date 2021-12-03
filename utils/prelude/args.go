package prelude

import (
	"flag"
)

var args *AppCmdArgs

type AppCmdArgs struct {
	BuildConfig string
	RunTaskName string
	CI          bool
	Force       bool
	Debug       bool
}

func resolveCmdArgs() *AppCmdArgs {
	// LogDebug("[utils/args] try to parse cmd args")
	var config = flag.String("c", "./config.jsonc", `build config, default="./config.jsonc"`)
	var task = flag.String("t", "", `specify to run a Task, input the name of the task, default=no`)
	var ci = flag.Bool("g", false, `CI mode (eg. Github Action)`)
	var force = flag.Bool("f", false, `force mode, Ignore the comparison result with the latest version of the database and force the task to be rebuilt`)
	var debug = flag.Bool("d", false, `debug mode`)

	flag.Parse()

	args := &AppCmdArgs{BuildConfig: *config, RunTaskName: *task, CI: *ci, Debug: *debug, Force: *force}
	// LogDebug("[utils/args] parsed: %+v", args)

	return args
}

func FetchCmdArgs() *AppCmdArgs {
	if args == nil {
		args = resolveCmdArgs()
	}
	return args
}
