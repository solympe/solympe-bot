REPO := https://github.com/solympe/solympe-bot
IMAGE_NAME := solympe-bot
VERSION ?= dev

 #TODO
build-datalayer: Dockerfile.bot
	docker build \
	-f Dockerfile.bot				\
	--build-arg "VERSION=$(VERSION)" \
	--build-arg "REPO=$(REPO)" \
	-t $(IMAGE_NAME) .