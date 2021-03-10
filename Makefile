GOPACKR := $(shell command -v packr2 2> /dev/null)
currentDir = $(shell pwd)
releasesOutputDir = ${currentDir}/releases/$(date +'%F')


## installation
install:
ifndef GOPACKR
	@echo "==> installing go-bindata"
    @go get -u github.com/gobuffalo/packr/v2/...
endif
	@echo "==> installing go dependencies"
	@go mod download
.PHONY: install

run:
	@echo "==> running job scrapper"
	@${currentDir}/scripts/run.sh
.PHONY: run


runquery:
	@echo "==> running job scrapper with query"
	@${currentDir}/scripts/runquery.sh
.PHONY: runquery

buildwindows:
	@echo "==> building job scrapper for windows"
	@${currentDir}/scripts/buildwin.sh
.PHONY: build

build:
	@echo "==> building job scrapper"
	@${currentDir}/scripts/build.sh
.PHONY: build

git:
	@echo "==> adding git tracked files"
	@git add -u
	@git commit
	@echo "==> pushing to git remote"
	@git push origin
.PHONY: git

clean:
	@go clean --cache
	@go mod tidy
	@git clean -f
.PHONY: clean
