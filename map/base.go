package stack

type (
	_typeKey_   int
	_typeValue_ int
)

type _TypeKey__TypeValue_Iterator func(key _typeKey_, value _typeValue_) (cont bool)

type _typeKey__typeValue_Interface interface {
	Len() int
	Get(key _typeKey_) (value _typeValue_, ok bool)
	Put(key _typeKey_, value _typeValue_)
	Del(key _typeKey_)
	Each(f _TypeKey__TypeValue_Iterator) (done bool)
	Keys() []_typeKey_
	Values() []_typeValue_
}
