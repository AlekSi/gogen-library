package stack

import (
	"sync"
)

type _TypeKey__TypeValue_MapTS struct {
	l    sync.RWMutex
	data map[_typeKey_]_typeValue_
}

// check interface
var _ _typeKey__typeValue_Interface = New_TypeKey__TypeValue_MapTS(1)

func New_TypeKey__TypeValue_MapTS(cap int) *_TypeKey__TypeValue_MapTS {
	return &_TypeKey__TypeValue_MapTS{
		data: make(map[_typeKey_]_typeValue_, cap),
	}
}

func (m *_TypeKey__TypeValue_MapTS) Len() int {
	m.l.RLock()
	l := len(m.data)
	m.l.RUnlock()
	return l
}

func (m *_TypeKey__TypeValue_MapTS) Get(key _typeKey_) (value _typeValue_, ok bool) {
	m.l.RLock()
	value, ok = m.data[key]
	m.l.RUnlock()
	return
}

func (m *_TypeKey__TypeValue_MapTS) Put(key _typeKey_, value _typeValue_) {
	m.l.Lock()
	m.data[key] = value
	m.l.Unlock()
	return
}

func (m *_TypeKey__TypeValue_MapTS) Del(key _typeKey_) {
	m.l.Lock()
	delete(m.data, key)
	m.l.Unlock()
	return
}

func (m *_TypeKey__TypeValue_MapTS) Each(f _TypeKey__TypeValue_Iterator) (done bool) {
	m.l.Lock()
	defer m.l.Unlock()

	for k, v := range m.data {
		if !f(k, v) {
			return false
		}
	}
	return true
}

func (m *_TypeKey__TypeValue_MapTS) Keys() []_typeKey_ {
	m.l.RLock()
	keys := make([]_typeKey_, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	m.l.RUnlock()
	return keys
}

func (m *_TypeKey__TypeValue_MapTS) Values() []_typeValue_ {
	m.l.RLock()
	values := make([]_typeValue_, 0, len(m.data))
	for _, v := range m.data {
		values = append(values, v)
	}
	m.l.RUnlock()
	return values

}
