package simarray

import (
	"reflect"
	"testing"
)

func Test_simIntArray_Src(t *testing.T) {
	array := simIntArray{}
	if array.array != nil {
		t.Fail()
	}

	src := []int{1, 2, 3}
	array.Src(src)

	if array.array == nil {
		t.Fail()
	}

	for _, value := range array.array {
		flag := false
		for _, s := range src {
			if s == value {
				flag = true
			}
		}
		if !flag {
			t.Fail()
		}
	}
}

func Test_simIntArray_Len(t *testing.T) {
	array := simIntArray{}
	array.Src([]int{1, 1, 3, 4})
	if array.Len() != 4 {
		t.Fail()
	}
}

func Test_simIntArray_ToArray(t *testing.T) {
	array := simIntArray{}
	src := []int{1, 2, 3, 4}
	array.Src(src)

	if !reflect.DeepEqual(array.ToArray(), src) {
		t.Fail()
	}
}

func Test_simIntArray_Contains(t *testing.T) {
	array := simIntArray{}
	array.Src([]int{100, 101, 102})

	if !array.Contains(100) {
		t.Fail()
	}
	if array.Contains("100") {
		t.Fail()
	}
	if array.Contains(104) {
		t.Fail()
	}
}
