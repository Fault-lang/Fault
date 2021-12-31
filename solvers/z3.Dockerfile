FROM ghcr.io/z3prover/z3:ubuntu-20.04-bare-z3-sha-5a77c30

WORKDIR /fault

COPY --from=0 /go/src/github.com/fault-lang/fault/fcompiler ./

RUN apt-get update && \
apt-get -y upgrade && \
apt-get install -y llvm

ENV SOLVERCMD="z3"
ENV SOLVERARG="-in"
ENV FAULT_HOST="/host"

# set entrypoint
ENTRYPOINT [ "./fcompiler"]

CMD [ "-mode=check", "-input=fspec",""]
