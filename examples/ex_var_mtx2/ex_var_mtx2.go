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

var keyMtx = var_mtx.NewVarRWMutex()

func read_worker(m *task.Mission, key int) {
	defer m.Done()

	keyMtx.RLock(key)
	defer keyMtx.RUnlock(key)
	log.Println("Start read:", key)
	defer log.Println("End read:", key)

	time.Sleep(time.Second)
}

func write_worker(m *task.Mission, key int) {
	defer m.Done()

	keyMtx.Lock(key)
	defer keyMtx.Unlock(key)
	log.Println("Start write:", key)
	defer log.Println("End write:", key)

	time.Sleep(3 * time.Second)
}

func main() {
	m := task.NewMission()
	defer m.Done()

	const KINDS = 3
	for i := 1; i < KINDS*3; i++ {
		go read_worker(m.New(), i%KINDS)
		go write_worker(m.New(), i%KINDS)
		time.Sleep(500 * time.Millisecond)
	}
}
