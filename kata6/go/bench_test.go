package kata6

import "testing"

func BenchmarkKata(b *testing.B) {
	for i := 0; i < b.N; i++ {
		make_kata([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	}
}
