package buble_sort


func Kata1(input []int, doReverse bool) {
	length := len(input)
	maxElement := length - 1
	currentEdge := maxElement

	for currentEdge > 0 {
		for i := 0; i < currentEdge; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
			}
		}
		currentEdge -= 1
	}

	if doReverse {
		for k := 0; k < length/2; k++ {
			input[k], input[maxElement-k] = input[maxElement-k], input[k]
		}
	}
}

func Kata2(input []int, doReverse bool) {
		maxElement := len(input) -1
		for edge := maxElement; edge > 0; edge-- {
			for i := range input[:edge] {
				if input[i]>input[i+1] {
					input[i], input[i+1] = input[i+1], input[i]
				}
			}
		}

		if doReverse {
			for k:= 0; k < len(input) /2; k++ {
				input[k], input[maxElement-k] = input[maxElement-k], input[k]
			}
		}
	}