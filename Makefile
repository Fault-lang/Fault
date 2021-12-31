NAME   := fault-lang/fault
TAG    := $$(git log -1 --pretty=%h)
IMG    := ${NAME}:${TAG}
LATEST := ${NAME}:latest
 
image:
	@docker build -t ${IMG} .
	@docker tag ${IMG} ${LATEST}
