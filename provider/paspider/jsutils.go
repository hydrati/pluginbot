package paspider

import (
	"github.com/go-rod/rod/lib/proto"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

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
