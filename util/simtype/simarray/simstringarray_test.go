package simarray

import (
	"reflect"
	"testing"
)

func Test_simStringArray_Src(t *testing.T) {
	array := simStringArray{}
	if array.array != nil {
		t.Fail()
	}

	src := []string{"1", "2", "3"}
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

func Test_simStringArray_Len(t *testing.T) {
	array := simStringArray{}
	array.Src([]string{"1", "2", "3", "4"})
	if array.Len() != 4 {
		t.Fail()
	}
}

func Test_simStringArray_ToArray(t *testing.T) {
	array := simStringArray{}
	src := []string{"1", "2", "3", "4"}
	array.Src(src)

	if !reflect.DeepEqual(array.ToArray(), src) {
		t.Fail()
	}
}

func Test_simStringArray_Contains(t *testing.T) {
	array := simStringArray{}
	array.Src([]string{"123", "345"})

	if !array.Contains("123") {
		t.Fail()
	}

	if array.Contains(123) {
		t.Fail()
	}
	if array.Contains("789") {
		t.Fail()
	}
}
