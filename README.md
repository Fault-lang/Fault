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

## Getting Started
_Fault is currently pre-alpha and not ready to develop real specs, but if you like pain and misery here's how to run the compiler..._

Fault is written in Go and can be run by downloading this repo and running this command:

`make fault-z3`

This will build Fault with Z3 as a solver backend. This will require you to have Docker installed. From there Fault specs can be run like so:

`fault -f example.fspec`

That will return the SMTLib2 output of the compiler. Please note that the compiler only supports part of the Fault grammar currently and the formatting to the results needs some work.

You can output different stages of compilation by using the `-mode` flag. By default this is set to `-mode=check`, but can be changed to output either `ast`, `ir`, or `smt` which will stop compilation early and output either Fault's AST, LLVM IR, or SMTLib2 respectively.

You can also start the compiler from the LLVM -> SMTLib2 stage by changing to `-input` flag to `-input=ll`. By default the compiler expects the input file to be a spec that fits the Fault grammar.

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
| Spec imports | :white_check_mark: | | |
| Conditionals | :white_check_mark: | | |
| Uncertain data types | :white_check_mark: | :white_check_mark: | |
| Non-negative data types | | | |
| Assertions | :white_check_mark: | :white_check_mark: | |

### Development Strategy
The assumption Fault is making is that since both system dynamic models and first order logic models represent things as state machines it should be possible for a language to take the imperative structure of system dynamic DSLs, compile them to the declarative structure of logic DSLs and create a model checker better suited for the day-to-day software work of professionals.

There are A LOT of assumptions there, so the pre-alpha development of Fault prioritizes the quickest paths to verifying those assumptions over a comprehensive implementation of any one stage of the compiler. It makes no sense to spend weeks/months crafting a thoughtful and elegant type checker only to find out that SMT solvers cannot handle to level of complexity most of Fault's potential users would need to represent in order for Fault to be useful. SMT solvers tend to be very particular, with lots of quirky performance issues.

But then that's part of the fun too. Developing Fault is an opportunity to learn more about how SMT solvers (specifically Z3) work.

### Current Status (9/28/2022)
Adding state chart support to the parser and LLVM compiler, plus implementing logic around "this", cleaning up some dead code, plus some minor adjustments here and there.

#### (8/05/2022)
In order to prepare for state chart organization, added support for Booleans. Also tweaked the syntax to allow values to be overwritten within a flow.

#### (7/24/2022)
For the past couples of months I've been stuck on the Fault interface. When Z3 returns results how should that be displayed so that the user understands the failure case described by the model? This is tricky because Z3 will solve for all defined variables, even the ones in branches not relevant to the rest of the model.
Initially I was playing around with the idea of using Bubble Tea to do a more robust visual interface from the command line, but eventually a scrapped that as being too complicated. The real problem was filtering out inactive branches after the phi. I had a couple of pieced together approaches, but there was a lot of technical debt in figuring out SMT generation that made the code difficult to read and overwhelming. I finally decided to just completely rewrite the SMT package to make it a little bit easier to figure out and the bake in filtering of return values for the interface.

#### (3/18/2022)
Took another look at assert generation and made them tweaks. Found some bugs in unknown variables. Big deal is implemented temporal logic on assert generation. Now in addition to generating asserts for traditional temporal logic like "always", "eventually" and "eventually-always", Fault also has a set of specific temporal functions like "no more than" (nmt) and "no fewer than" (nft) which allow model checking a stable loop.

#### (2/2/2022)
In the process of adding support for unknown variables I realized I never fully connected the dots on uncertain values (whoops). So finished both of those although the logic around asserts over multiple instances is kind of wonky and brittle. May need a rethink.

#### (1/26/2022)
Removed the fuzzer for the time being. It wasn't really doing what I needed it to do and the newest version Go starts to roll out fuzzing functionality by default 🎉 As added some error handling around the lexer/parser.

#### Status (1/7/2022)
This started with an honest attempt to set up CI/CD on Fault's repo so that other people can start contributing, but debugging the CI/CD pipeline made me realize I have cross platform capability issues :facepalm:. Tried a bunch of different things and conveniently forgot to swash my commits before merging to main. All my dirty laundry is there to study!

Anyway! Long story short: Fault now has an installer and runs on Docker. It also has support for alternative SMT solvers rather than a dependency on Z3 ... at least in theory!

#### Status (12/05/2021)
Development of Fault kind of goes like this: I write a spec and then implement whatever features are currently missing in order to get it to run. This time I added something unexpected: support for nestled stocks (stocks of stocks). Adds a little complexity but makes the specs look so much cleaner.

#### Status (12/01/2021)
Using Go channels to compile parallel runs seemed like a clevel solution, but truthfully the problems with correctly handling SSA weren't offset by any benefits in performance. Might revist in the future as models grow more complex. For right now plain ordinary sequencial processing of all premutations works well.

Other major thing is parsing the SMT returned by the solver and formatting those results into a human friendly form. Laid some of the ground work on Uncertain types.

#### Status (10/11/2021)
Just finished the happy path on conditionals, want to shift to spec imports next. Still have to test LLVM IR -> SMTLib2 after LLVM optimization passes. 

