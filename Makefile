PROTO_INCLUDE_PATH ?= /Users/warren.harper/.homebrew/include
GOPATH ?= $$GOPATH
CWD ?= $$(pwd)

.DEFAULT_GOAL := help

.PHONY: gen-grpc 
gen-grpc: ## Generate grpc specific code from the proto files
	protoc \
		-I ./api \
		-I $(PROTO_INCLUDE_PATH) \
		--go_out=plugins=grpc:pkg \
		config/config.proto

.PHONY: gen-grpc-gateway
gen-grpc-gateway: ## Generate grpc-gateway code from the proto file and api configuration file
	protoc \
		-I ./api \
		-I $(PROTO_INCLUDE_PATH) \
		-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true,grpc_api_configuration=$(CWD)/api/config/config.yml:pkg \
		config/config.proto

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
