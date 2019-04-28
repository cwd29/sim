package simarray

import (
	"errors"
	"reflect"
	"strings"
)

// SimArray is 指定以的数组类型
type SimArray struct {
	array ISimArray
}

// NewSimArray is 对SimArray进行初始化
func NewSimArray(src interface{}) (array *SimArray, err error) {
	if src == nil {
		return nil, errors.New("传入的数组不能为nil")
	}
	if rtype := reflect.ValueOf(src); rtype.Kind() == reflect.Slice {
		if rtype.Len() == 0 {
			return nil, errors.New("传入的数组长度不能为0")
		}
		array = &SimArray{}
		switch rtype.Index(0).Kind() {
		case reflect.String:
			array.array = &simStringArray{}
			break
		case reflect.Int32:
			array.array = &simIntArray{}
			break
		default:
			return nil, errors.New("对于除了string与int的数组之外的其他类型的数组还未进行处理")
		}
		array.array.Src(src)
		return
	}
	return nil, errors.New("传入必须是数组类型")
}

// Contains is 数组中是否包含某个元素
func (array *SimArray) Contains(cnt interface{}) bool {
	return array.array.Contains(cnt)
}

// ContainsAny is 数组中的元素包含sub的一部分的列表
func (array *SimArray) ContainsAny(sub string) (result []string) {
	for _, value := range array.array.ToArray().([]string) {
		if strings.Contains(value, sub) {
			result = append(result, value)
		}
	}
	return
}

// // ArrayContainsByRefelct is 用于比较一个数组中是否包含一个元素
// func (array *SimArray) ArrayContainsByRefelct(cnt interface{}) bool {
// 	cntType := reflect.ValueOf(cnt).Type()
// 	if array := reflect.ValueOf(array); array.Kind() == reflect.Slice {
// 		for i := 0; i < array.Len(); i++ {
// 			if array.Index(i).Type() != cntType {
// 				return false
// 			}
// 			if reflect.DeepEqual(array.Index(i).Interface(), cnt) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// // SimArrayContains is 一个数组中是否包含有某个元素
// func (array *SimArray) SimArrayContains(cnt interface{}) bool {
// 	if narray := reflect.ValueOf(array); narray.Kind() == reflect.Slice {
// 		switch narray.Index(0).Type().String() {
// 		// case "string":
// 		// 	return SimStringArray(array.([]string)).Contains(cnt)
// 		// case "int":
// 		// 	return SimIntArray(array.([]int)).Contains(cnt)
// 		}
// 	}
// 	return false
// }

// SimArrayMatches is 获取string数组中所有匹配正则的string
// func (array *SimArray) SimArrayMatches(format string) []string {
// 	return SimStringArray(array).Matches(format)
// }
