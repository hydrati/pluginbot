package config

import (
	"fmt"
	"io"

	"github.com/hyroge/pluginbot/utils/json"
	"github.com/hyroge/pluginbot/utils/output"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

type BuildPkgInfo struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}

type BuildRecentStatus struct {
	Time            uint64 `json:"time"`
	TimeDescription string `json:"timeDescription"`
	Success         bool   `json:"success"`
	ErrorMessage    string `json:"errorMessage"`
}

type BuildInfoItem struct {
	LatestVersion string               `json:"latestVersion"`
	Builds        []*BuildPkgInfo      `json:"builds"`
	RecentStatus  []*BuildRecentStatus `json:"recentStatus"`
}

func (i *BuildInfoItem) GetBuildTimes() int {
	return len(i.RecentStatus)
}

func (i *BuildInfoItem) GetFailureTimes() int {
	count := 0
	for _, v := range i.RecentStatus {
		if !v.Success {
			count += 1
		}
	}

	return count
}

func (i *BuildInfoItem) GetSuccessTimes() int {
	return i.GetBuildTimes() - i.GetFailureTimes()
}

func (i *BuildInfoItem) GetBarometerWeather(name string) (string, output.BarometerWeather) {
	b, f := i.GetBuildTimes(), i.GetFailureTimes()

	w := output.BaroByTimes(b-f, f)
	s := w.Print(name, b-f, f)
	return s, w
}

type BuildInfoList map[string]*BuildInfoItem

func (i *BuildInfoList) PrintBarometer() {
	fmt.Println("")
	LogInfo(output.BAROMETER_TITLE)
	for k, v := range *i {
		// LogInfo("Build Info of %s", k)
		s, _ := v.GetBarometerWeather(k)
		LogInfo(s)
	}
}

func UnmarshalBuildInfoList(f io.Reader) (*BuildInfoList, error) {
	LogInfo("[config/package] try to resolve package info")
	LogInfo("[config/package] resolve...")

	var config *BuildInfoList
	return config, json.UnmarshalJsonc(f, &config)
}

func MarshalBuildInfoList(o *BuildInfoList, w io.Writer) error {
	LogInfo("[config/package] marshal pkg info list...")
	return json.MarshalJsonToWriter(o, w)
}
