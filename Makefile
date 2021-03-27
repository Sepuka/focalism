PROGRAM_NAME=focalism

init:
	dep ensure -v

update:
	dep ensure -update

build:
	go build -o $(PROGRAM_NAME)

dependencies:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	dep ensure

tests:
	go test ./...