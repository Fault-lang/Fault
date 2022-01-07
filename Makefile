NAME   := fault-lang/fault
TAG    := $$(git log -1 --pretty=%h)
IMG    := ${NAME}:${TAG}
LATEST := ${NAME}:latest
VERSION := $$(git describe --tags --abbrev=0)

fault-z3:
	$(shell touch "fault.Dockerfile")
	cat Dockerfile ./solvers/z3.Dockerfile > fault.Dockerfile
	@docker build -t ${NAME}-z3:${TAG} --no-cache --build-arg BUILD_VERSION=${VERSION} --build-arg BUILD_DATE=$(date) -f fault.Dockerfile .
	@docker tag ${NAME}-z3:${TAG} ${NAME}-z3:latest
	@rm fault.Dockerfile
	@cp fault-lang.sh /usr/local/bin/fault

image:
	@docker build -t ${IMG} .
	@docker tag ${IMG} ${LATEST}
