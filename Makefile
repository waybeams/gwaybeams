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
GLFW_URL=github.com/go-gl/glfw/v3.2/glfw
GLFW_PATH=vendor/src/$(GLFW_URL)
GOGL_URL=github.com/go-gl/gl/v4.1-core/gl
GOGL_PATH=vendor/src/$(GOGL_URL)
GOMOBILE_URL=golang.org/x/mobile/cmd/gomobile
GOMOBILE_PATH=vendor/src/$(GOMOBILE_URL)
CAIRO_URL=github.com/golang-ui/cairo
CAIRO_PATH=vendor/src/$(CAIRO_URL)
XID_URL=github.com/rs/xid
XID_PATH=vendor/src/$(XID_URL)
NANO_URL=github.com/shibukawa/nanovgo
NANO_PATH=vendor/src/$(NANO_URL)

GOLANG_PATH=lib/go-$(GOLANG_VERSION)
GOLANG_BIN=$(GOLANG_PATH)/bin
GOLANG_BINARY=$(CURDIR)/$(GOLANG_BIN)/go
GOLANG_TEST_BINARY=./script/gotest-color

# TEST_FILES_EXPR=./src/...

.PHONY: test test-w dev-install build lint clean libraries

# Run linter
lint: $(GOLANG_BINARY)
	golint ./src/...

# Run all tests
test: $(GOLANG_BINARY) $(GOLANG_TEST_BINARY)
	@echo "-------------------------------------------------------------------------------"
	$(GOLANG_TEST_BINARY) test ./src/...

# Run all tests with verbose mode
test-v: $(GOLANG_BINARY) $(GOLANG_TEST_BINARY)
	@echo "-------------------------------------------------------------------------------"
	$(GOLANG_TEST_BINARY) test -v ./src/...

# Run the application binary
run: $(GOLANG_BINARY)
	$(GOLANG_BINARY) run ./src/examples/boxes/main.go

# Direct path to build Cairo for platform-specific debugging
build-cairo:
	$(GOLANG_BINARY) build ./vendor/src/github.com/golang-ui/cairo

# Build a static binary for current platform
build:
	$(GOLANG_BINARY) build -o out/main-debug src/examples/boxes/main.go
	$(GOLANG_BINARY) build -ldflags="-s -w" -o out/main src/examples/boxes/main.go
	ls -la out/

clean: 
	rm -rf dist
	rm -rf tmp
	rm -rf out
	rm -rf .gocache

libraries: $(GOGL_PATH) $(GLFW_PATH) $(GOMOBILE_PATH) $(CAIRO_PATH) $(XID_PATH) $(NANO_PATH)

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
	mkdir -p vendor

$(GLFW_PATH): vendor
	cd vendor/; $(GOLANG_BINARY) get -u -v $(GLFW_URL)
	touch $(GLFW_PATH)

$(GOGL_PATH): vendor
	cd vendor/; $(GOLANG_BINARY) get -u -v $(GOGL_URL)
	touch $(GOGL_PATH)

$(GOMOBILE_PATH): vendor
	cd vendor/; $(GOLANG_BINARY) get -u -v $(GOMOBILE_URL)
	touch $(GOMOBILE_PATH)

$(CAIRO_PATH): vendor
	cd vendor/; $(GOLANG_BINARY) get -u -v $(CAIRO_URL)
	touch $(CAIRO_PATH)

$(XID_PATH): vendor
	cd vendor/; $(GOLANG_BINARY) get -u -v $(XID_URL)
	touch $(XID_PATH)

