# TODOs

## LookupType — add Int support
`generator/unroll/variables.go` — `LookupType` currently only returns "Bool" or "Real".
Should probably handle Int as a distinct SMT type rather than treating all LLVM ints as Bool.

## Basic rule type — possibly dead code
`generator/rules/rules.go` — `Basic` is never instantiated anywhere in the codebase.
Investigate whether it can be removed.
