package simarray

// ISimArray is 自定义的一个数组的类型
type ISimArray interface {
	Src(src interface{})
	Len() int
	ToArray() interface{}

	Contains(src interface{}) bool
}
