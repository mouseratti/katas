package kata5

func Make_kata(in []bool) []bool {
	last := len(in) - 1
	for last > 0 {
		for i := 0; i < last; i++ {
			if in[i] && !in[i+1] {
				in[i], in[i+1] = false, true
			}
		}
		last -= 1
	}

	return []bool{}
}
