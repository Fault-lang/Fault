.PHONY: all java go gui fuzz clean

all: java go fuzz

java:
	@command -v antlr >/dev/null 2>&1 || { echo >&2 "I require antlr but it's not installed or not in PATH.  Aborting."; exit 1; }
	antlr FaultLexer.g4
	antlr FaultParser.g4
	javac -classpath /usr/local/Cellar/antlr/4.8_1/antlr-4.8-complete.jar *.java

go:
	@command -v antlr >/dev/null 2>&1 || { echo >&2 "I require antlr but it's not installed or not in PATH.  Aborting."; exit 1; }
	antlr FaultLexer.g4 -Dlanguage=Go -o ../parser
	antlr FaultParser.g4 -Dlanguage=Go -visitor -o ../parser

gui:
	@command -v antlr >/dev/null 2>&1 || { echo >&2 "I require antlr but it's not installed or not in PATH.  Aborting."; exit 1; }
	grun Fault spec -gui

fuzz:
	@command -v grammarinator-process >/dev/null 2>&1 || { echo >&2 "I require grammarinator but it's not installed or not in PATH.  Aborting."; exit 1; }
	grammarinator-process FaultLexer.g4 FaultParser.g4 -o ../fuzzer --no-actions

clean:
	rm -rf *.class *.java *.tokens *.interp *.go 