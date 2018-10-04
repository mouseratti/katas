package kata5

import "katas/fixtures"

func make_kata4(f *fixtures.Fixture) {
	marker := len(f.Input) - 1

	for marker > 0 {
		for pos, val := range f.Input[:marker] {
			if val > f.Input[pos+1] {
				f.Input[pos], f.Input[pos+1] = f.Input[pos+1], f.Input[pos]
			}
		}
		marker -= 1
	}

}
