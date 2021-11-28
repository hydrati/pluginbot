package paspider

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

var (
	ERR_NOT_FOUND_MD5            = errors.New("not found vaild md5")
	ERR_NOT_FOUND_PAGE_OBJ       = errors.New("not found page object")
	ERR_NOT_FOUND_LOCAL_DL_TABLE = errors.New("not found localization download table")

	REXP_VERSION_PATTERN     = regexp.MustCompile(`\d+.\d+(.\d+)*`)
	REXP_REDIRECT_DL_PATTERN = regexp.MustCompile(`/downloading/`)
	REXP_MD5_PATTERN         = regexp.MustCompile(`([a-zA-Z0-9]+)(\s?)+$`)
)

type PASpider struct {
	client *rod.Browser
	page   *rod.Page
	opts   proto.TargetCreateTarget
	name   string
}

func New(client *rod.Browser, opts proto.TargetCreateTarget, name string) (*PASpider, error) {
	LogInfo("[pas, %s] creating paspider, %s", name, opts.URL)
	page, err := client.Page(opts)
	if err != nil {
		return nil, err
	}

	return (&PASpider{client, page, opts, name}), nil
}

func MustNew(client *rod.Browser, opts proto.TargetCreateTarget, name string) *PASpider {
	r, err := New(client, opts, name)
	Must(err)
	return r
}

func (p *PASpider) GetLocationOrigin() (string, error) {
	LogInfo("[pas, %s] try to get page location origin", p.name)

	obj, err := p.EvaluateJs(`location`)
	if err != nil {
		return "", err
	}

	e := obj.Value.Get("origin").String()
	LogInfo("[pas, %s] got location origin, %s", p.name, e)

	return e, nil
}

func (p *PASpider) EvaluateJs(js string, args ...interface{}) (*proto.RuntimeRemoteObject, error) {
	page, err := p.GetPage()
	if err != nil {
		return nil, err
	}
	LogInfo("[pas, %s] try to eval js\n\tjs = `%s`\n\targs = `%+v`", p.name, js, args)
	return page.Eval(js, args)
}

func (p *PASpider) ClosePage() error {
	LogInfo("[pas, %s] close page", p.name)
	err := p.page.Close()
	p.page = nil
	return err
}

func (p *PASpider) GetPage() (*rod.Page, error) {
	LogInfo("[pas, %s] try to get page controller", p.name)
	if p.page == nil {
		LogError("[pas, %s] page is not ready", p.name)
		return nil, ERR_NOT_FOUND_PAGE_OBJ
	} else {
		LogInfo("[pas, %s] got page", p.name)
		return p.page, nil
	}
}

func (p *PASpider) MustGetPage() *rod.Page {
	page, err := p.GetPage()
	Must(err)
	return page
}

func (p *PASpider) WaitPageReady() error {
	page, err := p.GetPage()
	if err != nil {
		return err
	}

	LogInfo("[pas, %s] wait page loading...", p.name)
	page.WaitNavigation(proto.PageLifecycleEventNameDOMContentLoaded)()
	LogInfo("[pas, %s] page loaded", p.name)
	return nil
}

func (p *PASpider) MustWaitPageReady() {
	err := p.WaitPageReady()
	Must(err)
}

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

func (p *PASpider) EncodeUrl(url string) (string, error) {
	page, err := p.GetPage()
	if err != nil {
		return "", err
	}

	LogInfo("[pas, %s] try to encode url by js", p.name)
	obj, err := page.Eval("encodeURI(`" + url + "`)")
	if err != nil {
		return "", err
	}

	LogInfo("[pas, %s] ok, encoded url to string", p.name)
	return obj.Value.String(), nil
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
