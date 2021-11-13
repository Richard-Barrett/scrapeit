MKFILE_DIR := $(abspath $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST))))))

UNAME_S=$(shell uname -s)
ifeq ($(UNAME_S),Linux)
	GOOS := linux
endif
ifeq ($(UNAME_S),Darwin)
	GOOS := darwin
endif

MOD_PATH := github.com/Mirantis/dataeng

GO_IMAGE := golang:1.16.5

DOCKER_FLAGS := -v $(MKFILE_DIR):/go/$(MOD_PATH) -w /go/$(MOD_PATH) --rm -it
GO_FLAGS     := $(DOCKER_FLAGS) -u $$(id -u):$$(id -g) -e GOCACHE=/tmp/go=cache -e GOOS=$(GOOS)

GO := docker run $(GO_FLAGS) $(GO_IMAGE) go

.PHONY: all
all:
	$(GO) build -o bin/dataengctl $(MOD_PATH)

