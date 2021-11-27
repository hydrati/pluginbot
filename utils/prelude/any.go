package prelude

import "reflect"

type Any = interface{}

func IsNil(o Any) bool {
	return reflect.DeepEqual(o, nil)
}
