# syntax=docker/dockerfile:1
FROM golang:1.17

WORKDIR /go/src/github.com/fault-lang/fault/

COPY . .

RUN apt-get update && \
apt-get -y upgrade && \
apt-get install -y ca-certificates gcc

RUN go mod download

ENV SOLVERCMD=""
ENV SOLVERARG="" 

RUN go build -o fcompiler .


