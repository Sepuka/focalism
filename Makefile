PROGRAM_NAME=focalism

init:
	dep ensure -v

build:
	go build -o $(PROGRAM_NAME)