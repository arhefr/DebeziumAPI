package safemap

import (
	"errors"
	"sync"
)

type SafeMap struct {
	mu   *sync.Mutex
	mapa map[any]any
}

func NewSafeMap() *SafeMap {
	mapa := map[any]any{}
	return &SafeMap{mu: &sync.Mutex{}, mapa: mapa}
}

func (sm *SafeMap) Create(key any, val any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.mapa[key] = val
}

func (sm *SafeMap) Read(key any) (any, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	val, ok := sm.mapa[key]
	return val, ok
}

func (sm *SafeMap) ReadAll() []any {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	var reads []any
	for _, val := range sm.mapa {
		reads = append(reads, val)
	}
	return reads
}

func (sm *SafeMap) Update(key any, val any) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if _, ok := sm.mapa[key]; !ok {
		return errors.New("")
	}
	sm.mapa[key] = val

	return nil
}

func (sm *SafeMap) Delete(key any) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if _, ok := sm.mapa[key]; !ok {
		return errors.New("")
	}
	delete(sm.mapa, key)

	return nil
}
