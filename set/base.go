package set

type _type_ int

var _type_Present = struct{}{}

type _type_Interface interface {
	Len() int
	Reset()
	Has(value _type_) bool
	Put(value _type_)
	Del(value _type_)
}
