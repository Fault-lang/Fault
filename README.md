# Fault
Fault is a modeling language for building system dynamic models and checking them using a combination of first order logic and probability

## Project Status
Pre-alpha.

## Why "Fault"?
It is not possible to completely specify a system. All specifications must decide what parts of the system are in-scope and out-of-scope, and at what level of detail. Many formal specification approaches are designed to prove the system correct and it is very easy for an inexperienced practitioner to write a bad spec that gives a thumbs up to a flawed system.

Instead Fault is designed with the assumption that all systems will fail at some point, under some set of conditions. The name Fault was chosen to emphasize this point for users: Fault models that return no failure points are bad models. The user should keep trying until they've built a model that produces interesting and compelling failure scenarios.

## Origin Story
The development Fault is documented in the series "Marianne Writes a Programming Language":

- [audio](https://anchor.fm/mwapl)
- [transcripts](https://dev.to/bellmar/series/9711)

## Todos
_incomplete list. Items to be added as I think of them_

| Task | Happy Path | Edge Cases | Fuzz |
| :--: | :--: | :--: | :--: |
| BNF Grammar | :white_check_mark: | :white_check_mark: | :white_check_mark:|
| Lexer/Parser | :white_check_mark: | :white_check_mark: | |
| Type checking | :white_check_mark: | | |
| LLVM IR generation | :white_check_mark: | | |
| LLVM optimization passes | | | |
| SMTLib2 generation | :white_check_mark: | | |
| Spec imports | | | |
| Conditionals | :white_check_mark: | | |
| Uncertain data types | | | |
| Non-negative data types | | | |
| Assertions | | | |

### Development Strategy
The assumption Fault is making is that since both system dynamic models and first order logic models represent things as state machines it should be possible for a language to take the imperative structure of system dynamic DSLs, compile them to the declarative structure of logic DSLs and create a model checker better suited for the day-to-day software work of professionals.

There are A LOT of assumptions there, so the pre-alpha development of Fault prioritizes the quickest paths to verifying those assumptions over a comprehensive implementation of any one stage of the compiler. It makes no sense to spend weeks/months crafting a thoughtful and elegant type checker only to find out that SMT solvers cannot handle to level of complexity most of Fault's potential users would need to represent in order for Fault to be useful. SMT solvers tend to be very particular, with lots of quirky performance issues.

But then that's part of the fun too. Developing Fault is an opportunity to learn more about how SMT solvers (specifically Z3) work.

### Current Status (10/11/2021)
Just finished the happy path on conditionals, want to shift to spec imports next. Still have to test LLVM IR -> SMTLib2 after LLVM optimization passes. 

