# SMTLib2 Parser
ANTLR Grammar for SMTLib2. Not to be confused with Fault's parser, this parsers the responses from the SMT solver while the model checker is running

If updating from source repo, note that `string` and `String` need to be changed to `string_` and `String_` respectively in order to compile the Go files correctly.

Source Repo:
https://github.com/antlr/grammars-v4/blob/master/smtlibv2/SMTLIBv2.g4