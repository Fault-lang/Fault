.PHONY: all generate clean

all: generate

generate:
	@command -v grammarinator-generate >/dev/null 2>&1 || { echo >&2 "I require grammarinator but it's not installed or not in PATH.  Aborting."; exit 1; }
	grammarinator-generate -p FaultUnparser.py -l FaultUnlexer.py -r spec -o test_%d.spec -n 3

clean:
	rm -rf *.spec