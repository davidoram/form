# Defaults
GOOS ?= darwin
GOARCH ?= amd64

build-setup: clean
	go get -u github.com/kardianos/govendor

clean:
	rm -rf bin
	mkdir bin

build: build-fsite

test: test-lib test-fsite

build-fsite:
	cd cmd/f-site && env GOOS=$(GOOS) GOARCH=$(GOARCH) govendor build -o ../../bin/f-site-$(GOOS)-$(GOARCH)

test-fsite:
	cd cmd/f-site go test -v ./..

test-lib:
  # Uncomment when we get some tests
	# go test -v ./lib/...

run-fsite: build-fsite
	env GOOS=$(GOOS) GOARCH=$(GOARCH) bin/f-site-$(GOOS)-$(GOARCH)

refresh-javascript-css-libs:
	mkdir -p cmd/f-site/public/css/fonts
	mkdir -p cmd/f-site/public/javascript
	cd cmd/f-site/public/css        && curl --remote-name -L -X GET https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css
	cd cmd/f-site/public/css        && curl --remote-name -L -X GET https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css.map
	cd cmd/f-site/public/css        && curl --remote-name -L -X GET https://unpkg.com/formiojs@latest/dist/formio.full.min.css
	cd cmd/f-site/public/css        && curl --remote-name -L -X GET https://unpkg.com/formiojs@latest/dist/formio.full.min.css.map
	cd cmd/f-site/public/javascript && curl --remote-name -L -X GET https://unpkg.com/formiojs@latest/dist/formio.full.min.js
	cd cmd/f-site/public/javascript && curl --remote-name -L -X GET https://unpkg.com/formiojs@latest/dist/formio.embed.min.js
	cd cmd/f-site/public/css/fonts  && curl --remote-name -L -X GET https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/fonts/fontawesome-webfont.ttf
