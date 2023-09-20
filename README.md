# Fault
Fault is a modeling language for building system dynamic models and checking them using a combination of first order logic and probability

## Project Status
Pre-alpha.

## Install
Fault can be built for source if you like, but the best way to install Fault is by [downloading the correct release for your machine](https://github.com/Fault-lang/Fault/releases).

Once installed the model checker of Fault needs access to a SMT solver, otherwise Fault will default to generating SMT of models only. Microsoft's Z3 is the recommended solver at this time and [can be downloaded here](https://github.com/Z3Prover/z3/releases)

Then in order for Fault to find your solver you need to set two configuration variables

```
export SOLVERCMD="z3"
export SOLVERARG="-in"
``` 

For other install options please [see the Fault documentation](https://www.fault.tech)

## Why "Fault"?
It is not possible to completely specify a system. All specifications must decide what parts of the system are in-scope and out-of-scope, and at what level of detail. Many formal specification approaches are designed to prove the system correct and it is very easy for an inexperienced practitioner to write a bad spec that gives a thumbs up to a flawed system.

Instead Fault is designed with the assumption that all systems will fail at some point, under some set of conditions. The name Fault was chosen to emphasize this point for users: Fault models that return no failure points are bad models. The user should keep trying until they've built a model that produces interesting and compelling failure scenarios.

## Origin Story
The development Fault is documented in the series "Marianne Writes a Programming Language":

- [audio](https://anchor.fm/mwapl)
- [transcripts](https://dev.to/bellmar/series/9711)

### Current Status (9/19/2023)
Pushing a bunch of small patches and bug fixes in the ramp up to presenting Fault at Strange Loop 2023 \o/

#### (8/31/2023)
Been doing a lot of work on how model outputs are displayed. Needed to fix how dead branches are filtered first. The default display is now formatted like an event log. Also added a static format that just outputs variable values.

#### (5/24/2023)
Introducing strings to Fault, plus fixing some bugs in the behavior of imports. Strings in Fault behave like booleans and allow easily readable rules to be defined for the model checker to solve. They are not treated as immutable but haven't yet figured out how I want to syntax to change their state to look yet.

#### (4/26/2023)
Adding support for indexes (ie accessing historical values in the model) there are some edge cases around branches that need to be worked through.

#### (4/13/2023)
BREAKING CHANGE. Adding support for altering the initial values of new instances of stocks and flows ended up requiring a lot of fundalmental changes, the big one being separating the initialization of model structs for the logic of a loop. Previously these both lived in the run block (`for 2 run{...}`) but now the run block has a optional `init{}` clause where initializations go, leaving the run block with just the steps the happen in the loop. The documentation and example specs have been updated to reflect this new construction.  

#### (3/29/2023)
Been doing a lot of work on improving how Fault is packaged and ultimately released. The Dockerfiles will remain available, but going to be stepping away from Docker as the preferred way to installing Fault in favor of a traditional build and release pipeline.

#### (2/15/2023)
Pretty substantial rewrite of assert and assumption rule generation. First ditched assumptions as a unique AST node and added an assume flag to AssertionStatements so they could be treated the same. Then modified the order of rules involving infixs and found several logic bugs from the old approach.

Along the way, made changes to LLVM IR generation that allow rounds to be tracked so that Fault knows when two states of different variables coexist in time. Will eventually use the same approach to get rid of LLVM IR metadata for tracking concurrency, which will eliminate the issues with some of LLVM optimization passes removing metadata.

#### (12/30/2022)
Taking another stab at the interface question and display of results. Settled on using mermaid to visualize the solutions received from the solver which is much more useful than what I was trying otherwise. This reintroduces the dotviz generation from way back with the prototype so I've also included generated mermaid viz for the state machine and the active stock-flow subsystems.

#### (11/30/2022)
Trying to finish up state charts exposed some problems with conditionals that needed to be addressed, plus some bugs and funny edge cases. But state charts are now done, including a reachability analysis verifying that the system is appropriately specified that I think will be opt-in for now.

#### (11/2/2022)
While working on adding state chart support I finally hit the limit on what the half-assed namespace implementation I started with could support. So I ended up spending the entire month of October writing a preprocesser that annotated the AST with the proper ids for each nameable node, then integrating it into a compiler. It took a long time but the code in the LLVM compilation stage is so much cleaner and neater now. Along the way found some previously unknown bugs and added some more tests to bone up mutation scores (will need more of this later)

#### (9/28/2022)
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

