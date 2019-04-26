package simarray

import "testing"

func Test_IntArray_Len(t *testing.T) {
	array := SimIntArray{11, 12}
	if array.Len() != 2 {
		t.Fail()
	}
}

func Test_IntArray_Contains_不同类型1(t *testing.T) {
	array := SimIntArray{1, 2, 3, 4, 5, 6}

	if array.Contains("11") {
		t.Fail()
	}
}

func Test_IntArray_Contains_不包含(t *testing.T) {
	array := SimIntArray{1, 2, 3, 4, 5, 6}

	if array.Contains(7) {
		t.Fail()
	}
}

func Test_IntArray_Contains_包含(t *testing.T) {
	array := SimIntArray{1, 2, 3, 4, 5, 6}

	if !array.Contains(3) {
		t.Fail()
	}
}
