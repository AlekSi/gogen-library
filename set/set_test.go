package set

import (
	"math/rand"
	"testing"
)

func Random_Type_() _type_ {
	return _type_(rand.Int31())
}

func test_Type_(t *testing.T, s _type_Interface) {
	len := s.Len()
	if len != 0 {
		t.Errorf("unexpected len = %d", len)
	}

	if s.Has(Random_Type_()) {
		t.Errorf("set is not empty")
	}

	value1 := Random_Type_()
	s.Put(value1)
	len = s.Len()
	if len != 1 {
		t.Errorf("unexpected len = %d", len)
	}

	value2 := Random_Type_()
	s.Put(value2)
	len = s.Len()
	if len != 2 {
		t.Errorf("unexpected len = %d", len)
	}

	if !s.Has(value2) {
		t.Errorf("%v not present in set", value2)
	}

	s.Del(value2)
	if s.Has(value2) {
		t.Errorf("%v present in set", value2)
	}

	len = s.Len()
	if len != 1 {
		t.Errorf("unexpected len = %d", len)
	}
	s.Reset()
	len = s.Len()
	if len != 0 {
		t.Errorf("unexpected len = %d", len)
	}
}

func Test_Type_Set(t *testing.T) {
	s := New_Type_Set(1)
	test_Type_(t, s)
}

func Test_Type_SetTS(t *testing.T) {
	s := New_Type_SetTS(1)
	test_Type_(t, s)
}
