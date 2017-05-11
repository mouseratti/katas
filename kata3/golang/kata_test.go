package kata

import "testing"

func TestKata(t *testing.T) {
    for _, fix := range fixtures {
        subnet, broadcast := MakeKata(fix.input)
        if subnet != fix.subnet || broadcast != fix.broadcast {
            t.Errorf("subnet %v, broadcast %v != fix %v", subnet, broadcast, fix)
        }
    }
}

func BenchmarkKata(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MakeKata("109.130.40.1/28")
    }
}
