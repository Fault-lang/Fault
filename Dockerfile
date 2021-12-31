# syntax=docker/dockerfile:1
FROM ghcr.io/z3prover/z3:ubuntu-20.04-bare-z3-sha-5a77c30

WORKDIR /fault

COPY . .

RUN apt-get update && \
apt-get -y upgrade && \
apt-get install -y --no-install-recommends golang-1.16-go && \
apt-get install -y ca-certificates gcc llvm

ENV PATH="${PATH}:/usr/lib/go-1.16/bin"

RUN go mod download

RUN go build -o ./bin/fault .

# set entrypoint
ENTRYPOINT [ "./bin/fault"]

CMD [ "-mode=check", "-input=fspec",""]



