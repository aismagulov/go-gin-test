appname = go-gin-test
version = 1

run:
	go run main.go

build-bin:
	export CGO_ENABLED=0 && \
	export GOOS=linux && \
	go build -o bin/main main.go

build-image:
	docker build -t "$(appname):$(version)" .

build: build-bin build-image

run-image:
	docker stop $(appname) && \
	docker rm $(appname) && \
	docker run --name $(appname) -p 8080:8080 $(appname):$(version)