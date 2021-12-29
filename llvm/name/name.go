package name

import (
	"crypto/md5"
	"fmt"
)

var blockIndex uint64
var parallelIndex uint64
var anonFuncIndex uint64
var assertIndex uint64

func Block() string {
	name := fmt.Sprintf("block-%d", blockIndex)
	blockIndex++
	return name
}

func ParallelGroup(group string) string {
	data := []byte(fmt.Sprint(group, parallelIndex))
	parallelIndex++
	return fmt.Sprintf("%x", md5.Sum(data))
}

func AnonFunc() string {
	name := fmt.Sprintf("fn-%d", anonFuncIndex)
	anonFuncIndex++
	return name
}

func Assert() string {
	name := fmt.Sprintf("__assert-%d", assertIndex)
	assertIndex++
	return name
}

func Var(prefix string) string {
	name := fmt.Sprintf("%s-%d", prefix, blockIndex)
	blockIndex++
	return name
}
