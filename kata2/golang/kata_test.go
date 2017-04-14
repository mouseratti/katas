package kata2

import "testing"
import "fmt"

func TestKata(t *testing.T) {
    for _, fixture := range FIXTURES {
        if !fixture.Assert() {
            t.Error(fmt.Sprintf("%v is not equal to %v", fixture.Result(), fixture.expected))
        }
    }

}

func Test_rot13(t *testing.T) {
    for pos := range input {
        result := rot13(input[pos])
        if result != output[pos] {
            t.Error(fmt.Sprintf("%c(%v) is not equal to %c(%v)", result, result, output[pos], output[pos]))
        }
    }
}

func BenchmarkKata(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MakeKata("This is the best string I ever Watched")
    }
}
