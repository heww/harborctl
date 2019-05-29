# Makefile for a harbor project
#
# Targets:
#
# build: 	build harbor photon images
# clean:	clean images

# common
SHELL := /bin/bash
BUILDPATH=$(CURDIR)
MAKEPATH=$(BUILDPATH)/make
MAKEDEVPATH=$(MAKEPATH)/dev
SRCPATH=./src
TOOLSPATH=$(CURDIR)/tools
SEDCMD=$(shell which sed)
WGET=$(shell which wget)

# docker parameters
DOCKERCMD=$(shell which docker)
DOCKERBUILD=$(DOCKERCMD) build --pull
DOCKERRMIMAGE=$(DOCKERCMD) rmi
DOCKERIMASES=$(DOCKERCMD) images

# photon dockerfile
DOCKERFILEPATH=$(MAKEPATH)/photon

DOCKERFILEPATH_HARBORCTL=$(DOCKERFILEPATH)/harborctl
DOCKERFILENAME_HARBORCTL=Dockerfile
DOCKERIMAGENAME_HARBORCTL=heww/harborctl

DOCKERFILEPATH_MULTI_FILE_SWAGGER=$(DOCKERFILEPATH)/multi-file-swagger
DOCKERFILENAME_MULTI_FILE_SWAGGER=Dockerfile
DOCKERIMAGENAME_MULTI_FILE_SWAGGER=heww/multi-file-swagger

_build_harborctl:
	@echo "building harborctl container for photon..."
	@$(DOCKERBUILD) -f $(DOCKERFILEPATH_HARBORCTL)/$(DOCKERFILENAME_HARBORCTL) -t $(DOCKERIMAGENAME_HARBORCTL):$(VERSIONTAG) .
	@echo "Done."

_build_multi_file_swagger:
	@echo "building multi-file-swagger container..."
	@$(DOCKERBUILD) -f $(DOCKERFILEPATH_MULTI_FILE_SWAGGER)/$(DOCKERFILENAME_MULTI_FILE_SWAGGER) -t $(DOCKERIMAGENAME_MULTI_FILE_SWAGGER):$(VERSIONTAG) .
	@echo "Done.";

build: _build_harborctl

cleanimage:
	@echo "cleaning image for photon..."

.PHONY: clean
clean: cleanimage

