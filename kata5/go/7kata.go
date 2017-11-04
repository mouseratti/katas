package kata5

type Sortable interface {
	Equal(s Sortable) bool
	Greater(s Sortable) bool
}

func Sort(input []Sortable) []Sortable {
	lastElement := len(input) - 1
	var storage Sortable
	for lastElement > 0 {
		for i := 0; i < lastElement; i++ {
			if input[i].Greater(input[i+1]) {
				storage = input[i+1]
				input[i+1] = input[i]
				input[i] = storage
			}

		}
		lastElement -= 1
	}
	return input
}

type MyInt int

func (self *MyInt) Equal(i Sortable) bool {
	return self == i
}

func (self *MyInt) Greater(i Sortable) bool {
	return bool(i - self)
}
