package name

import "fmt"

var blockIndex uint64
var anonFuncIndex uint64

func Block() string {
	name := fmt.Sprintf("block-%d", blockIndex)
	blockIndex++
	return name
}

func AnonFunc() string {
	name := fmt.Sprintf("fn-%d", anonFuncIndex)
	anonFuncIndex++
	return name
}

func Var(prefix string) string {
	name := fmt.Sprintf("%s-%d", prefix, blockIndex)
	blockIndex++
	return name
}
