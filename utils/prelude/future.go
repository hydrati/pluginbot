package prelude

type WaitFunc func(chan None)
type None = struct{}

func AsyncWait(f WaitFunc) chan None {
	ch := make(chan None, 1)
	go (func() {
		f(ch)
		close(ch)
	})()
	return ch
}
