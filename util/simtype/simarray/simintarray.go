package simarray

type simIntArray struct {
	array []int
}

func (array *simIntArray) New(src interface{}) {
	array.array = src.([]int)
}

func (array *simIntArray) Len() int {
	return len(array.array)
}

func (array *simIntArray) ToArray() interface{} {
	return array.array
}

// // SimIntArray is 自定的int数组类型
// type SimIntArray []int

// // Len is 获取int数组的长度
// func (array SimIntArray) Len() int {
// 	return len(array)
// }

// // Contains is 判断int数组中是否包含某个元素
// func (array SimIntArray) Contains(cnt interface{}) bool {
// 	if reflect.ValueOf(cnt).Type().String() != "int" {
// 		return false
// 	}

// 	for _, element := range array {
// 		if element == cnt.(int) {
// 			return true
// 		}
// 	}
// 	return false
// }
