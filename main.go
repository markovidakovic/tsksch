package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Duration time.Duration
}

func NewTask(id int, duration time.Duration) *Task {
	return &Task{
		ID:       id,
		Duration: duration,
	}
}

type Scheduler struct {
	tasks []*Task
	wg    sync.WaitGroup
	mtx   sync.Mutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) AddTask(t *Task) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.tasks = append(s.tasks, t)
}

func (s *Scheduler) Run() {
	for _, task := range s.tasks {
		s.wg.Add(1)
		go func(t *Task) {
			defer s.wg.Done()

			fmt.Printf("Task %d started\n", t.ID)
			time.Sleep(t.Duration)
			fmt.Printf("Task %d finished\n", t.ID)
		}(task)
	}
	s.wg.Wait()
}

func main() {
	sch := NewScheduler()

	sch.AddTask(NewTask(1, 3*time.Second))
	sch.AddTask(NewTask(2, 5*time.Second))
	sch.AddTask(NewTask(3, 2*time.Second))

	sch.Run()
}
