###########################################################
# build script
###########################################################

# Operating System (darwin or linux)
PLATFORM:=$(shell uname | tr A-Z a-z)
ARCH=amd64
PROJECT_ROOT=$(shell git rev-parse --show-toplevel)
PROJECT_NAME="FindingYou"

# GOLANG
GOOS=$(PLATFORM)
GOARCH=$(ARCH)

GOLANG_VERSION=1.10
GOLANG_SRC=tmp/golang

GOLANG_PATH=lib/go-$(GOLANG_VERSION)
GOLANG_BIN=$(GOLANG_PATH)/bin
GOLANG_BINARY=$(GOLANG_BIN)/go

.PHONY: test test-w dev-install build lint clean

build: bin/$(PROJECT_NAME)

# Run all tests
test: $(NODE)
	$(MOCHA) --reporter dot $(TEST_FILES)

test-w: $(NODE)
	$(MOCHA) --reporter dot $(TEST_FILES) -w

publish: clean build
	npm publish

dist/indexcards.js: client.js src/*
	$(WEBPACK) --devtool source-map --config $(WEBPACK_CLIENT_CONFIG) client.js dist/indexcards.js

dist/indexcards.min.js: client.js src/*
	$(WEBPACK) --optimize-minimize --config $(WEBPACK_CLIENT_CONFIG) client.js dist/indexcards.min.js

dist/indexcards.min.gz:
	gzip --best -c dist/indexcards.min.js > dist/indexcards.min.gz

dist/express.js:
	$(WEBPACK) --config $(WEBPACK_SERVER_CONFIG) express.js dist/express.js
	
lint:
	$(ESLINT) --config $(PROJECT_ROOT)/.eslintrc.json .

module-install: 
	$(NPM) install

integrate: clean lint build test

clean: 
	rm -rf dist
	rm -rf tmp
	rm -f .tmp-view.html

# Intall development dependencies (OS X and Linux only)
dev-install: $(GOLANG_BINARY)

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

# Install npm dependencies
$(NODE_MODULES_BIN): $(PROJECT_ROOT)/package.json
	$(NPM) install --development

