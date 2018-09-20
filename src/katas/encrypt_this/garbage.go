package encrypt_this

type garbage []interface{}

var _GARBAGE_SYMBOLS = garbage{
	'!',
	'.',
	',',
}

func (g *garbage) contains(i interface{}) bool {
	for _, val := range *g {
		if val == i {
			return true
		}
	}
	return false
}
