package notifier

import (
	"container/list"
	"sync"

	// . "github.com/hyroge/pluginbot/utils/prelude"
	"github.com/zyxar/argo/rpc"
)

type NotifierFunc func([]rpc.Event)

type DefaultNotifier struct {
	OnDownloadStart      NotifierFunc
	OnDownloadPause      NotifierFunc
	OnDownloadStop       NotifierFunc
	OnDownloadError      NotifierFunc
	OnDownloadComplete   NotifierFunc
	OnBtDownloadComplete NotifierFunc
}

type NotifierCallbackFunc func(*NotifierEvent)

type NotifierEvent struct {
	Event string
	Gids  []rpc.Event
}

type CallbackNotifier struct {
	mutex  sync.Mutex
	recv   chan *NotifierEvent
	closed bool
	done   chan struct{}
	queue  map[string]*list.List // List[WaitNotifierFunc]
}

func NewCallbackNotifier() *CallbackNotifier {
	n := CallbackNotifier{
		recv:   make(chan *NotifierEvent, 10),
		done:   make(chan struct{}, 1),
		queue:  make(map[string]*list.List),
		closed: false,
		mutex:  sync.Mutex{},
	}

	go n.poll_events()
	return &n
}

type CallbackHandler struct {
	elem *list.Element
}

type NotifierWaitCallbackFunc func(*NotifierEvent) bool
type NotifierWaiterFunc func() *NotifierEvent

func (n *CallbackNotifier) CreateWaiter(evn string, cb NotifierWaitCallbackFunc) NotifierWaiterFunc {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	wg.Add(1)

	ch := make(chan *NotifierEvent, 1)
	h := func(ev *NotifierEvent) {
		if !mutex.TryLock() {
			return
		}
		ok := cb(ev)
		if ok {
			ch <- ev
			wg.Done() //
		}
	}

	w := func() *NotifierEvent {
		handle := n.On(evn, h)
		wg.Wait()
		n.Remove(evn, handle)
		ev := <-ch
		close(ch)
		return ev
	}

	return w
}

func (n *CallbackNotifier) On(ev string, cb NotifierCallbackFunc) *CallbackHandler {
	if n.queue[ev] == nil {
		n.queue[ev] = list.New()
	}
	cbs := n.queue[ev]
	return &CallbackHandler{cbs.PushBack(cb)}
}

func (n *CallbackNotifier) Remove(ev string, cb *CallbackHandler) {
	if n.queue[ev] != nil {
		n.queue[ev].Remove(cb.elem)
	}
}

func (n *CallbackNotifier) Close() {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	n.closed = true
	n.done <- struct{}{}
}

func (n *CallbackNotifier) poll_events() {
	for {
		select {
		case <-n.done:
			return
		case ev, ok := <-n.recv:
			if ok {
				if n.queue[ev.Event] != nil {
					cbs := n.queue[ev.Event]
					for i := cbs.Front(); i != nil; i = i.Next() {
						if cb, ok := i.Value.(NotifierCallbackFunc); ok {
							go cb(ev)
						}
					}
				}
			}
		}
	}
}

func (n *CallbackNotifier) OnDownloadStart(events []rpc.Event) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	if !n.closed {
		n.recv <- &NotifierEvent{"DownloadStart", events}
	}
}

func (n *CallbackNotifier) OnDownloadPause(events []rpc.Event) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	if !n.closed {
		n.recv <- &NotifierEvent{"DownloadPause", events}
	}
}

func (n *CallbackNotifier) OnDownloadStop(events []rpc.Event) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	if !n.closed {
		n.recv <- &NotifierEvent{"DownloadStop", events}
	}
}

func (n *CallbackNotifier) OnDownloadError(events []rpc.Event) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	if !n.closed {
		n.recv <- &NotifierEvent{"DownloadError", events}
	}
}

func (n *CallbackNotifier) OnDownloadComplete(events []rpc.Event) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	if !n.closed {
		n.recv <- &NotifierEvent{"DownloadComplete", events}
	}
}

func (n *CallbackNotifier) OnBtDownloadComplete(events []rpc.Event) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	if !n.closed {
		n.recv <- &NotifierEvent{"BtDownloadComplete", events}
	}
}
