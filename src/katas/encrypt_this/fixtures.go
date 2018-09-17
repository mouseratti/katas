package encrypt_this

var fixtures = []struct {
	input    string
	expected string
}{
	{"Hello", "72olle"},
	{"good", "103doo"},
	{"hello world", "104olle 119drlo"},
	{"Fuck my brain!!!", "70kcu 109y 98nair"},
}
var fixtures_encrypt_word = []struct {
	input    string
	expected string
}{
	{"Hello", "72olle"},
	{"good", "103doo"},
	{"brain!!!!!!!!", "98nair"},
	{"p!nta,", "112ant!"},
}
var fixtures_unused_symbols = []struct {
	input    string
	expected string
}{
	{"Hello!!!!!!!!!!!!", "Hello"},
	{"good,", "good"},
	{"bad.", "bad"},
	{"plokho...", "plokho"},
	{"pr!vet", "pr!vet"},
}

var fixtures_isInSet = []struct {
	element  interface{}
	set      []interface{}
	expected bool
}{
	{1, []interface{}{1, 2, 3}, true},
	{"a", []interface{}{"a", "b", "c"}, true},
	{5, []interface{}{"a", 2, int64(10)}, false},
	{'!', []interface{}{'!', ',', '.'}, true},
}
