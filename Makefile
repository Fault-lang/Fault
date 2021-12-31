NAME   := fault-lang/fault
TAG    := $$(git log -1 --pretty=%h)
IMG    := ${NAME}:${TAG}
LATEST := ${NAME}:latest

fault-z3:
	$(shell touch "fault.Dockerfile")
	cat Dockerfile ./solvers/z3.Dockerfile > fault.Dockerfile
	@docker build -t ${NAME}-z3:${TAG} --no-cache -f fault.Dockerfile .
	@docker tag ${NAME}-z3:${TAG} ${NAME}-z3:latest
	$(shell rm "fault.Dockerfile")

# docker run -v ~/:/host:ro fault-lang/fault-z3 -mode=smt -filepath=Fault/smt/testdata/bathtub2.fspec
image:
	@docker build -t ${IMG} .
	@docker tag ${IMG} ${LATEST}
