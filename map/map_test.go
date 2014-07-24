package stack

import (
	"math/rand"
	"testing"
)

func Random_TypeKey_() _typeKey_ {
	return _typeKey_(rand.Int31())
}

func Random_TypeValue_() _typeValue_ {
	return _typeValue_(rand.Int31())
}

func testMap(t *testing.T, m _typeKey__typeValue_Interface) {
	len := m.Len()
	if len != 0 {
		t.Errorf("unexpected len = %d", len)
	}

	var zero _typeValue_
	value, ok := m.Get(Random_TypeKey_())
	if ok || value != zero {
		t.Errorf("unexpected %t value = %v", ok, value)
	}

	key1, value1 := Random_TypeKey_(), Random_TypeValue_()
	m.Put(key1, value1)
	len = m.Len()
	if len != 1 {
		t.Errorf("unexpected len = %d", len)
	}

	key2, value2 := Random_TypeKey_(), Random_TypeValue_()
	m.Put(key2, value2)
	len = m.Len()
	if len != 2 {
		t.Errorf("unexpected len = %d", len)
	}

	value, ok = m.Get(key1)
	if !ok || value != value1 {
		t.Errorf("unexpected %t value = %v", ok, value)
	}

	value, ok = m.Get(key2)
	if !ok || value != value2 {
		t.Errorf("unexpected %t value = %v", ok, value)
	}

	m.Del(key2)
	value, ok = m.Get(key2)
	if ok || value != zero {
		t.Errorf("unexpected %t value = %v", ok, value)
	}
}

func Test_TypeKey__TypeValue_MapTS(t *testing.T) {
	m := New_TypeKey__TypeValue_MapTS(1)
	testMap(t, m)
}
