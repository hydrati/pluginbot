package main

import (
	"fmt"
	"os"
	"sync"

	Spider "github.com/hyroge/pluginbot/provider/paspider"
	"github.com/hyroge/pluginbot/utils/brand"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

func init() {
	brand.DisplayStartup()
	LogInfo("[init] look up build tools...")

	if _, err := os.Stat("./tools/"); os.IsNotExist(err) {
		LogError("[init] not found build tools")
		Must(err)
	}
	LogInfo("[init] found build tools")
	os.Setenv("PATH", os.Getenv("PATH")+";./tools/")

}

func main() {
	client := Spider.MustLaunchBrowserDefault()
	group := &sync.WaitGroup{}
	group.Add(2)

	results := make(chan *Spider.PAEntry, 2)

	worker := func(url string, name string) {
		defer group.Done()

		options := Spider.CreatePageOptions{
			URL: url,
		}
		lang := "Chinese (Simplified)"

		entry, err := Spider.FetchEntry(client, options, lang, name)
		Must(err) // panic when error

		fmt.Printf("got entry: %+v\n", entry)
		results <- entry

	}

	go worker("https://portableapps.com/apps/internet/firefox_portable", "firefox")
	go worker("https://portableapps.com/apps/development/cppcheck-portable", "cppcheck")
	group.Wait()
	close(results)

	fmt.Println()
	LogInfo("[main] tasks done")
	fmt.Println("=========================")

	for entry := range results {
		LogInfo("got %+v", entry)
	}

	// a, err := unarr.NewArchive("test.7z")
	// if err != nil {
	// 	panic(err)
	// }
	// LogInfo("%+v", a.List())
	// defer a.Close()

	// archive, err := lzmadec.NewArchive("test.7z")
	// Must(err)

	// // list all files inside archive
	// for _, e := range archive.Entries {
	// 	fmt.Printf("name: %s, size: %d\n", e.Path, e.Size)
	// 	fmt.Println(e.Attributes)
	// }
	// firstFile := archive.Entries[0].Path
	// fmt.Println(firstFile)
}
