PROGRAM_NAME=focalism

init:
	dep ensure -v

build:
	go build -o $(PROGRAM_NAME)

dependencies:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	dep ensure