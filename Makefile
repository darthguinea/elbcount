export APP=$(shell basename $(CURDIR))
export GOPATH=${PWD}
export GOBIN=${PWD}

all:
	go get
	go build

install:
	cp ${APP} ${INSTALL_PATH}

clean:
	go clean

clean-all:
	go clean
	rm -rf src/github.com
