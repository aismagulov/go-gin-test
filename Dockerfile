FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY /bin/main .

EXPOSE 8080

ENTRYPOINT ["./main"]