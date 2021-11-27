package prelude

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var (
	LOG_INFO  = color.New(color.FgGreen).Add(color.Bold)
	LOG_WARN  = color.New(color.FgYellow).Add(color.Bold)
	LOG_ERROR = color.New(color.FgRed).Add(color.Bold)

	LOG_INFO_HEADER  = LOG_INFO.Sprint("[INFO]")
	LOG_WARN_HEADER  = LOG_WARN.Sprint("[WARN]")
	LOG_ERROR_HEADER = LOG_ERROR.Sprint("[ERRO]")
)

func LogInfo(f string, i ...interface{}) {
	t := time.Now().UTC().Format(time.RFC3339)
	s := fmt.Sprintf(f, i...)
	fmt.Printf("%s %v: %s\n", LOG_INFO_HEADER, t, s)
}

func LogWarn(f string, i ...interface{}) {
	t := time.Now().UTC().Format(time.RFC3339)
	s := fmt.Sprintf(f, i...)
	fmt.Printf("%s %v: %s\n", LOG_WARN_HEADER, t, s)
}

func LogError(f string, i ...interface{}) {
	t := time.Now().UTC().Format(time.RFC3339)
	s := fmt.Sprintf(f, i...)
	fmt.Printf("%s %v: %s\n", LOG_ERROR_HEADER, t, s)
}
