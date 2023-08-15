package sum

import (
	"github.com/stretchr/testify/assert"
	// S "test/sum"
	"testing"
)

func TestSum(t *testing.T) {
	a := 3
	b := 5
	expected := 8
	result := Sum(a, b)
	assert.Equal(t, expected, result, "Unexpected result")

}
