GOPATH ?= $(HOME)/go
BINDIR ?= $(GOPATH)/bin
NAMESPACE ?= https://github.com/Delisa-sama/$(APP_NAME)\
TMPDIR ?= $(shell dirname $$(mktemp -u))
APP_NAME ?= grpc-boilerplate
COVERAGE_OUT ?= /tmp/$(APP_NAME)-coverage.out
API_VER ?= v1

tools:
	@echo Installing tools from tools.go
	@grep '_ "' tools.go | grep -o '"[^"]*"' | xargs -tI % go install %

test:
	go test ./... -coverprofile=$(COVERAGE_OUT)
	go tool cover -func=$(COVERAGE_OUT) | grep ^total

$(COVERAGE_OUT):
	$(MAKE) test

cover:
	go tool cover -html=$(COVERAGE_OUT)
	rm -f $(COVERAGE_OUT)

benchmark:
	go test ./... -short -bench=. -run="Benchmark*"

lint: tools
	golint -set_exit_status ./...

vet:
	go vet ./...

fmt:
	test -z $$(for d in $$(go list -f {{.Dir}} ./...); do gofmt -e -l -w $$d/*.go; done)

imports: tools
	test -z $$(for d in $$(go list -f {{.Dir}} ./...); do goimports -e -l -local $(NAMESPACE) -w $$d/*.go; done)

check: fmt imports vet lint test

models: tools
	sqlboiler psql --no-tests --no-context

dependencies:
	go mod tidy

migrate:
	sql-migrate up

build: service-build gateway-build

service-build:
	go build ./service/main.go

gateway-build:
	go build ./gateway/gateway.go

service-proto: tools
	protoc -I. --go_out=plugins=grpc,paths=source_relative:. api/$(API_VER)/*.proto

docs-proto: tools
	protoc -I. --swagger_out=logtostderr=true,grpc_api_configuration=api/$(API_VER)/api.yaml:. api/$(API_VER)/*.proto

gateway-proto: tools
	protoc -I. --grpc-gateway_out=logtostderr=true,paths=source_relative,grpc_api_configuration=api/$(API_VER)/api.yaml:. api/$(API_VER)/*.proto

proto: service-proto gateway-proto docs-proto

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean:
	rm -f main gateway coverage.out api/$(API_VER)/*.pb.go api/$(API_VER)/*.pb.gw.go api/$(API_VER)/*.swagger.json
	rm -rf models
