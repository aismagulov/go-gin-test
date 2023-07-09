appname = go-gin-test
version = 1

run:
	go run main.go

build:
	CGO_ENABLED=0 \
	GOOS=linux \
	go build -o bin/main main.go

build-image:
	docker build -t "$(appname):$(version)" .

run-image:
	docker run --name $(appname) -p 8080:8080 $(appname):$(version)