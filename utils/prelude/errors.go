package prelude

func Must(e error) {
	if e != nil {
		LogError("oops! pluginbot crashed! (by must)")
		panic(e)
	}
}
