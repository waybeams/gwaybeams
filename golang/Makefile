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
GLFW=vendor/src/$(GLFW_URL)
GOGL_URL=github.com/go-gl/gl/v4.1-core/gl
GOGL=vendor/src/$(GOGL_URL)

GOLANG_PATH=lib/go-$(GOLANG_VERSION)
GOLANG_BIN=$(GOLANG_PATH)/bin
GOLANG_BINARY=$(GOLANG_BIN)/go
GOLANG_TEST_BINARY=./script/gotest-color

# TEST_FILES_EXPR=./src/...

.PHONY: test test-w dev-install build lint clean

build: out/go-skia-demo

out/go-skia-demo:
	mkdir -p out
	$(GOLANG_BINARY) build -o out/go-skia-demo -a github.com/lukebayes/go-skia-demo

# Run all tests
test: $(GOLANG_BINARY) $(GOLANG_TEST_BINARY)
	@echo "-------------------------------------------------------------------------------"
	$(GOLANG_TEST_BINARY) test -v ./src/...

# Get the requested packages and install into our vendor path
# Call like: 
#    make get ARGS="-u -v github.com/go-gl/glfw/v3.2/glfw"
get: $(GOLANG_BINARY)
	cd vendor && $(GOLANG_BINARY) get ${ARGS}

# Run the application binary
run: $(GOLANG_BINARY)
	$(GOLANG_BINARY) run ./src/github.com/lukebayes/findingyou/findingyou.go

# Build a static binary for current platform
build: $(SKIA_LIB)
	echo "Build complete"

clean: 
	rm -rf dist
	rm -rf tmp
	rm -rf out

# Intall development dependencies (OS X and Linux only)
dev-install: $(GOLANG_BINARY) $(SKIA_LIB) $(GLFW) $(GOGL)

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

# Create the vendor directory
vendor:
	mkdir -p vendor

$(GLFW): vendor
	cd vendor/; ../$(GOLANG_BINARY) get -v $(GLFW_URL)

$(GOGL): vendor
	cd vendor/; ../$(GOLANG_BINARY) get -v $(GOGL_URL)

