package kata

import "testing"

type fixa struct {
    input    string
    index    int
    expected string
}

var _FIXTURES []fixa = []fixa{
    {"ebana", 1, "eana"},
    {"ebana", 2, "ebna"},
    {"ebta", 1, "eta"},
    {"ebta", 0, "bta"},
}

func Test_cut(t *testing.T) {
    for _, f := range _FIXTURES {
        result := f.input
        cut(&result, f.index)
        if result != f.expected {
            t.Error(result, " is not equal to ", f.expected)
        }
    }
}

func Benchmark_cut(b *testing.B) {
    for i := 0; i < b.N; i++ {
        str := "ebana"
        cut(&str, 3)
    }
}
