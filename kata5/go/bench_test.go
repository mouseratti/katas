package kata5

import "testing"

func BenchmarkKata(b *testing.B) {
	for i := 0; i < b.N; i++ {
		make_kata([]int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	}
}

func BenchmarkKata2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		make_kata2([]int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	}
}

func BenchmarkKata3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		make_kata3([]int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	}
}

func BenchmarkKata4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		make_kata4([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	}
}


func BenchmarkKata5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		make_kata5([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	}
}

func BenchmarkKata5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		make_kata5([]int{10,9,8,7,6,5,4,3,2,1})
	}
}
