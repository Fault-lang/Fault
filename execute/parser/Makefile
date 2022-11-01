.PHONY: all smtparser java gui clean

all: smtparser java

java:
	@command -v antlr >/dev/null 2>&1 || { echo >&2 "I require antlr but it's not installed or not in PATH.  Aborting."; exit 1; }
	antlr SMTLIBv2.g4
	javac -classpath /usr/local/Cellar/antlr/4.11.1/antlr-4.11.1-complete.jar *.java


gui:
	@command -v antlr >/dev/null 2>&1 || { echo >&2 "I require antlr but it's not installed or not in PATH.  Aborting."; exit 1; }
	grun SMTLIBv2 start -gui

smtparser:
	@command -v antlr >/dev/null 2>&1 || { echo >&2 "I require antlr but it's not installed or not in PATH.  Aborting."; exit 1; }
	antlr SMTLIBv2.g4 -Dlanguage=Go -visitor -o ../parser

clean:
	rm -rf *.class *.java *.tokens *.interp *.go 