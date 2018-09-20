package chandy_lamport

import "sync"

// An implementation of a map that synchronizes read and write accesses.
// Note: This class intentionally adopts the interface of `sync.Map`,
// which is introduced in Go 1.9+ but not available before that.
// This provides a simplified version of the same class without
// requiring the user to upgrade their Go installation.
type SyncMap struct {
	internalMap map[interface{}]interface{}
	lock sync.RWMutex
}

func NewSyncMap() *SyncMap {
	m := SyncMap{}
	m.internalMap = make(map[interface{}]interface{})
	return &m
}

func (m *SyncMap) Load(key interface{}) (value interface{}, ok bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	value, ok = m.internalMap[key]
	return
}

func (m *SyncMap) Store(key, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.internalMap[key] = value
}

func (m *SyncMap) LoadOrStore(key, value interface{}) (interface{}, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	existingValue, ok := m.internalMap[key]
	if ok {
		return existingValue, true
	}
	m.internalMap[key] = value
	return value, false
}

func (m *SyncMap) Delete(key interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.internalMap, key)
}

func (m *SyncMap) Range(f func(key, value interface{}) bool) {
	m.lock.RLock()
	for k, v := range m.internalMap {
		if !f(k, v) {
			break
		}
	}
	defer m.lock.RUnlock()
}
