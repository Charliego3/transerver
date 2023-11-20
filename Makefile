GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get

BINARY_NAME=transerver
CONFIG_FILE=weaver.toml

all: clean build

build:
	$(GO_BUILD) -o $(BINARY_NAME) -v -ldflags="-X main.MenuIcon=command.square.fill"

clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)

run: clean tidy
	$(GO_BUILD) -o $(BINARY_NAME) -v
	SERVICEWEAVER_CONFIG=${CONFIG_FILE} ./$(BINARY_NAME)

tidy: clean
	${GO_CMD} mod tidy

test:
	$(GO_TEST) -v ./...

deps:
	$(GO_GET)