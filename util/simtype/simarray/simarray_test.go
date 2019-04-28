package simarray

import (
	"reflect"
	"testing"
)

func Test_SimArray_NewSimArray(t *testing.T) {
	src := []string{"1", "2"}
	array, err := NewSimArray(src)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(array.array.ToArray(), src) {
		t.Fail()
	}
}

func Test_SimArray_Contains(t *testing.T) {
	src := []string{"1111", "2222"}
	array, err := NewSimArray(src)
	if err != nil {
		t.Fail()
	}
	if !array.Contains("1111") {
		t.Fail()
	}
	if array.Contains(1111) {
		t.Fail()
	}
	if array.Contains("3333") {
		t.Fail()
	}
}

func Test_SimStringArray_ContainsAny(t *testing.T) {
	h1 := "go say hello world"
	h2 := "csharp say hello world"
	src := []string{h1, h2, "go say ok", "csharp say ok"}
	array, err := NewSimArray(src)
	if err != nil {
		t.Fail()
	}
	result := array.ContainsAny("hello world")
	if len(result) != 2 {
		t.Fail()
	}

	for _, data := range result {
		if h1 != data && h2 != data {
			t.Fail()
		}
	}
}

// func Test_ArrayContainsByRefelct(t *testing.T) {
// 	iscnt := ArrayContainsByRefelct([]string{"111", "1223", "122"}, "111")
// 	if !iscnt {
// 		t.Fail()
// 	}
// }

// func Test_SimArrayContains(t *testing.T) {
// 	iscnt := SimArrayContains([]int{111, 1223, 122}, 111)
// 	if !iscnt {
// 		t.Fail()
// 	}
// }

// func mockArrayContains(src []string, cnt string) bool {
// 	for _, obj := range src {
// 		if obj == cnt {
// 			return true
// 		}
// 	}
// 	return false
// }

// func Benchmark_ArrayContains_mock(b *testing.B) { // 普通的查询
// 	src := []string{"1111", "2131", "122", "3131", "31", "121313132", "111121", "331", "12213", "13131", "121", "1313132"}
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		mockArrayContains(src, "121")
// 	}
// 	b.StopTimer()
// }

// func Benchmark_ArrayContains_simarray(b *testing.B) { // simarray的查询
// 	src := []string{"1111", "2131", "122", "3131", "31", "121313132", "111121", "331", "12213", "13131", "121", "1313132"}
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		ArrayContainsByRefelct(src, "121")
// 	}
// 	b.StopTimer()
// }

// func Benchmark_SimArrayContains_simarray(b *testing.B) {
// 	src := []string{"1111", "2131", "122", "3131", "31", "121313132", "111121", "331", "12213", "13131", "121", "1313132"}
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		SimArrayContains(src, "121")
// 	}
// 	b.StopTimer()
// }
