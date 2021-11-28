package brand

import (
	"fmt"

	"github.com/fatih/color"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

var (
	STARTUP_LOGO_STYLE = color.New(color.FgCyan).Add(color.Bold)
	STARTUP_LOGO       = `
 ▄▄▄▄▄▄▄ ▄▄▄     ▄▄   ▄▄ ▄▄▄▄▄▄▄ ▄▄▄ ▄▄    ▄ ▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄ 
█       █   █   █  █ █  █       █   █  █  █ █  ▄    █       █       █
█    ▄  █   █   █  █ █  █   ▄▄▄▄█   █   █▄█ █ █▄█   █   ▄   █▄     ▄█
█   █▄█ █   █   █  █▄█  █  █  ▄▄█   █       █       █  █ █  █ █   █  
█    ▄▄▄█   █▄▄▄█       █  █ █  █   █  ▄    █  ▄   ██  █▄█  █ █   █  
█   █   █       █       █  █▄▄█ █   █ █ █   █ █▄█   █       █ █   █  
█▄▄▄█   █▄▄▄▄▄▄▄█▄▄▄▄▄▄▄█▄▄▄▄▄▄▄█▄▄▄█▄█  █▄▄█▄▄▄▄▄▄▄█▄▄▄▄▄▄▄█ █▄▄▄█  
	`
	TITLE_STYLE = color.New(color.FgCyan).Add(color.Bold)
	VERSION     = `0.0.0-dev`

	AUTHOR_STYLE = color.New(color.FgHiBlack).Add(color.Italic)
	AUTHOR       = `By Hydrogen & Edgeless Team`
)

func DisplayStartup() {
	STARTUP_LOGO_STYLE.Println(STARTUP_LOGO)

	TITLE_STYLE.Printf(`Edgeless Pluginbot`)
	fmt.Printf(" v%s    ", VERSION)

	AUTHOR_STYLE.Println(AUTHOR)
	fmt.Printf("\n")
	LogInfo("[init] startup")
}
