.PHONY: print print1 touchfile build moveBuild

all: helloworld type expression

helloworld:
	go build src/helloworld.go
	mv helloworld bin/

type:
	@go build src/type.go
	@mv type bin/

expression:
	@go build src/expression.go
	@mv expression bin/

clean:
	rm -rf bin/*
