package kata6

import "katas/fixtures"

func make_kata3(f *fixtures.Fixture) {
	for pos, val := range f.Input {
		if pos == 0 {
			continue
		}
		if val < f.Input[pos-1] {
			for i := 0; i < pos; i++ {
				if val < f.Input[i] {
					for j := pos; j > i; j-- {
						f.Input[j] = f.Input[j-1]
					}
					f.Input[i] = val
					break
				}
			}
		}
	}
}
