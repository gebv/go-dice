GITHASH=`git log -1 --pretty=format:"%h" || echo "???"`

CURDATETS=`date -u +%FT%T`

VERSION=v0.1

APPVERSION=${VERSION}-${GITHASH}
BUILDDATE=${CURDATETS}

build:
	go build -v -ldflags "-X main.VERSION=${APPVERSION} -X main.BUILDATE=${BUILDDATE}" -o ./bin/dice ./dice/main.go ./dice/collections.go

manual-test: build
	go generate ./examples

test:
	go test -v -count 1 -race ./...
