package paspider

import (
	"fmt"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

func (p *PASpider) GetLocalizationDownloadTable() (*rod.Element, error) {
	page, err := p.GetPage()
	if err != nil {
		return nil, err
	}

	LogInfo("[pas, %s] try to get download links table as object", p.name)
	obj, err := page.Evaluate(&rod.EvalOptions{
		ByValue: false,
		JS:      `() => document.querySelector('.download-links')`,
	})
	if err != nil {
		return nil, err
	}

	if obj.Type == proto.RuntimeRemoteObjectTypeObject && obj.Subtype == proto.RuntimeRemoteObjectSubtypeNull {
		LogInfo("[pas, %s] not found invaild download links table", p.name)
		return nil, ERR_NOT_FOUND_LOCAL_DL_TABLE
	}

	LogInfo("[pas, %s] try to convert to element", p.name)
	return page.ElementFromObject(obj)
}

func (p *PASpider) GetRealDownloadUrl(href string) (string, error) {
	origin, err := p.GetLocationOrigin()
	if err != nil {
		return "", err
	}

	u := origin + strings.Clone(href)
	LogInfo("[pas, %s] original download url: %s", p.name, u)

	r := REXP_REDIRECT_DL_PATTERN.ReplaceAllString(u, `/redirect/`)
	LogInfo("[pas, %s] real download url: %s", p.name, r)

	return p.EncodeUrl(r)
}

func (p *PASpider) GetFirstDownloadUrl() (string, error) {
	dl, err := p.GetDownloadBox()
	if err != nil {
		return "", err
	}
	LogInfo("[pas, %s] try to get first download url", p.name)
	LogInfo("[pas, %s] lookup `a` tag", p.name)
	a, err := dl.Element("a")
	if err != nil {
		return "", err
	}

	LogInfo("[pas, %s] got `a` tag, try to get href", p.name)

	href, err := a.Attribute("href")
	if err != nil {
		return "", err
	}

	LogInfo("[pas, %s] get href: %s, converting", p.name, href)
	return p.GetRealDownloadUrl(*href)
}

func (p *PASpider) GetLocalizationDownloads() (PALocalizationDownloadEntryMap, error) {
	LogInfo("[pas, %s] try to get localization download links", p.name)

	table, err := p.GetLocalizationDownloadTable()
	if err != nil {
		LogError("[pas, %s] get table error", p.name)
		return nil, err
	}

	LogInfo("[pas, %s] get table children...", p.name)
	children, err := table.Elements("tr")
	if err != nil {
		LogError("[pas, %s] get failed", p.name)
		return nil, err
	}

	m := make(PALocalizationDownloadEntryMap)
	for i, child := range children[1:] {
		LogInfo("[pas, %s] try to get table item %d/%d", p.name, i, len(children)-1)
		get_text := func(nth uint) (string, error) {
			n, err := child.Element(fmt.Sprintf(`td:nth-child(%d)`, nth))
			if err != nil {
				LogError("[pas, %s] get table item td error", p.name)
				return "", err
			}
			return n.Text()
		}

		lang, err := get_text(1)
		if err != nil {
			LogError("[pas, %s] get lang error", p.name)
			return nil, err
		}

		hash, err := get_text(4)
		if err != nil {
			LogError("[pas, %s] get hash error", p.name)
			return nil, err
		}

		a, err := child.Element(`td:nth-child(3) > a`)
		if err != nil {
			LogError("[pas, %s] get download link node error", p.name)
			return nil, err
		}

		r, err := a.Attribute("href")
		if err != nil {
			LogError("[pas, %s] get download link error", p.name)
			return nil, err
		}

		link, err := p.GetRealDownloadUrl(*r)
		if err != nil {
			LogError("[pas, %s] get real download link error", p.name)
			return nil, err
		}

		m[lang] = &PALocalizationDownloadEntry{Lang: lang, Link: link, Hash: hash}
		LogInfo("[pas, %s] got table item %d, %+v", p.name, i, lang)
	}

	LogInfo("[pas, %s] got localization download links entries", p.name)
	return m, nil
}
