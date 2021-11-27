package prelude

import (
	"github.com/pkg/errors"
)

func Must(e error) {
	if e != nil {
		LogError("oops! pluginbot crashed!")
		panic(errors.Wrap(e, "panic by must()"))
	}
}
