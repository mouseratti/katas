package kata6

import "katas/fixtures"

func make_kata1(f *fixtures.Fixture) {
	last := len(f.Input) - 1
	var storage int64
	for i := 1; i <= last; i ++ {
		for j := 0; j < i; j++ {
			if f.Input[i] < f.Input[j] {
				storage = f.Input[i]

				for k := i; k > j; k-- {
					f.Input[k] = f.Input[k-1]
				}
				f.Input[j] = storage
			}
		}
	}
}
