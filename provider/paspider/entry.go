package paspider

import (
	"github.com/go-rod/rod"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

func FetchEntry(client *rod.Browser, opts CreatePageOptions, def_lang string, name string) (*SpiderResult, error) {
	LogDebug("[pas, %s] start spider task", name)
	spider, err := New(client, opts, name)
	if err != nil {
		LogError("[pas, %s] create spider error", name)
		return nil, err
	}

	defer spider.ClosePage()

	err = spider.WaitPageReady()
	if err != nil {
		LogError("[pas, %s] wait page error", name)
		return nil, err
	}

	entry := &SpiderResult{}

	ver, err := spider.GetVersion()
	if err != nil {
		LogError("[pas, %s] get version error", name)
		return nil, err
	}
	entry.Version = ver

	m, err := spider.GetLocalizationDownloads()

	if err == nil {
		LogDebug("[pas, %s] get localization downloads success", name)
		if m[def_lang] != nil {
			LogDebug("[pas, %s] got default lang, \n%+v", name, m[def_lang])
			d := m[def_lang]
			entry.Hash = d.Hash
			entry.Url = d.Link
			LogDebug("[pas, %s] got entry, return\n%+v", name, entry)
			return entry, nil
		} else {
			LogDebug("[pas, %s] not found default lang link, fallback to first-url", name)
		}
	}

	f, err := spider.GetFirstMD5()
	if err != nil {
		LogError("[pas, %s] get first md5 error", name)
		return nil, err
	}
	LogDebug("[pas, %s] got first md5 hash", name)
	entry.Hash = f

	u, err := spider.GetFirstDownloadUrl()
	if err != nil {
		LogError("[pas, %s] get first download url error", name)
		return nil, err
	}
	LogDebug("[pas, %s] got first download url hash", name)
	entry.Url = u

	LogDebug("[pas, %s] got entry, return\n%+v", name, entry)
	return entry, nil
}
