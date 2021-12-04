package prelude

func Must(e error) {
	if e != nil {
		LogError("oops! pluginbot crashed! (by must)")
		panic(e)
	}
}

func MustOk(e bool) {
	if e != true {
		LogError("oops! pluginbot crashed! (by must-ok)")
		panic("val is not true")
	}
}
