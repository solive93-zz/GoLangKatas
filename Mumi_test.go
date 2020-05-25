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
		{"Convierte una vocal en mumy", args{"a"}, "mumy"},
		{"Convierte todas las vocales iguales de una palabra", args{"caca"}, "cmumycmumy"},
		{"Convierte todas loas vocales de una palabra", args{"hola"}, "hmumylmumy"},
		{"Devuelve la la misma palabra si no tiene un 30% de vocales", args{"chus"}, "chus"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t*testing.T) {
			assert.Equal(t, test.want, Convert(test.args.input))
		})
	}
}