# Build file for Waybeams workspace

GOPATH="${HOME}/go:${CURDIR}"
GOPHERJS="${HOME}/go/bin/gopherjs"

dev-install:
	go mod download

clean:
	rm -rf bin/*

test:
	go test ./... | ./script/colorize

test-v:
	go test -v ./... | ./script/colorize

bench:
	go test ./... -bench=. | ./script/colorize


# Example Tasks
run: run-desktop

build: build-desktop build-js

run-desktop:
	go run ./examples/todo/cmd/desktop/main.go

build-desktop:
	go build -ldflags="-s -w" -o bin/desktop ./examples/todo/cmd/desktop/...

serve:
	${GOPHERJS} serve ./examples/todo/cmd/browser/main.go

run-js:
	${GOPHERJS} run ./examples/todo/cmd/browser/main.go

build-js:
	GO111MODULE=on GOPATH=${GOPATH} ${GOPHERJS} build -m -o bin/todo.min.js ./examples/todo/cmd/browser/main.go
	# gopherjs build -m -o bin/todo.min.js ./todo/cmd/browser/main.go
	rm -f bin/todo.min.gz
	gzip -c -8 bin/todo.min.js > bin/todo.min.gz
	ls -la bin/

