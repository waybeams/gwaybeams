# Build file for Waybeams workspace
# As of 8/7/2018, this project relies on the updated build of Go Modules and
# a binary built from at least go1.11beta2. In order for this makefile to work,
# `which go` should return a reference to the beta build of the go binary.
# For me, this meant moving the older (v1.10) binary to a name that does not
# match "go" and then symlinking the beta binary to "go" in a folder that's in
# my path.

dev-install:
	go build ./...

clean:
	rm -rf bin/*

run:
	go run ./examples/todo/cmd/desktop/main.go

build:
	go build -ldflags="-s -w" -o bin/desktop ./examples/todo/cmd/desktop/...

serve:
	gopherjs serve examples/todo/cmd/browser/main.go

run-js:
	gopherjs run examples/todo/cmd/browser/main.go

build-js:
	gopherjs build -m -o bin/todo.min.js ./examples/todo/cmd/browser/main.go
	rm -f bin/todo.min.gz
	gzip -c -8 bin/todo.min.js > bin/todo.min.gz
	ls -la bin/

test:
	go test ./... | ./script/colorize

test-v:
	go test -v ./... | ./script/colorize

bench:
	go test ./... -bench=. | ./script/colorize
