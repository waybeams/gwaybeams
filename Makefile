# Build file for Waybeams workspace
# As of 8/7/2018, this project relies on the updated build of Go Modules and
# a binary built from at least go1.11beta2. In order for this makefile to work,
# `which go` should return a reference to the beta build of the go binary.
# For me, this meant moving the older (v1.10) binary to a name that does not
# match "go" and then symlinking the beta binary to "go" in a folder that's in
# my path.

test:
	go test ./... | ./script/colorize

test-v:
	go test -v ./... | ./script/colorize

bench:
	go test ./... -bench=. | ./script/colorize
