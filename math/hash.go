package math

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

// Non-cryptographic hash (FNV-1a)
func Hash(data any) uint {
	f := fnv.New64a()
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	_ = enc.Encode(data)
	f.Write(buf.Bytes())
	return uint(f.Sum64())
}
