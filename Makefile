GOBINDATA := $(shell command -v go-bindata 2> /dev/null)
currentDir = $(shell pwd)
releasesOutputDir = ${currentDir}/releases/$(date +'%F')


## installation
install:
ifndef GOBINDATA
	@echo "==> installing go-bindata"
	@go get -u github.com/go-bindata/go-bindata/...
endif
	@echo "==> installing go dependencies"
	@go mod download
.PHONY: install

run:
	@echo "==> running job scrapper"
	@${currentDir}/scripts/run.sh
.PHONY: run

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
