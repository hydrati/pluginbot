package paspider

import (
	"errors"
	"regexp"

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
