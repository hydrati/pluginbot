package paspider

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	. "github.com/hyroge/pluginbot/utils/prelude"
)

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
