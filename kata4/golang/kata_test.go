package kata

import "testing"

func Test_fizzbuzz(t *testing.T) {
	fixtures := map[int]string{
		1:  "1",
		2:  "2",
		6:  "fizz",
		15: "fizzbuzz",
		25: "buzz",
		30: "fizzbuzz",
		33: "fizz",
	}

	for key, val := range fixtures {
		result := fizzbuzz(key)
		if val != result {
			t.Errorf("fizzbuzz(%v) is %v != expected value %v", key, result, val)
		}
	}
}

func TestKata(t *testing.T) {
	MakeKata()
}

// func BenchmarkKata(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         MakeKata("10.30.40.1/28")
//     }
// }
