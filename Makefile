SHELL := /bin/bash
BUILDPATH=$(CURDIR)
MAKEPATH=$(BUILDPATH)/make

# makefile
MAKEFILEPATH_PHOTON=$(MAKEPATH)/photon

# docker parameters
DOCKERCMD=$(shell which docker)
DOCKERTAG=$(DOCKERCMD) tag
DOCKERRMIMAGE=$(DOCKERCMD) rmi
DOCKERIMAGES=$(DOCKERCMD) images

# docker images
VERSIONTAG=dev
DOCKERIMAGENAME_HARBORCTL=heww/harborctl
DOCKERIMAGENAME_MULTI_FILE_SWAGGER=heww/multi-file-swagger

SWAGGER = docker run --rm \
	-v $(shell pwd)/make/swagger/templates:/templates \
	-v $(shell pwd):/go/src/github.com/heww/harborctl \
	-u $(shell id -u):$(shell id -g) \
	-w /go/src/github.com/heww/harborctl \
	quay.io/goswagger/swagger:v0.19.0

MULTI_FILE_SWAGGER = docker run --rm \
	-v $(shell pwd):/go/src/github.com/heww/harborctl \
	-u $(shell id -u):$(shell id -g) \
	-w /go/src/github.com/heww/harborctl \
	${DOCKERIMAGENAME_MULTI_FILE_SWAGGER}:$(VERSIONTAG)

build_multi_file_swagger:
	make -f $(MAKEFILEPATH_PHOTON)/Makefile _build_multi_file_swagger -e VERSIONTAG=$(VERSIONTAG)

deps:
	@if [ "$(shell ${DOCKERIMAGES} -q ${DOCKERIMAGENAME_MULTI_FILE_SWAGGER}:$(VERSIONTAG) 2> /dev/null)" == "" ]; then \
		make build_multi_file_swagger; \
		echo "deps done"; \
	fi

.PHONY: harbor-sdk
harbor-sdk: deps
	- rm -rf pkg/harbor/{models,client}
	${MULTI_FILE_SWAGGER} -o yaml pkg/harbor/spec/index.yaml > pkg/harbor/swagger.yaml
	${SWAGGER} generate client -f pkg/harbor/swagger.yaml --target pkg/harbor --template-dir=/templates

.PHONY: clean
clean:
	- rm -rf pkg/harbor/{models,client}

.PHONY: build
build:
	make -f $(MAKEFILEPATH_PHOTON)/Makefile build \
	-e VERSIONTAG=$(VERSIONTAG)
