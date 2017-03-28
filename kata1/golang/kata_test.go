package kata

import "testing"

var FIXTURES [][]string = [][]string{
    {"eeee", "))))"},
    {"eeeae", ")))()"},
}

func TestKata1(t *testing.T) {
    for _, fixture := range FIXTURES {
        result := MakeKata(fixture[0])
        if result != fixture[1] {
            t.Error(result, " is not equal to ", fixture[0])
        }
    }
}
