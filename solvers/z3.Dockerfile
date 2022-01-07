FROM ghcr.io/z3prover/z3:ubuntu-20.04-bare-z3-sha-5a77c30

ARG BUILD_DATE
ARG BUILD_VERSION

WORKDIR /fault

COPY --from=0 /go/src/github.com/fault-lang/fault/fcompiler ./

RUN apt-get update && \
apt-get -y upgrade && \
apt-get install -y llvm

ENV SOLVERCMD="z3"
ENV SOLVERARG="-in"
ENV FAULT_HOST="/host"

# set label info
LABEL org.opencontainers.image.vendor="Fault-lang"
LABEL org.opencontainers.image.authors="Marianne Bellotti" 
LABEL org.opencontainers.image.created=${BUILD_DATE} 
LABEL org.opencontainers.image.version=${BUILD_VERSION}
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.description="Fault using Z3Solver as its engine"


# set entrypoint
ENTRYPOINT [ "./fcompiler"]

CMD [ "-mode=check", "-input=fspec",""]
