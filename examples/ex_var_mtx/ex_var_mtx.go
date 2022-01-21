package main

import (
	"log"
	"time"

	"github.com/l4go/task"
	"github.com/l4go/var_mtx"
)

type Test struct {
	Id int
}

var keyMtx = var_mtx.NewVarMutex()

func worker(m *task.Mission, id int, key int) {
	defer m.Done()

	keyMtx.Lock(key)
	defer keyMtx.Unlock(key)
	log.Println("Start:", id, key)
	defer log.Println("End:", id, key)

	time.Sleep(3 * time.Second)
}

func main() {
	m := task.NewMission()
	defer m.Done()

	for i := 1; i < 10; i++ {
		go worker(m.New(), i, i%3)
		time.Sleep(500 * time.Millisecond)
	}
}
