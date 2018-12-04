package utils

import (
	"time"
)

func NewTask(c int, fn func(p ... interface{})) *Task {
	return &Task{c: c, q: make(chan int, c*20), handle: fn}
}

type Task struct {
	c      int
	q      chan int
	handle func(ps ... interface{})
}

func (t *Task) Do(i ... interface{}) {
	t.q <- 1

	for {
		if len(t.q) > t.c {
			time.Sleep(time.Millisecond * 200)
			continue
		}

		go func() {
			t1 := time.Now().UnixNano()
			t.handle(i...)
			<-t.q
			t2 := (time.Now().UnixNano() - t1) / 1000000
			if t2 < 200 {
				t.c += 1
			} else {
				t.c -= 1
			}
		}()
		return
	}
}
