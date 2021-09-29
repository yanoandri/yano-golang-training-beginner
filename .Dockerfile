FROM golang:alpine AS builder 

RUN apk update

RUN mkdir /myapp/

WORKDIR /myapp/

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o start /myapp/server.go

FROM alpine

COPY --from=builder /myapp/start /myapp/start
COPY --from=builder /myapp/entrypoint.sh /myapp/entrypoint.sh

RUN chmod +x /myapp/entrypoint.sh