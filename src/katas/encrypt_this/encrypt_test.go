package encrypt_this

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	{"p!zda,", "112azd!"},
}

func Test_make_kata(t *testing.T) {
	for _, f := range fixtures {
		t.Run(f.input, func(t *testing.T) {
			assert.Equal(t, f.expected, make_kata1(f.input), f.input)
		})
	}

}

func Test_encrypt_word(t *testing.T) {
	for _, f := range fixtures_encrypt_word {
		t.Run(f.input, func(t *testing.T) {
			assert.Equal(t, f.expected, encrypt_word(f.input), f.input)
		})
	}

}
