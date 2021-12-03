package main

import (
	"time"

	"github.com/hyroge/pluginbot/build/task"
	"github.com/hyroge/pluginbot/config"
	"github.com/hyroge/pluginbot/provider/paspider"
	"github.com/hyroge/pluginbot/utils/output"
	"github.com/hyroge/pluginbot/utils/worker"

	_ "github.com/hyroge/pluginbot/utils/init"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

type A struct {
	A *int
	B *int
	C *int
}

func main() {
	start_time := time.Now().Unix()
	client := paspider.MustLaunchBrowserDefault()
	LogInfo("started browser")
	LogDebug("[main] browser ready")
	LogDebug("%+v", config.FetchBuildConfig())
	LogDebug("%s", output.BaroPrintByTimes("Firefox", 15, 0))

	//cfg := config.FetchBuildConfig()

	tasks, err := task.CheckResolveAllTaskInTasksFolder()
	LogInfo("resolved all tasks")
	// wg := sync.WaitGroup{}
	Must(err)

	LogInfo("build with 64 threads...")
	pool := worker.NewPool(64, tasks)
	wait := pool.Run(func(job interface{}) interface{} {
		task := job.(*config.Task)
		if task.PAUrl == nil {
			return nil
		}
		pa_entry, err := paspider.FetchEntry(client, paspider.CreatePageOptions{
			URL: *task.PAUrl,
		}, "Chinese (Simplified)", task.Name)
		Must(err)
		LogDebug("[pa, %s] %+v", task.Name, pa_entry)
		return pa_entry
	})
	result := wait()
	for i := result.Front(); i != nil; i = i.Next() {
		LogInfo("Got, %+v", i.Value)
	}
	LogInfo("Ok, %d/%d", result.Len(), tasks.Len())
	end_time := time.Now().Unix()

	LogDebug("used %d s", (end_time - start_time))

	// for i := tasks.Front(); i != nil; i = i.Next() {
	// 	task := i.Value.(*config.Task)
	// 	if task.PAUrl != nil {
	// 		wg.Add(1)
	// 		go func() {
	// 			pa_entry, err := paspider.FetchEntry(client, paspider.CreatePageOptions{
	// 				URL: *task.PAUrl,
	// 			}, "Chinese (Simplified)", task.Name)
	// 			Must(err)
	// 			LogDebug("[pa, %s] %+v", task.Name, pa_entry)
	// 			wg.Done()
	// 		}()
	// 	}
	// }

	// wg.Wait()

	// f, err := os.Open(cfg.BuildInfoPath)
	// Must(err)
	// defer f.Close()

	// db, err := config.UnmarshalBuildInfoList(f)
	// Must(err)
	// db.PrintBarometer()
	// // fs.CopyDirRecursive("vendor", "vendor2", false)
	// // _, err = fs.ReadDirRecursive("vendor")
	// // Must(err)

	// fmt.Println(json.MarshalJsonToString(db))
	// s := []string{"a", "ss", "bb"}
	// LogDebug("%+v", slices.IncludeInSliceString(s, "ss"))
	// fmt.Println(task.CheckResolveTaskFromPath("./tests/example.pa-task.json"))

	// cmd := aria2.NewCmd("tools/aria2c.exe", cfg.Aria2SpawnArgs)
	// Must(cmd.Start())

	// handle := notifier.NewCallbackNotifier()
	// client, err := aria2.NewClient(aria2.RpcOptions{
	// 	Host:      cfg.Aria2Host,
	// 	Port:      cfg.Aria2Port,
	// 	Secret:    cfg.Aria2Secret,
	// 	Transport: "ws",
	// 	Timeout:   "1s",
	// 	Notifier:  handle,
	// })
	// Must(err)

	// w := handle.CreateWaiter("DownloadStart", func(ev *notifier.NotifierEvent) bool {
	// 	fmt.Printf("hhh: %+v\n", ev)
	// 	return true
	// })

	// go func() {
	// 	guard, err := client.GetClient()
	// 	Must(err)
	// 	defer guard.Close() // unlock client!

	// 	ver, err := guard.Get().GetVersion()
	// 	Must(err)
	// 	LogDebug("%+v", ver)

	// 	_, err = guard.Get().AddURI([]string{"https://zfile.edgeless.top/s/ub3caa"})
	// 	Must(err)
	// }()

	// fmt.Println(w())

	// defer Must(client.Close()) // close rpc client
	// defer Must(cmd.Close())    // close process at last

	// group := &sync.WaitGroup{}
	// group.Add(2)

	// results := make(chan *paspider.PAEntry, 2)

	// worker := func(url string, name string) {
	// 	defer group.Done()

	// 	options := paspider.CreatePageOptions{
	// 		URL: url,
	// 	}
	// 	lang := "Chinese (Simplified)"

	// 	entry, err := paspider.FetchEntry(client, options, lang, name)
	// 	Must(err) // panic when error

	// 	fmt.Printf("got entry: %+v\n", entry)
	// 	results <- entry

	// }

	// go worker("https://portableapps.com/apps/internet/firefox_portable", "firefox")
	// go worker("https://portableapps.com/apps/development/cppcheck-portable", "cppcheck")
	// group.Wait()
	// close(results)

	// fmt.Println()
	// LogDebug("[main] tasks done")
	// fmt.Println("=========================")

	// for entry := range results {
	// 	LogDebug("got %+v", entry)
	// }

	// a, err := unarr.NewArchive("test.7z")
	// if err != nil {
	// 	panic(err)
	// }
	// LogDebug("%+v", a.List())
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
