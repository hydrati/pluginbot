package init

import (
	"errors"
	"os"
	"runtime"

	"github.com/hyroge/pluginbot/config"
	"github.com/hyroge/pluginbot/utils/brand"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

func init() {
	StartupInit()
}

func StartupInit() {
	brand.DisplayStartup()
	PrintRuntimeInfo()
	_ = FetchCmdArgs()
	Must(PlatformCheck())
	config.FetchBuildConfig()
	LookupBuildTools()
}

func PrintRuntimeInfo() {
	LogInfo("[init] compiler version: %v, %v", runtime.Version(), runtime.Compiler)
	LogInfo("[init] os platform: %v+%v", runtime.GOOS, runtime.GOARCH)
}

func PlatformCheck() error {
	if runtime.GOOS != "windows" || runtime.GOARCH != "amd64" {
		LogError("[init] need windows+amd64, crashed")
		return errors.New("need windows+amd64")
	}
	return nil
}

func LookupBuildTools() {
	LogInfo("[init] look up build tools...")

	if _, err := os.Stat("./tools/"); os.IsNotExist(err) {
		LogError("[init] not found build tools")
		Must(err)
	}
	LogInfo("[init] found build tools")
	os.Setenv("PATH", os.Getenv("PATH")+";./tools/")
	LogInfo("[init] added to path")
}
