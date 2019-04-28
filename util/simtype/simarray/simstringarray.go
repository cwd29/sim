package simarray

import "reflect"

type simStringArray struct {
	array []string
}

func (array *simStringArray) Src(src interface{}) {
	array.array = src.([]string)
}

func (array *simStringArray) Len() int {
	return len(array.array)
}

func (array *simStringArray) ToArray() interface{} {
	return array.array
}

func (array *simStringArray) Contains(cnt interface{}) bool {
	if reflect.ValueOf(cnt).Kind() != reflect.String {
		return false
	}
	ncnt := cnt.(string)
	for _, value := range array.array {
		if value == ncnt {
			return true
		}
	}
	return false
}

// // SimStringArray is 自定义的string数组类型
// type SimStringArray []string

// // Len is 获取string数组的长度
// func (array SimStringArray) Len() int {
// 	return len(array)
// }

// // Contains is 判断string数组中是否包含某个元素
// func (array SimStringArray) Contains(cnt interface{}) bool {
// 	if reflect.ValueOf(cnt).Type().String() != "string" {
// 		return false
// 	}

// 	for _, element := range array {
// 		if element == cnt.(string) {
// 			return true
// 		}
// 	}
// 	return false
// }

// // Matches is 获取string数组中所有匹配正则的string
// func (array SimStringArray) Matches(format string) (result []string) {
// 	for _, element := range array {
// 		if match, _ := regexp.MatchString(format, element); match {
// 			result = append(result, element)
// 		}
// 	}
// 	return
// }
