package var_mtx

import (
	"sync"
)

type refMutex struct {
	ref int32
	mtx *sync.Mutex
}

type VarMutex struct {
	mtx     *sync.Mutex
	var_mtx map[interface{}]*refMutex
}

func NewVarMutex() *VarMutex {
	return &VarMutex{
		mtx:     &sync.Mutex{},
		var_mtx: map[interface{}]*refMutex{},
	}
}

func (km *VarMutex) incr_mtx(key interface{}) *sync.Mutex {
	km.mtx.Lock()
	defer km.mtx.Unlock()

	rm, ok := km.var_mtx[key]
	if ok {
		rm.ref++
		if rm.ref <= 0 {
			panic("too many locks")
		}
	} else {
		rm = &refMutex{ref: 1, mtx: &sync.Mutex{}}
		km.var_mtx[key] = rm
	}
	return rm.mtx
}

func (km *VarMutex) decr_mtx(key interface{}) *sync.Mutex {
	km.mtx.Lock()
	defer km.mtx.Unlock()

	rm, ok := km.var_mtx[key]
	if !ok {
		panic("not locked")
	}

	mtx := rm.mtx
	rm.ref--
	if rm.ref <= 0 {
		rm.mtx = nil
		delete(km.var_mtx, key)
	}
	return mtx
}

func (km *VarMutex) Lock(key interface{}) {
	mtx := km.incr_mtx(key)
	mtx.Lock()
}

func (km *VarMutex) Unlock(key interface{}) {
	mtx := km.decr_mtx(key)
	mtx.Unlock()
}
