package kata5
// buble sort


func make_kata(input []int64) []int64 {
	lastElement := len(input) - 1
	for {
		for i := 0; i <= lastElement; i++ {
			if needToPush(input, i) {
				push_buble(input, i)
			}
		}
		lastElement -= 1
		if lastElement < 0 {
			break
		}
	}

	return input

}

func push_buble(slice []int64, position int) {
	var storage int64
	storage = slice[position + 1]
	slice[position + 1] = slice[position]
	slice[position] = storage
	return

}

func needToPush(input []int64, position int) bool {
	return position < len(input) -1 && input[position] > input[position +1]
}