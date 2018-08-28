# Defaults
GOOS ?= darwin
GOARCH ?= $(CURRENT_ARCH)

build-setup: clean
	go get -u github.com/kardianos/govendor

clean:
	rm -rf bin
	mkdir bin

build-fsite:
	cd cmd/f-site && env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ../../bin/f-site-$(GOOS)-$(GOARCH)

run-fsite: build-fsite
	env GOOS=$(GOOS) GOARCH=$(GOARCH) bin/f-site-$(GOOS)-$(GOARCH)
