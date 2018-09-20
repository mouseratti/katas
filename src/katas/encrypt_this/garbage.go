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

type _GarbageMap map[interface{}]struct{}

func (g _GarbageMap) contains(i interface{}) bool {
	if _, ok := g[i]; ok {
		return true
	}
	return false
}

func makeGarbageMap(input []interface{}) _GarbageMap {
	var s struct{}
	gm := _GarbageMap{}
	for _, v := range input {
		gm[v] = s
	}
	return gm
}
