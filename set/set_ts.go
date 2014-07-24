package set

import (
	"sync"
)

type _Type_SetTS struct {
	l    sync.RWMutex
	data map[_type_]struct{}
	cap  int
}

// check interface
var _ _type_Interface = New_Type_SetTS(1)

func New_Type_SetTS(cap int) *_Type_SetTS {
	s := &_Type_SetTS{cap: cap}
	s.Reset()
	return s
}

func (s *_Type_SetTS) Len() int {
	s.l.RLock()
	l := len(s.data)
	s.l.RUnlock()
	return l
}

func (s *_Type_SetTS) Reset() {
	s.l.Lock()
	s.data = make(map[_type_]struct{}, s.cap)
	s.l.Unlock()
}

func (s *_Type_SetTS) Has(value _type_) bool {
	s.l.RLock()
	_, ok := s.data[value]
	s.l.RUnlock()
	return ok
}

func (s *_Type_SetTS) Put(value _type_) {
	s.l.Lock()
	s.data[value] = _type_Present
	s.l.Unlock()
}

func (s *_Type_SetTS) Del(value _type_) {
	s.l.Lock()
	delete(s.data, value)
	s.l.Unlock()
}
