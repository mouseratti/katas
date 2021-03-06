package kata5

//func make_kata2(input ...int64) []int64 {
//	last := len(input) - 1
//	for {
//		for i := 0; i < last; i++ {
//			if need_push(i, input...) {
//				push_up(i, input...)
//			}
//		}
//		last -= 1
//		if last < 0 {
//			break
//		}
//	}
//
//	return input
//}
//
//func need_push(pos int, inp ...int64) bool {
//	return len(inp) - 1 > pos && inp[pos] > inp[pos + 1]
//}
//func push_up(pos int, inp ...int64) {
//	s := inp[pos + 1]
//	inp[pos + 1], inp[pos] = inp[pos], s
//}

func make_kata2(input []int64) []int64 {
	length := len(input)
	last := length - 1
	for {
		for i := 0; i < last; i++ {
			try_push(input, i, last)
		}
		last -= 1
		if last == 0 {
			break
		}
	}
	return input
}

func try_push(input []int64, pos int, last int) {
	if pos == last {
		return
	}
	if input[pos] > input[pos+1] {
		s := input[pos]
		input[pos] = input[pos+1]
		input[pos+1] = s

	}
	return
}
