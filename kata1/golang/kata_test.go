package kata

import "testing"

var FIXTURES [][]string = [][]string{
    {"⌘eeee", "())))"},
    {"eeeae", ")))()"},
    {"The Goal of this", "))()()(())()))(("},
    {"Where Each Character", "())))))))))))))))())"},
    {"Exercise is to Convert", ")()))))))))))))))(()))"},
    {"(bra) p(ora)", ")()))(()()))"},
}

func TestKata1(t *testing.T) {
    for _, fixture := range FIXTURES {
        result := MakeKata(fixture[0])
        if result != fixture[1] {
            t.Errorf("%T %v len %v != %T %v len %v",
                result, result, len(result),
                fixture[1], fixture[1], len(fixture[1]),
            )
        }
    }
}

func BenchmarkKata1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MakeKata("This is the best string I ever Watched")
    }
}
