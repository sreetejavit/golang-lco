package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	ListOfInputs := []addvalue{
		addvalue{1, 2},
		addvalue{2, 3},
		addvalue{3, 4},
	}

	for _, tt := range ListOfInputs {
		result := Add(tt.one, tt.two)
		c := tt.one + tt.two + 1

		assert.Equal(t, c, result)

	}
}
