package kata6

import "katas/fixtures"

func make_kata2(f *fixtures.Fixture) {
	length := len(f.Input)
	last := length - 1
	current := last - 1
	var storage int64
	var position int

	for current >= 0 {
		for i := last; i > current; i-- {
			if f.Input[current] > f.Input[i] {
				position = i
				storage = f.Input[current]
				for j := current; j < position; j++ {
					f.Input[j] = f.Input[j+1]
				}
				f.Input[position] = storage
			}
		}
		current-=1
	}
}
