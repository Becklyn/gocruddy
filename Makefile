go-docker-run-base := docker run --rm --volume $(PWD)/.docker/go/cache:/root/.cache/go-build --volume $(PWD)/.docker/go/mod:/mod --volume $(PWD)/.docker/go/build:/build --volume $(PWD):/current --env CGO_ENABLED=0 --env GO111MODULE=on --env GOOS=linux --env GOARCH=amd64 --env GOPATH=/mod --workdir=/current golang:1.16-alpine

TESTPACKAGES = `go list ./... | grep -v test`

.PHONY: help
help: ## Show this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"; printf "\Targets:\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m	 %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: setup
setup: ## Setup the development environment
	mkdir -p ./.docker
	mkdir -p ./.docker/go
	mkdir -p ./.docker/go/mod
	mkdir -p ./.docker/go/build
	mkdir -p ./.docker/go/cache

.PHONY: download
download: ## Download all needed dependencies
	$(go-docker-run-base) go mod download

.PHONY: install
install: ## Install dependencies: make install MOD=github.com/your/module
	$(go-docker-run-base) go get -u $(MOD)

.PHONY: tidy
tidy: ## Tidy up go dependencies
	$(go-docker-run-base) go mod tidy

.PHONY: test
test: ## Test the application
	$(go-docker-run-base) go test `go list ./... | grep -v test` -count=1 -timeout 15s

.PHONY: cover
cover: ## Run test coverage (needs a local go installation)
	go test $(TESTPACKAGES) -coverprofile=coverage.out -covermode=atomic -count=1 -timeout 25s
	go tool cover -html=coverage.out
	go tool cover -func coverage.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}'
	rm coverage.out
