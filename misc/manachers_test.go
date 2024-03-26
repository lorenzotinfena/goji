package misc_test

import (
	"testing"

	"github.com/lorenzotinfena/goji/misc"
)

func TestManachersAlgorithms(t *testing.T) {
	misc.ManachersAlgorithm([]byte("abcbcbab"))
	res := misc.ManachersAlgorithm([]byte("abba"))
	t.Log(res[2*1+2])
}
