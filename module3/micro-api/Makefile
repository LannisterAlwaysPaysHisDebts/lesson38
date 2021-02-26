build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go14 build -o micro-api main.go

dockerBuild:
	docker build -t caot1995/micro-api-test .

.PHONY: build dockerBuild