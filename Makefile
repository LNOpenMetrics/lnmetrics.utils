CC=go
FMT=gofmt
NAME=lnmetrics.utils
BASE_DIR=/script
OS=linux
ARCH=386
ARM=

default: fmt lint check

fmt:
	$(CC) fmt ./...

lint:
	golangci-lint run

check:
	$(CC) test -v ./...

check-full:
	richgo test ./... -v

doc:
	$(CC) doc .
