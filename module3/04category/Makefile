GOPATH := $(shell go env GOPATH)
.PHONY: proto
proto:
	protoc -I ./ --go_out=./ --micro_out=./ ./proto/category/*.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o category-service ./server/*.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: dockerBuild
dockerBuild:
	docker build -t caot1995/category-service:latest .
