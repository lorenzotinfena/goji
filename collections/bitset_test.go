package collections_test

import (
	"testing"

	"github.com/lorenzotinfena/goji/collections"
	"github.com/stretchr/testify/assert"
)

func TestBitset(t *testing.T) {
	b := collections.NewBiset(1234)
	// test 1
	for i := 0; i < b.Len(); i++ {
		b.Set(i, i%2 == 1)
	}
	b.ShiftLeft(1)

	for i := 0; i < b.Len(); i++ {
		assert.Equal(t, b.Get(i), i%2 == 0)
	}
	b.ShiftRight(1)
	for i := 0; i < b.Len(); i++ {
		assert.Equal(t, b.Get(i), i%2 == 1)
	}

	// test 2
	for i := 0; i < b.Len(); i++ {
		b.Set(i, i%3 == 0)
	}
	b.ShiftRight(13)
	b.ShiftLeft(13)
	assert.True(t, b.Get(186))
	assert.False(t, b.Get(185))
	assert.False(t, b.Get(187))
}
