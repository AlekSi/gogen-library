package set

type _Type_Set struct {
	data map[_type_]struct{}
	cap  int
}

// check interface
var _ _type_Interface = New_Type_Set(1)

func New_Type_Set(cap int) *_Type_Set {
	s := &_Type_Set{cap: cap}
	s.Reset()
	return s
}

func (s *_Type_Set) Len() int {
	l := len(s.data)
	return l
}

func (s *_Type_Set) Reset() {
	s.data = make(map[_type_]struct{}, s.cap)
}

func (s *_Type_Set) Has(value _type_) bool {
	_, ok := s.data[value]
	return ok
}

func (s *_Type_Set) Put(value _type_) {
	s.data[value] = _type_Present
}

func (s *_Type_Set) Del(value _type_) {
	delete(s.data, value)
}
