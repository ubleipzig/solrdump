SHELL := /bin/bash
TARGETS := solrdump
PKGNAME := solrdump
VERSION := 0.1.15

SEMVER := $(shell echo $(VERSION) | sed 's/^v//')

.PHONY: all
all: $(TARGETS)

$(TARGETS): %: cmd/%/main.go
	go build -ldflags "-w -s -X main.Version=$(VERSION)" -o $@ $<

# Cross-compiled linux/amd64 binary used for packaging (deb/rpm), kept in a
# separate directory so it never overwrites the native dev binary above and
# the host platform (e.g. macOS arm64) does not leak into the package.
build/solrdump: cmd/solrdump/main.go
	@mkdir -p build
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s -X main.Version=$(VERSION)" -o $@ $<

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	go clean
	rm -f $(TARGETS)
	rm -rf build
	rm -f $(PKGNAME)_*.deb
	rm -f $(PKGNAME)-*.rpm

.PHONY: imports
imports:
	goimports -w .

# nfpm-based packaging.
.PHONY: deb
deb: build/solrdump
	SEMVER=$(SEMVER) GOARCH=amd64 nfpm package -p deb -f nfpm.yaml

.PHONY: rpm
rpm: build/solrdump
	SEMVER=$(SEMVER) GOARCH=amd64 nfpm package -p rpm -f nfpm.yaml
