package calc_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBit(t *testing.T) {

	assert.Equal(t, 1<<1, 2)
	assert.Equal(t, 1<<2, 4)
	assert.Equal(t, 1<<3, 8)
	assert.Equal(t, 1<<4, 16)

	assert.Equal(t, 2>>1, 1)
	assert.Equal(t, 4>>2, 1)
	assert.Equal(t, 8>>3, 1)
}
