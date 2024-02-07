package trie_test

import (
	"testing"

	"github.com/lorenzotinfena/goji/collections/trie"
	"github.com/stretchr/testify/assert"
)

func TestHAMT(t *testing.T) {
	hamt := trie.NewHAMT[int, int]()
	hamt.Set(1, 10)
	hamt.Set(10, 100)
	hamt.Set(100, 1000)
	hamt.Set(1000, 1)
	hamt.Set(10000, 1)
	v, present := hamt.Get(1)
	assert.True(t, present)
	assert.Equal(t, v, 10)
	v, present = hamt.Get(10)
	assert.True(t, present)
	assert.Equal(t, v, 100)
	v, present = hamt.Get(100)
	assert.True(t, present)
	assert.Equal(t, v, 1000)
	v, present = hamt.Get(1000)
	assert.True(t, present)
	assert.Equal(t, v, 1)
	v, present = hamt.Get(10000)
	assert.True(t, present)
	assert.Equal(t, v, 1)
}
