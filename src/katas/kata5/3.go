package kata5

import "katas/fixtures"

func make_kata3(f *fixtures.Fixture) {
	last := len(f.Input) - 1
	for last > 0 {
		for i := 0; i < last; i++ {
			if f.Input[i] > f.Input[i+1] {
				f.Input[i], f.Input[i+1] = f.Input[i+1], f.Input[i]
			}
		}
		last -= 1
	}

}

