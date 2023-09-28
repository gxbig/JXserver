package threadPool

import (
	"errors"
	"github.com/name5566/leaf/log"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// ErrInvalidPoolCap return if pool size <= 0
	ErrInvalidPoolCap = errors.New("invalid pool cap")
	// ErrPoolAlreadyClosed put task but pool already closed
	ErrPoolAlreadyClosed = errors.New("pool already closed")
	Pools                *Pool
)

const (
	// RUNNING pool is running
	RUNNING = 1
	// STOPED pool is stoped
	STOPED = 0
)

// Task task to-do
type Task struct {
	Handler func(v ...interface{})
	Params  []interface{}
}

// Pool task pool
type Pool struct {
	capacity       uint64
	runningWorkers uint64
	state          int64
	taskC          chan *Task
	PanicHandler   func(interface{})
	sync.Mutex
}

func InitPool(capacity uint64) {
	pools, err := NewPool(capacity)
	log.Release("初始化协程池Pool成功!")
	if err != nil {
		panic(err)
	}
	Pools = pools
}
func GetPool() *Pool {
	return Pools
}

// NewPool init pool
func NewPool(capacity uint64) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrInvalidPoolCap
	}
	return &Pool{
		capacity: capacity,
		state:    RUNNING,
		taskC:    make(chan *Task, capacity),
	}, nil
}

// GetCap get capacity
func (p *Pool) GetCap() uint64 {
	return p.capacity
}

// GetRunningWorkers get running workers
func (p *Pool) GetRunningWorkers() uint64 {
	return atomic.LoadUint64(&p.runningWorkers)
}
func (p *Pool) incRunning() {
	atomic.AddUint64(&p.runningWorkers, 1)
}
func (p *Pool) decRunning() {
	atomic.AddUint64(&p.runningWorkers, ^uint64(0))
}
func (p *Pool) WaitRun(Handler func(v ...interface{}), Params []interface{}) error {
	task := &Task{
		Handler: Handler,
		Params:  Params,
	}
	err := p.Put(task)
	if errors.Is(err, ErrPoolAlreadyClosed) {
		InitPool(Pools.capacity)
		err := Pools.Put(task)
		if err != nil {
			return err
		}
	}
	return nil
}

// Put put a task to pool
func (p *Pool) Put(task *Task) error {
	if p.getState() == STOPED {
		return ErrPoolAlreadyClosed
	}
	// safe run worker
	p.Lock()
	if p.GetRunningWorkers() < p.GetCap() {
		p.run()
	}
	p.Unlock()
	// send task safe
	p.Lock()
	if p.state == RUNNING {
		p.taskC <- task
	}
	p.Unlock()
	return nil
}
func (p *Pool) run() {
	p.incRunning()
	go func() {
		defer func() {
			p.decRunning()
			if r := recover(); r != nil {
				p.run()
				if p.PanicHandler != nil {
					p.PanicHandler(r)
				} else {
					log.Release("Worker panic: %s\n", r)

				}
			}
		}()
		for {
			select {
			case task, ok := <-p.taskC:
				if !ok {
					return
				}
				task.Handler(task.Params...)
			}
		}
	}()
}
func (p *Pool) getState() int64 {
	p.Lock()
	defer p.Unlock()
	return p.state
}
func (p *Pool) setState(state int64) {
	p.Lock()
	defer p.Unlock()
	p.state = state
}

// close safe
func (p *Pool) close() {
	p.Lock()
	defer p.Unlock()
	close(p.taskC)
}

// Close close pool graceful
func (p *Pool) Close() {
	if p.getState() == STOPED {
		return
	}
	p.setState(STOPED)     // stop put task
	for len(p.taskC) > 0 { // wait all task be consumed
		time.Sleep(1e6) // reduce CPU load
	}
	p.close()
}
