package kata5

import "katas/fixtures"

func make_kata5(f *fixtures.Fixture) {
	marker := len(f.Input) - 1

	for marker > 0 {
		for i := 0; i < marker; i++ {
			if f.Input[i] > f.Input[i+1] {
				f.Input[i], f.Input[i+1] = f.Input[i+1], f.Input[i]
			}
		}
		marker -= 1
	}
}

