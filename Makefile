.PHONY: test build

# Source a local .env
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

#################################################################################
# BUILD COMMANDS
#################################################################################
dependencies:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/golang/mock/mockgen@v1.6.0
	npm i -g @redocly/cli@latest

gen: dependencies
	go generate ./...

	docker run -v ${PWD}:/activity  openapitools/openapi-generator-cli generate -i /activity/api/openapi.yaml -g typescript-node -o /activity/client/ts/ --git-user-id jesse0michael --git-repo-id activity --additional-properties=npmName=@jesse0michael/activity,npmVersion=1.0.0

	redocly build-docs api/openapi.yaml -o docs/index.html 

build-cli: 
	go build -o ./bin/activity ./cmd/activity
	
#################################################################################
# TEST COMMANDS
#################################################################################
test:
	go test -cover ./... 

lint:
	golangci-lint run ./...

cover:
	go test -coverpkg ./internal/... -coverprofile coverage.out ./... && go tool cover -html=coverage.out

vuln: dependencies
	govulncheck -test ./...

behavior: build-cli
	go test ./test/...  -tags=behavior
