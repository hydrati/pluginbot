package main

import (
	"fmt"

	Spider "github.com/hyroge/pluginbot/provider/paspider"
	"github.com/hyroge/pluginbot/utils/brand"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

func main() {
	brand.DisplayStartup()

	client := Spider.MustLaunchBrowserDefault()

	options := Spider.CreatePageOptions{
		URL: "https://portableapps.com/apps/internet/firefox_portable",
	}
	lang := "Chinese (Simplified)"

	entry, err := Spider.FetchEntry(client, options, lang, "Firefox")
	Must(err) // panic when error

	fmt.Printf("%+v\n", entry)
}
