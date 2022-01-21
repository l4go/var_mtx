package var_mtx

import (
	"sync"
)

type refRWMutex struct {
	ref int32
	mtx *sync.RWMutex
}

type VarRWMutex struct {
	mtx     *sync.Mutex
	var_mtx map[interface{}]*refRWMutex
}

func NewVarRWMutex() *VarRWMutex {
	return &VarRWMutex{
		mtx:     &sync.Mutex{},
		var_mtx: map[interface{}]*refRWMutex{},
	}
}

func (km *VarRWMutex) incr_mtx(key interface{}) *sync.RWMutex {
	km.mtx.Lock()
	defer km.mtx.Unlock()

	rm, ok := km.var_mtx[key]
	if ok {
		rm.ref++
		if rm.ref <= 0 {
			panic("too many locks")
		}
	} else {
		rm = &refRWMutex{ref: 1, mtx: &sync.RWMutex{}}
		km.var_mtx[key] = rm
	}
	return rm.mtx
}

func (km *VarRWMutex) decr_mtx(key interface{}) *sync.RWMutex {
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

func (km *VarRWMutex) Lock(key interface{}) {
	mtx := km.incr_mtx(key)
	mtx.Lock()
}

func (km *VarRWMutex) Unlock(key interface{}) {
	mtx := km.decr_mtx(key)
	mtx.Unlock()
}

func (km *VarRWMutex) RLock(key interface{}) {
	mtx := km.incr_mtx(key)
	mtx.RLock()
}

func (km *VarRWMutex) RUnlock(key interface{}) {
	mtx := km.decr_mtx(key)
	mtx.RUnlock()
}
