package fixtures

import (
	"math/rand"
	"time"
)

type Fixture struct {
	Input []int64
}

func (self *Fixture) Is_sorted() (result bool) {
	result = true
	last := len(self.Input) -1
	for i := 0; i < last; i++ {

		if self.Input[i] > self.Input[i+1] {
			result = false
			break
		}
	}
	return
}

func GetFixture(length int) Fixture {
	rand.Seed(time.Now().UnixNano())
	f := Fixture{make([]int64, length)}
	for i := 0; i < length; i++ {
		f.Input[i] = rand.Int63n(100)
	}
	return f
}
