###########################################################
# build script
###########################################################

# Operating System (darwin or linux)
PLATFORM:=$(shell uname | tr A-Z a-z)
ARCH=amd64
PROJECT_ROOT=$(shell git rev-parse --show-toplevel)
PROJECT_NAME='FindingYou'

# Google Depot Tools
TOOLS_URL="https://chromium.googlesource.com/chromium/tools/depot_tools.git"
DEPOT_TOOLS=lib/depot_tools

# GOLANG
GOOS=$(PLATFORM)
GOARCH=$(ARCH)

GOLANG_VERSION=1.10
GOLANG_SRC=tmp/golang

# TODO(lbayes): Extract this duplicate garbage to a simpler config file
ASSERT_URL=github.com/waybeams/assert
ASSERT_PATH=./vendor/src/$(ASSERT_URL)
GLFW_URL=github.com/go-gl/glfw/v3.2/glfw
GLFW_PATH=./vendor/src/$(GLFW_URL)
GOGL_URL=github.com/go-gl/gl/v4.1-core/gl
GOGL_PATH=./vendor/src/$(GOGL_URL)
GOMOBILE_URL=golang.org/x/mobile/cmd/gomobile
GOMOBILE_PATH=./vendor/src/$(GOMOBILE_URL)
GOXJS_URL=github.com/goxjs/gl
GOXJS_PATH=./vendor/src/$(GOXJS_URL)
XID_URL=github.com/rs/xid
XID_PATH=./vendor/src/$(XID_URL)
NANO_URL=github.com/shibukawa/nanovgo
NANO_PATH=./vendor/src/$(NANO_URL)
EASE_URL=github.com/fogleman/ease
EASE_PATH=./vendor/src/$(EASE_URL)
CLOCK_URL=github.com/benbjohnson/clock
CLOCK_PATH=./vendor/src/$(CLOCK_URL)

GOLANG_PATH=lib/go-$(GOLANG_VERSION)
GOLANG_BIN=$(GOLANG_PATH)/bin
GOLANG_BINARY=$(CURDIR)/$(GOLANG_BIN)/go
GOLANG_TEST_BINARY=./script/gotest-color

# TEST_FILES_EXPR=./src/...

.PHONY: test test-w dev-install build lint clean libraries

# Run linter
lint: $(GOLANG_BINARY)
	# NOTE: Ignoring "should have comment or be unexported" lint warning.
	golint ./src/... | grep -v unexported

# Run all tests
test: $(GOLANG_BINARY) $(GOLANG_TEST_BINARY)
	@echo "-------------------------------------------------------------------------------"
	$(GOLANG_TEST_BINARY) test ./src/... 

# Run all tests with verbose mode
test-v: $(GOLANG_BINARY) $(GOLANG_TEST_BINARY)
	@echo "-------------------------------------------------------------------------------"
	$(GOLANG_TEST_BINARY) test -v ./src/...

# Run all benchmarks
bench: $(GOLANG_BINARY)
	$(GOLANG_BINARY) test -bench=. ./src/...

# Run the application binary
run: $(GOLANG_BINARY)
	$(GOLANG_BINARY) run ./examples/todomvc/main.go

# Build a static binary for current platform
build: $(GOLANG_BINARY)
	$(GOLANG_BINARY) build -o out/todomvc-debug examples/todomvc/main.go
	$(GOLANG_BINARY) build -ldflags="-s -w" -o out/todomvc examples/todomvc/main.go
	ls -la out/

clean: 
	rm -rf dist
	rm -rf tmp
	rm -rf out
	rm -rf .gocache

libraries: $(GOGL_PATH) $(ASSERT_PATH) $(GLFW_PATH) $(GOMOBILE_PATH) $(XID_PATH) $(NANO_PATH) $(EASE_PATH) $(CLOCK_PATH)

# Intall development dependencies (OS X and Linux only)
dev-install: $(GOLANG_BINARY) libraries

# Download and unpack the Golang binaries into lib/.
$(GOLANG_BINARY):
	# Download sources
	mkdir -p tmp
	wget -O tmp/go.src.tar.gz "https://dl.google.com/go/go1.10.src.tar.gz"
	# Unpack source files
	mkdir -p $(GOLANG_SRC)
	tar -xvf tmp/go.src.tar.gz -C $(GOLANG_SRC) --strip 1
	# Build from source
	cd $(GOLANG_SRC)/src && GOOS=$(GOOS) GOARCH=$(GOARCH) ./bootstrap.bash
	mkdir -p lib
	mv tmp/go-$(GOOS)-$(GOARCH)-bootstrap $(GOLANG_PATH)
	touch $(GOLANG_BINARY)
	rm -rf tmp

$(DEPOT_TOOLS):
	mkdir -p $(DEPOT_TOOLS)
	git clone $(TOOLS_URL) $(DEPOT_TOOLS)

# Deal with library dependencies
vendor:
	mkdir -p vendor/src

$(ASSERT_PATH): vendor
	cd vendor/; $(GOLANG_BINARY) get -u -v $(ASSERT_URL)
	touch $(ASSERT_PATH)

$(GLFW_PATH): vendor
	cd vendor/; $(GOLANG_BINARY) get -u -v $(GLFW_URL)
	touch $(GLFW_PATH)

$(GOGL_PATH): vendor
	cd vendor/; $(GOLANG_BINARY) get -u -v $(GOGL_URL)
	touch $(GOGL_PATH)

$(GOMOBILE_PATH): vendor
	cd vendor/; $(GOLANG_BINARY) get -u -v $(GOMOBILE_URL)
	touch $(GOMOBILE_PATH)

$(GOXJS_PATH): vendor
	cd vendor/; $(GOLANG_BINARY) get -u -v $(GOXJS_URL)
	touch $(GOXJS_PATH)

$(NANO_PATH): vendor
	cd vendor/; $(GOLANG_BINARY) get -u -v $(NANO_URL)
	touch $(NANO_PATH)

$(EASE_PATH): vendor
	cd vendor/; $(GOLANG_BINARY) get -u -v $(EASE_URL)
	touch $(EASE_PATH)

$(CLOCK_PATH): vendor
	cd vendor/; $(GOLANG_BINARY) get -u -v $(CLOCK_URL)
	touch $(CLOCK_PATH)

$(XID_PATH): vendor
	cd vendor/; $(GOLANG_BINARY) get -u -v $(XID_URL)
	touch $(XID_PATH)

ci-install:
	go get -u -v $(GLFW_URL)
	go get -u -v $(GOGL_URL)
	go get -u -v $(GOMOBILE_URL)
	go get -u -v $(NANO_URL)
	go get -u -v $(EASE_URL)
	go get -u -v $(CLOCK_URL)
	go get -u -v $(XID_URL)

# Run all tests for Circle CI
ci-test:
	go test -v ./src/...

ci-build:
	go build -o out/main-debug examples/todomvc/src/main.go
	go build -ldflags="-s -w" -o out/main examples/todomvc/src/main.go
