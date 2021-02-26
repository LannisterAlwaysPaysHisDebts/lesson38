proto:
	find proto/ -name '*.proto' -exec protoc -I ./ --go_out=./ --micro_out=./ {}  \;

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cartApi-api *.go

docker:
	docker build -t caot1995/cartApi-api:latest .

.PHONY: proto build docker