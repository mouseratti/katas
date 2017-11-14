package kata6

import "katas/fixtures"

//TODO fix this kata
func make_kata2(f *fixtures.Fixture) {
	length := len(f.Input)
	last := length - 1
	current := last
	var storage int64
	var position int

	for current >= 0 {
		for i := current - 1; i < length; i++ {
			if f.Input[current] < f.Input[i] {
				position = i
			} else {
				position = last
			}
			storage = f.Input[current]
			for j := current; j < position; j++ {
				f.Input[j] = f.Input[j+1]
			}
			f.Input[position] = storage
		}
		current-=1
	}
}
