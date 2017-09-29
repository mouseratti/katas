package kata6

type Sorted interface {
	Sort()
}

type KataFixture6 struct {
	input    []int
	expected []int
}

func (self *KataFixture6) Sort() {
}

func (self *KataFixture6) changePosition(position int, newPosition int) {
	switch newPosition - position {
	case 1, -1:
		s := self.input[position]
		self.input[position], self.input[newPosition] = self.input[newPosition], s
		return
	default:
		break
	}

	for self.input[position] != self.input[newPosition] {
		nextPosition := getNextPos(position, newPosition)
		self.changePosition(position, nextPosition)
		position = nextPosition
	}
	return

}

//func make_kata(s Sorted) Sorted {
//	s.Sort()
//	return s
//}

func getNextPos(pos int, newPos int) (result int) {
	result = pos
	if pos > newPos {
		result = pos - 1
	} else {
		result = pos + 1
	}
	return result
}
