package paspider

import (
	"regexp"

	"github.com/go-rod/rod"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

func (p *PASpider) GetAppTitle() (string, error) {
	page, err := p.GetPage()
	if err != nil {
		return "", err
	}
	LogInfo("[pas, %s] get page title...", p.name)
	tnode, err := page.Element("#page-title")
	if err != nil {
		return "", err
	}

	LogInfo("[pas, %s] get page title text...", p.name)
	text, err := tnode.Text()
	if err != nil {
		return "", err
	}

	LogInfo("[pas, %s] got page title text", p.name)
	return text, nil
}

func (p *PASpider) GetDownloadBox() (*rod.Element, error) {
	page, err := p.GetPage()
	if err != nil {
		return nil, err
	}
	LogInfo("[pas, %s] get default download box", p.name)
	return page.Element(".download-box")
}

func (p *PASpider) GetFirstMD5() (string, error) {
	page, err := p.GetPage()
	if err != nil {
		return "", err
	}

	LogInfo("[pas, %s] try to get first md5", p.name)
	var htag *rod.Element
	for _, elem := range page.MustElements(`strong`) {
		if m, err := regexp.MatchString(`(?i)MD5 Hash`, elem.MustText()); err == nil && m {
			LogInfo("[pas, %s] found a md5 hash tag, %+v", p.name, elem)
			htag = elem
			break
		}
	}

	var text string = ""
	if htag != nil {
		tp, err := htag.Parent()
		if err != nil {
			LogError("[pas, %s] md5 found, get parent error", p.name)
			return "", err
		}
		text0, err := tp.Text()
		if err != nil {
			LogError("[pas, %s] md5 found, get text error", p.name)
			return "", err
		}
		h := REXP_MD5_PATTERN.FindStringSubmatch(text0)
		LogInfo("[pas, %s] found, try to match hash string, %+v", p.name, h)
		if len(h) > 1 {
			text = h[1]
		}
	} else {
		LogInfo("[pas, %s] not found first md5", p.name)
	}

	if len(text) != 32 {
		LogError("[pas, %s] not found md5 hash string, %s", p.name)
		return "", ERR_NOT_FOUND_MD5
	} else {
		return text, nil
	}
}

func (p *PASpider) GetVersion() (string, error) {
	page, err := p.GetPage()
	if err != nil {
		LogError("[pas, %s] get page error", p.name)
		return "", err
	}

	LogInfo("[pas, %s] try to get app version", p.name)

	dlinfon, err := page.Element("p.download-info")
	if err != nil {
		LogError("[pas, %s] get download info error", p.name)
		return "", err
	}

	dlinfo, err := dlinfon.Text()
	if err != nil {
		LogError("[pas, %s] get download info text error", p.name)
		return "", err
	}

	ver := REXP_VERSION_PATTERN.FindString(dlinfo)
	LogInfo("[pas, %s] found app version, %s", p.name, ver)

	return ver, nil
}
