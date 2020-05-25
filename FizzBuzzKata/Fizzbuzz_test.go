package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t*testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"By default returns the same number that is passed in", args{7}, "7"},
		{"Returns Fizz when the number is divisible by three", args{6}, "Fizz"},
		{"Returns Buzz when the number is divisible by five", args{5}, "Buzz"},
		{"Returns FizzBuzz when the number is divisible both by three and five", args{15}, "FizzBuzz"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t*testing.T) {
			assert.Equal(t, test.want, Calculate(test.args.number))
		})
	}
}