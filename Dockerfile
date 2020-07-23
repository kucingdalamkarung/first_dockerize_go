FROM golang:alpine

RUN apk update && apk upgrade && \
    apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY . /app

RUN go build -o app main.go

ENTRYPOINT ./app