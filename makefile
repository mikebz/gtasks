.DEFAULT_GOAL = test

vmajor   := 0
vminor   := 1
version  := $(vmajor).$(vminor).$(shell git rev-list --count HEAD)
buildnum := $(shell git rev-parse --short HEAD)

output   := out

name     := gtask
package  := github.com/mikebz/$(name)
packages := $(shell go list ./... | grep -v /vendor/)

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
gobin 	:= $(shell go env GOPATH)/bin
else
gobin	:= $(shell go env GOBIN)
endif

.PHONY: build
build:
	echo building version $(version)
	echo build number $(buildnum)
	mkdir -p $(output)
	go build -o $(output)/$(name) \
		-ldflags "-X github.com/mikebz/kgr/cmd.Version=$(version) -X 'github.com/mikebz/kgr/cmd.Build=$(buildnum)'"

.PHONY: install
install:
	echo installing version $(version)
	echo installing build number $(buildnum)
	go install -ldflags "-X github.com/mikebz/kgr/cmd.Version=$(version) -X 'github.com/mikebz/kgr/cmd.Build=$(buildnum)'"

.PHONY: test
test:
	go test -v $(packages)

.PHONY: bench
bench:
	go test -bench=. -v $(packages)

.PHONY: lint
lint: lint-license-headers
	go vet -v $(packages)

.PHONY: check
check: lint test

.PHONY: clean
clean:
	git clean -xddff

"$(gobin)/addlicense":
	go install github.com/google/addlicense@v1.1.1

.PHONY: license-headers
license-headers: "$(gobin)/addlicense"
	"$(gobin)/addlicense" -v -c "Mike Borozdin" -f LICENSE_TEMPLATE -ignore=vendor/** -ignore=out/** . 2>&1 | sed '/ skipping: / d'

.PHONY: lint-license-headers
lint-license-headers: "$(gobin)/addlicense"
	"${gobin}/addlicense" -check -ignore=vendor/** -ignore=out/** . 2>&1 | sed '/ skipping: / d'