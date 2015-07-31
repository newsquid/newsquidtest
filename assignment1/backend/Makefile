.PHONY: run clean

default: run

run: backend
	./backend

backend: *.go api/*.go deps
	go build

deps: .deps

.deps:
	-go get -t ./...
	touch .deps

clean:
	-@rm -f backend .deps
