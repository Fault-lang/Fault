# TODOs

## LookupType — add Int support
`generator/unroll/variables.go` — `LookupType` currently only returns "Bool" or "Real".
Should probably handle Int as a distinct SMT type rather than treating all LLVM ints as Bool.
