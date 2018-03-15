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

# SKIA
SKIA_URL="https://skia.googlesource.com/skia.git"
# SKIA_URL="https://github.com/google/skia.git"
SKIA_SRC=lib/skia
SKIA_LIB=lib/skia/out/Shared/libskia.so

# GOLANG
GOOS=$(PLATFORM)
GOARCH=$(ARCH)

GOLANG_VERSION=1.10
GOLANG_SRC=tmp/golang

GLFW_URL=github.com/go-gl/glfw/v3.2/glfw
GLFW_PATH=vendor/src/$(GLFW_URL)
GOGL_URL=github.com/go-gl/gl/v4.1-core/gl
GOGL_PATH=vendor/src/$(GOGL_URL)
GOMOBILE_URL=golang.org/x/mobile/cmd/gomobile
GOMOBILE_PATH=vendor/src/$(GOMOBILE_URL)

GOLANG_PATH=lib/go-$(GOLANG_VERSION)
GOLANG_BIN=$(GOLANG_PATH)/bin
GOLANG_BINARY=$(CURDIR)/$(GOLANG_BIN)/go
GOLANG_TEST_BINARY=./script/gotest-color

# TEST_FILES_EXPR=./src/...

.PHONY: test test-w dev-install build lint clean libraries

build: out/go-skia

out/go-skia:
	mkdir -p out
	$(GOLANG_BINARY) build -a -o out/go-skia -a src/skia_main.go

# Run all tests
test: $(GOLANG_BINARY) $(GOLANG_TEST_BINARY)
	@echo "-------------------------------------------------------------------------------"
	$(GOLANG_TEST_BINARY) test -v ./src/...

# Run the application binary
run: $(GOLANG_BINARY)
	$(GOLANG_BINARY) run ./src/gnomplate.go

# Build a static binary for current platform
build: $(SKIA_LIB)
	echo "Build complete"

clean: 
	rm -rf dist
	rm -rf tmp
	rm -rf out
	rm -rf .gocache

libraries: $(SKIA_LIB) $(GLFW_PATH) $(GOGL_PATH) $(GOMOBILE_PATH)

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

# Download and build Skia
$(SKIA_SRC): $(DEPOT_TOOLS)
	git clone $(SKIA_URL) $(SKIA_SRC)
	cd $(SKIA_SRC); python tools/git-sync-deps

# Build Skia .SO file
$(SKIA_LIB): $(SKIA_SRC)
	cd $(SKIA_SRC); gn gen out/Shared --args='is_official_build=false is_component_build=true'
	cd $(SKIA_SRC); ninja -C out/Shared

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

