
.PHONY: lint

lint:
	golint
	govet
	errcheck
	staticcheck
	unused
	gosimple
