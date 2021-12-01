package slices

import (
	"reflect"

	. "github.com/hyroge/pluginbot/utils/prelude"
)

func IncludeInSliceAny(s []Any, e Any) bool {
	for _, v := range s {
		if reflect.DeepEqual(v, e) {
			return true
		}
	}
	return false
}

func IncludeInSliceString(s []string, e string) bool {
	for _, v := range s {
		if reflect.DeepEqual(v, e) {
			return true
		}
	}
	return false
}
