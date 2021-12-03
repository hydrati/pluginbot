package paspider

import (
	. "github.com/hyroge/pluginbot/utils/prelude"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type CreatePageOptions = proto.TargetCreateTarget

const (
	DEBUG = false
)

func LaunchBrowserDefault() (*rod.Browser, error) {
	LogDebug("[pas-launcher] launch browser with default...")
	if DEBUG {
		LogDebug("[pas-launcher] debug mode enabled")
	}
	l, err := launcher.New().Headless(!DEBUG).Launch()
	if err != nil {
		LogError("[pas-launcher] launch error")
		return nil, err
	}
	LogDebug("[pas-launcher] create browser controller...")
	client := rod.New().ControlURL(l)
	LogDebug("[pas-launcher] connect to browser...")
	err = client.Connect()
	if err != nil {
		LogError("[pas-launcher] connect error")
		return nil, err
	}
	LogDebug("[pas-launcher] browser ready")
	return client, nil
}

func MustLaunchBrowserDefault() *rod.Browser {
	r, err := LaunchBrowserDefault()
	Must(err)
	return r
}
