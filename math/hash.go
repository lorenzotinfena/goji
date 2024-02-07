package math

import (
	"fmt"
	"hash/fnv"
	"reflect"
)

// Non-cryptographic hash (FNV-1a)
func Hash(data any) uint {
	f := fnv.New64a()
	if reflect.ValueOf(data).Kind() == reflect.Ptr {
		f.Write([]byte(fmt.Sprintf("%p", interface{}(data))))
	} else {
		f.Write([]byte(fmt.Sprintf("%v", data)))
	}
	return uint(f.Sum64())
}
