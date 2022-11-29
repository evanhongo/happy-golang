PROJECT="happy-golang"
TIME=$$(date +%Y-%m-%d_%H:%M)
COMMIT=$$(git rev-parse --short HEAD)
VERSION=$$( if $$(git describe --abbrev=0 > /dev/null 2>&1); then echo $$(git describe --abbrev=0) ; else echo $(COMMIT) ; fi )

##@ Info
.PHONY: version
version:  ## Show version
	echo $(VERSION)


##@ Docs
.PHONY: docs
docs: ## Generate swagger docs
	swag i -d ./api,./entity -g main.go --propertyStrategy camelcase -o ./api/docs


##@ Proto
.PHONY: proto
proto: ## Generate protobuf related files
	protoc --twirp_out=. --go_out=. rpc/job/job.proto


##@ Testing
test:  ## Run test
	go test -v -cover ./...


##@ Build
.PHONY: build
build: ## Build binary
	GOOS=linux GOARCH=amd64 go build -o app.exe ./cmd/server


##@ Run
.PHONY: run
run:  ## Run main service
	go run cmd/server/*.go

##@ Client
.PHONY: protobuf client
client:  ## 
	go run cmd/client/main.go

##@ Help
.PHONY: help
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help