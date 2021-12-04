package worker

import (
	"container/list"
	"sync"

	. "github.com/hyroge/pluginbot/utils/prelude"
)

type WorkerPool struct {
	wg     *sync.WaitGroup
	mutex  *sync.Mutex
	jobs   *list.List
	cur    *list.Element
	result *list.List
	future *WorkerPoolFuture
	ch     chan *list.List
	count  int
	run    bool
	done   bool
	gots   int
}

func NewPool(count int, jobs *list.List) *WorkerPool {
	p := &WorkerPool{&sync.WaitGroup{}, &sync.Mutex{}, jobs, jobs.Front(), list.New(), nil, nil, count, false, false, 0}
	future, ch := newFuture(p.wg)
	p.future = future
	p.ch = ch

	return p
}

func (p *WorkerPool) lock() {
	p.mutex.Lock()
}

func (p *WorkerPool) unlock() {
	p.mutex.Unlock()
}

func (p *WorkerPool) pull() *list.Element {
	p.lock()
	defer p.unlock()

	job := p.cur
	if job != nil {
		p.gots += 1
		p.cur = p.cur.Next()
	}
	return job
}

func (p *WorkerPool) resolve(result interface{}) {
	p.lock()
	defer p.unlock()

	p.result.PushBack(result)
	if p.result.Len() >= p.jobs.Len() {
		p.ch <- p.result
	}

	p.wg.Done()
}

func (p *WorkerPool) IsDone() bool {
	p.lock()
	defer p.unlock()
	if !p.done {
		p.done = p.gots >= p.jobs.Len()
	}
	return p.done
}

type WorkerFunc func(interface{}) interface{}
type WorkerPoolAwaiter func() (*list.List, bool)

type WorkerPoolFuture struct {
	wg *sync.WaitGroup
	ch chan *list.List
}

func newFuture(wg *sync.WaitGroup) (*WorkerPoolFuture, chan *list.List) {
	ch := make(chan *list.List, 1)
	return &WorkerPoolFuture{wg, ch}, ch
}

func (p *WorkerPoolFuture) GetAwaiter() WorkerPoolAwaiter {
	return func() (*list.List, bool) {
		p.wg.Wait()
		close(p.ch)
		val, ok := <-p.ch
		return val, ok
	}
}

func (p *WorkerPoolFuture) Await() (*list.List, bool) {
	return p.GetAwaiter()()
}

func (p *WorkerPool) Run(f WorkerFunc) *WorkerPoolFuture {
	if !p.run {
		p.wg.Add(p.jobs.Len())
		for i := 0; i < p.count; i += 1 {
			go func(id int) {
				defer func() {
					if err := recover(); err != nil {
						LogError("worker(%d) in pool error: %+v", id, err)
					}
				}()
				for {
					if !p.IsDone() {
						job := p.pull()
						if job != nil {
							p.resolve(f(job.Value))
						}
					} else {
						return
					}
				}
			}(i)
		}
		p.run = true
		return p.future
	} else {
		panic("pool running")
	}
}
