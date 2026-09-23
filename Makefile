.PHONY: build test vet fmt apidiff

build:
	go build ./...

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofmt -l .

# apidiff compares this package's exported API surface against the most
# recent git tag and reports any incompatible changes. Run before
# considering a change to any exported identifier complete. See the
# update-tour-manager-go-client skill for how to handle a reported
# breaking change.
apidiff:
	go run golang.org/x/exp/cmd/gorelease@latest
