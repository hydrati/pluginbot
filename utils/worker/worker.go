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
	count  int
	run    bool
	done   bool
	gots   int
}

func NewPool(count int, jobs *list.List) *WorkerPool {
	return &WorkerPool{&sync.WaitGroup{}, &sync.Mutex{}, jobs, jobs.Front(), list.New(), count, false, false, 0}
}

func (p *WorkerPool) Lock() {
	p.mutex.Lock()
}

func (p *WorkerPool) Unlock() {
	p.mutex.Unlock()
}

func (p *WorkerPool) Pull() *list.Element {
	p.Lock()
	defer p.Unlock()

	job := p.cur
	if job != nil {
		p.gots += 1
		p.cur = p.cur.Next()
	}
	return job
}

func (p *WorkerPool) Done(result interface{}) {
	p.Lock()
	defer p.Unlock()

	p.result.PushBack(result)

	p.wg.Done()
}

func (p *WorkerPool) IsDone() bool {
	p.Lock()
	defer p.Unlock()
	if !p.done {
		p.done = p.gots >= p.jobs.Len()
	}
	return p.done
}

type WorkerFunc func(interface{}) interface{}

// returns result waiter
func (p *WorkerPool) Run(f WorkerFunc) func() *list.List {
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
						job := p.Pull()
						if job != nil {
							p.Done(f(job.Value))
						}
					} else {
						return
					}
				}
			}(i)
		}
		p.run = true
		return func() *list.List {
			p.wg.Wait()
			return p.result
		}
	} else {
		panic("pool running")
	}
}
