package simarray

import "testing"

func Test_StringArray_Len(t *testing.T) {
	array := SimStringArray{"1", "2", "3", "4"}
	if array.Len() != 4 {
		t.Fail()
	}
}

func Test_StringArray_Contains_不同类型(t *testing.T) {
	array := SimStringArray{"1", "2", "3", "4"}
	if array.Contains(1) {
		t.Fail()
	}
}

func Test_StringArray_Contains_不包含(t *testing.T) {
	array := SimStringArray{"1", "2", "3", "4"}
	if array.Contains("5") {
		t.Fail()
	}
}

func Test_stringArray_Contains_包含(t *testing.T) {
	array := SimStringArray{"1", "2", "3", "4"}
	if !array.Contains("3") {
		t.Fail()
	}
}
