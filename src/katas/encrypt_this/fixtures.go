package encrypt_this

import "strings"

var longString string = "For legacy purposes, we will continue to provide pre-generated changelogs and patches (both to the previous mainline and incremental patches to previous stable). However,"

var fixtures = []struct {
	input    string
	expected string
}{
	{"Hello", "72olle"},
	{"good", "103doo"},
	{"hello world", "104olle 119drlo"},
	{"Fuck my brain!!!", "70kcu 109y 98nair"},
}
var fixtures_encrypt_tokens = []struct {
	input    []string
	expected []string
}{
	{[]string{"hello"}, []string{"104olle"}},
	{[]string{"hello", "world"}, []string{"104olle", "119drlo"}},
	{[]string{"Fuck", "my", "brain!!!"}, strings.Split("70kcu 109y 98nair", " ")},
	{
		strings.Split(longString, " "),
		[]string{"70ro", "108ygace", "112srposeu", "119e", "119lli", "99entinuo", "116o", "112eovidr", "112de-generater", "99sangelogh", "97dn", "112stchea", "40hotb", "116o", "116eh", "112seviour", "109einlina", "97dn", "105lcrementan", "112stchea", "116o", "112seviour", "115)ablet", "72rweveo"},
	},
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

var largeSet garbage = garbage{1, "2", "3.0", "chetyure", '5', "!!!!!!", struct{}{}, 50000}
