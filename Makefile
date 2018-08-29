# Defaults
GOOS ?= darwin
GOARCH ?= $(CURRENT_ARCH)

build-setup: clean
	go get -u github.com/kardianos/govendor

clean:
	rm -rf bin
	mkdir bin

build: build-fsite

test: test-lib test-fsite

build-fsite:
	cd cmd/f-site && env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ../../bin/f-site-$(GOOS)-$(GOARCH)

test-fsite:
	cd cmd/f-site go test -v ./..

test-lib:
  # Uncomment when we get some tests
	# go test -v ./lib/...

run-fsite: build-fsite
	env GOOS=$(GOOS) GOARCH=$(GOARCH) bin/f-site-$(GOOS)-$(GOARCH)
