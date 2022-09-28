# Fault Contributors Guide 
## About Fault
### What does Fault do?
#### The Fancy Answer
Fault is a domain specific language and model checker for [system dynamic models](https://systemdynamics.org/what-is-system-dynamics/). Fault allows you to define a feedback loop in a procedural style, that then cross-compiles to logical rules that are intelligible to an SMT solver ([Z3](https://github.com/Z3Prover/z3) by default)

#### The Pragmatic Answer
Fault is a tool to help engineers reason about systems by making it easier to model their behavior. Fault blurs the line between a verification model and a simulation. It uses SMT to find violations of invariants the same way a verification tool might, but it also includes constructs to allow the model to be manipulated based on probably behaviors.

### What do I use Fault for?
Fault is designed to model failure states in complex systems. The idea is you use Fault to do the initial hypothesis testing around interactions you think might trigger system failure. Although it uses many of the same technologies as other model checkers, the point of Fault is not to prove or verify a system is correct, but to explore the limits of it with logic programming.

### Technologies
Fault is written in Go. It uses [Antlr](https://www.antlr.org/) to generate a parser/lexer from a BNF grammar, and generates LLVM intermediate representation before ultimately compiling to [SMTLib2](https://smtlib.cs.uiowa.edu/)

### Am I Good Enough to Contribute?
Almost certainly yes! Fault was started because Marianne Bellotti wanted to learn the basics about compilers ([See: Marianne Writes a Programming Language](https://dev.to/mwapl)) and there are A LOT of mistakes and just general head scratching constructions still in the code base. Don't be intimidated by the topic area, you're definitely good enough to work on Fault if you are interested.

### How Fault Models Work
Traditional formal specifications are designed to verify and prove a given system is correct. They are difficult for most engineers to use because it is too easy to write a bad spec that verifies as correct. All you have to do is oversimpilify the model.

Fault specs, on the other hand, are written to help engineers reason about failure modes in complex systems. Fault assumes all systems will fail eventually and uses logic programming to define how interactions between various subsystems can create scenarios that trigger a failure mode.

The Fault compiler uses two different styles of programming to model check. The first and primary pass involves **logic programming**. The second and optional pass involves **probabilistic programming**.

#### Logic Programming
Logic programming is a style of declarative programming in which we write a series of rules and on execution the compiler produces a solution that satisfies all of the rules.

Fault specs are system diagrams that document how interactions between different components of a system change the distribution of resources (called stocks) throughout the system. The resources in question can be memory, I/O, CPU, requests, users, etc. The key characteristic of a stock is that when the system is thought of as a state machine, there is a function (or flow) that cause the stock to increase or decrease.

The Fault compiler takes this specification and translates it into a set of logic rules that will produce a solution the represents one valid state of the model given those rules. A user can write Fault models just by defining stocks and flows, but the real power of Fault's logic programming comes from including assertions in the model. Fault assertions are negated while compiling to SMT, producing a logic program that will search for a model state that will disprove the spec's own assertions.

States can also be removed from consideration by using `assume` statements in the model itself. Assume statements are structured the same as assertions except the compiler does not negate them. This means instead of searching for a state that invalidates the assumption, Fault  disregards any solution that violates those statements. In effect, it just assumes they are not relevant and ignores them.

#### Probability in Fault
Traditionally, logic based model checking is used to verify systems by proving that certain states are mathematically impossible. This is extremely useful when answering questions about algorithmic correctness, and less useful when answering questions about complex system behavior at scale. Complex systems tend to take the form of feedback loops and pure logic models produce many scenarios that although possible are also so unlikely they convey no useful information about the quality of the system.

For that reason, Fault includes some basic probabilistic programming with `uncertain` data types. A model can specify that a particular value is unknown but that the mean value is 10 and its standard deviation is 2.3 (expressed in Fault as `uncertain(10, 2.3)`). Fault will solve for this value as normal, then place the solution on a probability distribution and determine it's likelihood.

Probabilities in Fault currently assume that variables are independent, which is not the case if a model uses multiple uncertain values within the same feedback loop. At some point as Fault matures, a good way of either detecting or handling conditional probabilities needs to be developed.

By default Fault does not filter out unlikely values, but the user can set a tolerance with a configuration flag (_Note: not yet implemented_). As SMT solvers generate one solution at a time, setting a tolerance will cause Fault to append new rules to eliminate unlikely values and rerun the SMT to generate new potential solutions. This can increase the runtime of the model dramatically.

#### Conditionals, Branches and Phi Values
Fault specs can include conditionals or other types of branching behaviors. Logic programming handles branches with what are called Phi values. A Phi caps the end of a conditional with the end state of the branch taken. For example:

```
(assert (= x_1 (+ x_0 10.0)))
(assert (= x_2 (- x_1 20.0)))
(assert (= x_3 (- x_0 20.0)))
(assert (= x_4 (+ x_3 10.0)))
(assert (or (= x_5 x_2) (= x_5 x_4)))
```

The above SMTLib code represents two possible branches. In the first branch `x` is increased by 10 and then decreased by 20. In the second branch the order is reversed. The variable `x_5` is a Phi. States of `x` for both branches are deduced by the solver, then the next rule tells the solver to choose between the end state of branch 1 (`x_2`) and the end state of branch 2 (`x_4`) and assign that value to `x_5`.

There are two places in Fault where Phis come up: the first is if/else constructs, the second is when steps of a model are run in parallel.

Parallel steps are used to simulate race conditions. The example above is actually code from a Fault spec involving a parallel step. There is no actual concurrency here, instead the compiler generates branches for all possible orders of steps. If the spec above included an assertion that the value of `x` can never be less than some value, decreasing the value of x by 20 before increasing it might trigger a failure.

## The Structure of Fault Model
### FSystem files
All formal models (that is models based on the logic) are state machines, Fault brings this way of thinking to the forefront with the fsystem file which is the global view of the system being models. Fsystem files list all the components of the system and the states that component can be in. The body of each state is a definition of under what conditions the component transitions out of that state and into another state (or whether the state of that component causes another component to change state)

When the events that trigger a state change are too complex to express in simple conditionals, fsystem files can import fspec files. 

### FSpec files
Fspec files represent a change in resources that might represent a transition from one state to another. They use stocks to represent pools of resources and flows to represent change functions decreasing or increasing those resources as well as control logic. FSpec are meant to represent feedback loops.

## Compiler Walkthru

### Overview
The Fault compiler parses spec files written in Fault, constructs an AST for them, generates LLVM IR for the steps in each loop of the run block, then instead of opscode it generates an SMTLib2 representation and runs the model through an SMT solver (by default Z3), it then parses the results returned by the solver, and if the spec makes use of any uncertain data types it calculates the probability of the failure state actually happening and displays those results to the user.

Fault -> Antlr Parse Tree -> Fault AST -> Type Annotated Fault AST -> LLVM IR -> SMTLib2 Syntax -> SMTLib2 Solution -> Model Results

## Folders and Files

### AST
Every node in Fault's AST has three standard methods: `Token()` which returns the token associated with the node, `String()` which prints the content of the node and subnodes, and `Position()` which returns the line and column position in the original spec where the node comes from as a slice of `[start_line, start_column, stop_line, stop_column]`. 

### Parser/Lexer (Fault -> Antlr Parse Tree)
Fault's parser and lexer are generated by Antlr 4, using two grammar files `FaultLexer.g4` and `FaultParser.g4` respectively. If you add a new reserved word to Fault, you must include it in `FaultLexer.g4` before using it in `FaultParser.g4`. Otherwise the grammar will not compile correctly.

The grammars themselves are only useful when compiled by antlr. Included with the grammar files is a Makefile that makes that easier. `make golang` will generated the parser, lexer and bases for the listener used by Fault. 

`make gui` allows Fault code to be written to standard in and displays a visualization of the parse tree generated by Antlr. This is useful when debugging new features and requires you run the `make java` command first. 

The grammar of Fault follows roughly the syntax as Golang itself, with fewer data structures. Fault handles integers, floats, and booleans. The parser supports strings, although strings don't really do anything in Fault just yet. There are no lists, trees, dictionaries or sets but it does have stocks and flows which behave like objects.

### Listener (Antlr Parse Tree -> Fault AST)
Part of the parsing phase, the listener walks the Antlr specific parse tree and converts it to Fault's AST. The listener must traverse every node and uses a stack to pass information between children nodes and parents nodes. The way this works is that the listener moves down the tree until it reaches a terminal node and then returns, converting Antlr parser tree to Fault AST as it moves back up. Fault AST nodes are put on the top of the stack as the listener exits a Antlr node, then popped off the stack and wrapped in new parent node which itself is put on the stack.

For example:

```
func (l *FaultListener) ExitSpec(c *parser.SpecContext) {
	var spec = &ast.Spec{}
	for _, v := range l.stack {
		spec.Statements = append(spec.Statements, v.(ast.Statement))
	}
	l.AST = spec
}

func (l *FaultListener) EnterSpecClause(c *parser.SpecClauseContext) {
	l.currSpec = c.IDENT().GetText()
}
```
The `EnterSpecClause` executes as the listener enters the spec for the first time on the way down the tree, whereas the `ExitSpec` function executes as the listener comes back up.

The listener will check to make sure all the files being imported actually exist, then parse them first and add their AST trees to the Spec's tree.

Fault has a specific syntax around arrows: `foo.value -> x` decreases a stock called foo by x and `foo.value <- x` increases it by x. While parsing, the compiler converts these to `foo.value = foo.value - x` and `foo.value = foo.value + x` respectively

Fault grammar adds an underscore to some node names in order to avoid conflicts with Go's reserved words. A string node in the Antlr tree is a `string_` node instead. A float node `float_` instead.

### Type Checker (Fault AST -> Type Annotated Fault AST)
Like most type checkers, Fault's type checker annotates the existing AST with additional information. It uses a map to store all the stocks/flows defined in the spec and the dependencies imported. This map is reused by the LLVM stage. If the spec includes nestled stocks (stocks of stocks), the `Complex` flag in the relevant instance node is set to true.

For now the type rules are pretty basic:
- Ints can be converted to Floats, Naturals, and Booleans.
- Floats can be converted to Naturals or Booleans.
- Naturals can be converted to Booleans.
- Asserts and assumptions must be statements that ultimately produce a boolean without type conversion. Arithmetic operators can be used but `assert x + 5` is not a valid assertion
- Uncertain values are ignored, they can't be converted to anything else and are treated as Unknown types until execution stage 

The type checker calculates the scope of any numeric values. This is old code that traces back to the original implementation design which used fixed point instead of reals. It will be removed in future releases.

### LLVM (Type Annotated Fault AST -> LLVM IR)
Before generating SMTLib2, Fault converts the AST to the intermediate representation favored by LLVM. Because Fault models contain chunks of procedural code that roughly follow the structure of C inspired languages, they can be optimized via LLVM before converting to SMTLib. For elements that do not have a proper LLVM IR representation (such as steps that run in parallel in the model), metadata is attached to the relevant nodes. Therefore, optimization passes that remove metadata may cause Fault specs to run incorrectly. All numeric values are converted to Doubles at this stage.

The compiler converts all variable names to `[spec name]_[scope name]_[parameter name]`

Steps in the run block are assigned a unique group name (which is a hash) so that steps that are meant to be run in parallel can be identified after LLVM optimization process.

Identifiers in assert or assume statements are converted to a special `AssertVar` node because when generating SMT we need to generate cartesian products for these variables in order to cover constraints on the entire state space. (The conversion to `AssertVar` should probably be moved to the listener stage.)

### SMT (LLVM IR -> SMTLib2 Syntax)
At this stage functions are completely unrolled, variables are converted to single static assignment and SMTLib2 is generated. SMTLib2 syntax is very similar to Lisp and uses Polish notation. 

The SMTLib2 Fault produces is pretty primative and avoids using many of SMTLib2 advanced features in order to reduce complexity in the compiler. The original plan had been to represent all numeric types as Ints and used fixed point math. SMTLib2 does support floating points, but it is the relatively new area of research and [full of fun gotchas](https://stackoverflow.com/questions/54502823/are-floating-point-smt-logics-slower-than-real-ones/54505692). Another fun fact discovered during initial performance tests is that in Z3 Ints are arbitary sized data structures, whereas Reals are Doubles. This means Reals tend to run faster than Ints.

Temp variables produced by LLVM are converted back to their proper identifiers and a number appended to the end of the variable name in order to make the code SSA. This number increments every time a variable is encountered. The current counts stored in a map called `ssa` with the base variable name used as a key. Constants declared in the model do not have a number appended to them as their values should not change.

The compiler first converts the tree into a slice of rule types. Rule types include `assrt` (assert and assume statements), `infix`, `ite` (if then else), `invariant` (subnode of assrt nodes), and `wrap`/`vwrap`/`wrapGroup` which are used to wrap various static values.

The way the Go library parses LLVM IR is a little wacky, particularly around conditionals, so rules are tagged with branches if they appear in a conditional statement in order to help the SMT generation.

Both steps taken in parallel and conventional conditionals require a phi variable to store the resulting value from the winning branch. For this reason when the compiler encounters such a situation it does another pass to determine at what state each variable at each branch ended up with. If then generates a rule for a new variable to cap the parallel function or conditional.

Fault assertions are negated by the compiler so that Z3 finds an invariant and disproves the model. Assumptions are left the way they are in order to eliminate or filter to certain states.

SMTLib2 requires that all variables be initialized separately from their use, so as Fault generates rules it stores them in four groups: `g.inits`, `g.constants`, `g.rules`, and `g.asserts`. The final part of the SMT conversion is writing all these rules to strings starting with the variable initializations, then all the constant declarations, then all the steps of the model, and finally all the assertions.

### Execution (SMTLib2 Syntax -> SMTLib2 Solution -> Model Results)
Once the SMTLib has been generated, Fault runs this code through Z3 and parses the solution (also in SMTLib2) so that the command line tool can display the results for the user. The `responses.go` is a listener file for SMTLib2 necessary to parse the solution returned by Z3. For each Uncertain variable Fault generates a normal distribution from the given Sigma and Mu, then adds the probability of each value returned by Z3 occuring. 
