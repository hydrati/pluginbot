package main

import (
	"fmt"
	"os"

	"github.com/hyroge/pluginbot/build/task"
	"github.com/hyroge/pluginbot/config"
	Spider "github.com/hyroge/pluginbot/provider/paspider"
	"github.com/hyroge/pluginbot/utils/aria2"
	"github.com/hyroge/pluginbot/utils/json"
	"github.com/hyroge/pluginbot/utils/output"
	"github.com/hyroge/pluginbot/utils/slices"
	"github.com/zyxar/argo/rpc"

	_ "github.com/hyroge/pluginbot/utils/init"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

type A struct {
	A *int
	B *int
	C *int
}

func main() {
	_ = Spider.MustLaunchBrowserDefault()
	LogInfo("[main] browser ready")
	LogInfo("%+v", config.FetchBuildConfig())
	LogInfo("%s", output.BaroPrintByTimes("Firefox", 15, 0))

	cfg := config.FetchBuildConfig()

	f, err := os.Open(cfg.BuildInfoPath)
	Must(err)
	defer f.Close()

	db, err := config.UnmarshalBuildInfoList(f)
	Must(err)

	(*db)["Rufus"].PushRecentStatus(&config.BuildRecentStatus{
		Time:            1,
		TimeDescription: "114514",
		Success:         false,
		ErrorMessage:    "homo",
	}, 3)

	db.PrintBarometer()
	// fs.CopyDirRecursive("vendor", "vendor2", false)
	// _, err = fs.ReadDirRecursive("vendor")
	// Must(err)

	fmt.Println(json.MarshalJsonToString(db))
	s := []string{"a", "ss", "bb"}
	LogInfo("%+v", slices.IncludeInSliceString(s, "ss"))
	fmt.Println(task.CheckResolveTaskFromPath("./tests/example.pa-task.json"))
	client, err := aria2.NewClient(aria2.RpcOptions{
		Host:      "localhost",
		Port:      6800,
		Secret:    "edgeless",
		Transport: "ws",
		Timeout:   "1s",
		Notifier:  rpc.DummyNotifier{},
	})

	Must(err)

	guard, err := client.GetClient()
	Must(err)
	defer guard.Close()

	ver, err := guard.Get().GetVersion()
	Must(err)
	LogInfo("%+v", ver)

	// group := &sync.WaitGroup{}
	// group.Add(2)

	// results := make(chan *Spider.PAEntry, 2)

	// worker := func(url string, name string) {
	// 	defer group.Done()

	// 	options := Spider.CreatePageOptions{
	// 		URL: url,
	// 	}
	// 	lang := "Chinese (Simplified)"

	// 	entry, err := Spider.FetchEntry(client, options, lang, name)
	// 	Must(err) // panic when error

	// 	fmt.Printf("got entry: %+v\n", entry)
	// 	results <- entry

	// }

	// go worker("https://portableapps.com/apps/internet/firefox_portable", "firefox")
	// go worker("https://portableapps.com/apps/development/cppcheck-portable", "cppcheck")
	// group.Wait()
	// close(results)

	// fmt.Println()
	// LogInfo("[main] tasks done")
	// fmt.Println("=========================")

	// for entry := range results {
	// 	LogInfo("got %+v", entry)
	// }

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
