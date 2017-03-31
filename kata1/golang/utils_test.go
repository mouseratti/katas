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

// type fixture struct {
//     input    string
//     symbol   rune
//     count    int
//     expected string
// }

// var FIXTURES_REMOVE_FROM_STRING []fixture = []fixture{
//     {"eeeeee", 'e', 1, "eeeee"},
//     {"abeecc", 'e', 2, "abcc"},
//     {"qwertyyyc", 'y', 0, "qwert"},
//     {"qwertyyycyyyyyy", 'y', 0, "qwert"},
//     {"qwertyyyc", 'q', 0, "wertyyyc"},
// }

// func Test_RemoveFromString(t *testing.T) {
//     for _, fixtura := range FIXTURES_REMOVE_FROM_STRING {
//         result := RemoveFromString(fixtura.input, fixtura.symbol, fixtura.count)
//         if result != fixtura.expected {
//             t.Error(result, " is not equal to ", fixtura.expected)
//         }
//     }
// }
