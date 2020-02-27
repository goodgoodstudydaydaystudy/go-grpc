package mysql

import "fmt"

type task struct {
	f func() error
}

func NewTask(f func() error) *task {
	return &task{
		f: f,
	}
}

// 执行任务
func (t *task) Execute()  {
	t.f()
}


type pool struct {
	entryChannel chan *task
	poolCap int
	jobsChannel  chan *task
}

func NewPool(cap int) *pool {
	return &pool{
		entryChannel: make(chan *task),
		poolCap:      cap,
		jobsChannel:  make(chan *task),
	}
}

func (p *pool) worker(id int) {
	for task := range p.jobsChannel {
		task.Execute()
		fmt.Println("worker id:", id)
	}
}

func (p *pool)runPool() {
	for i := 0; i < p.poolCap; i++ {
		go p.worker(i)
	}

	for task := range p.entryChannel  {
		p.jobsChannel <- task
	}

	close(p.jobsChannel)

	close(p.entryChannel)

}