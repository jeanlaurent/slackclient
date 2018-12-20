FROM golang:1.11.1-alpine3.8 as gobuilder
WORKDIR /build
COPY main.go ./
RUN env GOOS=linux GOARCH=amd64 go build -o slack

FROM alpine:3.8
RUN apk add --no-cache ca-certificates
COPY --from=gobuilder /build/slack /usr/bin/slack
CMD ["slack", "hello-world"]