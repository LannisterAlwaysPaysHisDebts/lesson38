
.PHONY: proto
proto:
	protoc -I ./ --micro_out=./ --go_out=./ ./proto/cart/*.proto && protoc -I ./ --micro_out=./ --go_out=./ ./proto/cartApi/*.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cart-service ./server/*.go

.PHONY: buildApi
buildApi:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cartApi-api ./api/*.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t cart-service:latest
