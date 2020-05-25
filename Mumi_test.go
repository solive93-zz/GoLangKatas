package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t*testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Replaces a vowel to mumy", args{"a"}, "mumy"},
		{"Replaces all same vowels in the same word", args{"caca"}, "cmumycmumy"},
		{"Replaces all vowels in the same word", args{"hola"}, "hmumylmumy"},
		{"Returns the input word if there's less than thirty per cent of vowels", args{"chus"}, "chus"},
		{"Replace for mumy just once if there's 2 vowels together", args{"laura"}, "lmumyrmumy"},
		{"Double vowel in the beginning of the word", args{"aerial"}, "mumyrmumyl"},
		{"Triple vowel still returns only one mumy", args{"miau"}, "mmumy"},

	}
	for _, test := range tests {
		t.Run(test.name, func(t*testing.T) {
			assert.Equal(t, test.want, Convert(test.args.input))
		})
	}
}