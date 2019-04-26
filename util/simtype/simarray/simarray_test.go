package simarray

import (
	"testing"
)

func Test_NewSimArray_success(t *testing.T) {
	src := []string{"1", "2"}
	array, err := NewSimArray(src)
	if err != nil {
		t.Fail()
	}

	for _,value := range array.array
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
