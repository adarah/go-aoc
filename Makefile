.DEFAULT_GOAL=build

build:
	go build

install:
	go install

test:
	go test ./... -v

bench:
	go test -bench=./... -v