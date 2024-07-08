# Fault
Fault is a modeling language for building system dynamic models and checking them using a combination of first order logic and probability

## Project Status
Pre-alpha.

## Install
Fault can be built from source if you like, but the best way to install Fault is by [downloading the correct release for your machine](https://github.com/Fault-lang/Fault/releases).

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
