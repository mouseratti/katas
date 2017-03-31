package kata

import "testing"

var FIXTURES [][]string = [][]string{
    {"⌘eeee", "())))"},
    {"eeeae", ")))()"},
    {"The Goal of this", "))()()(())()))(("},
    {"Where Each Character", "())))))))))))))))())"},
    {"Exercise is to Convert", ")()))))))))))))))(()))"},
}

type fixture struct {
    input    string
    symbol   rune
    count    int
    expected string
}

var FIXTURES_REMOVE_FROM_STRING []fixture = []fixture{
    {"eeeeee", 'e', 1, "eeeee"},
    {"abeecc", 'e', 2, "abcc"},
    {"qwertyyyc", 'y', 0, "qwert"},
    {"qwertyyycyyyyyy", 'y', 0, "qwert"},
    {"qwertyyyc", 'q', 0, "wertyyyc"},
}

func Test_RemoveFromString(t *testing.T) {
    for _, fixtura := range FIXTURES_REMOVE_FROM_STRING {
        result := RemoveFromString(fixtura.input, fixtura.symbol, fixtura.count)
        if result != fixtura.expected {
            t.Error(result, " is not equal to ", fixtura.expected)
        }
    }
}

// func TestKata1(t *testing.T) {
//     for _, fixture := range FIXTURES {
//         result := MakeKata(fixture[0])
//         if result != fixture[1] {
//             t.Error(result, " is not equal to ", fixture[1])
//         }
//     }
// }

// func BenchmarkKata1(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         MakeKata("This is the best string I ever Watched")
//     }
// }
