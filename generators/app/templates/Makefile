BIN_NAME = <%= appName %>
BIN_PATH = ./build/${BIN_NAME}

.PHONY: docs build run

dep:
	go get -u github.com/swaggo/swag/cmd/swag

run:
	${BIN_PATH}

build:
	go build -tags=jsoniter -o ${BIN_PATH}

docs:
	swag init

cert:
	mkdir cert && generate-certificate.sh

dev: docs build run
