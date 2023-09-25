FROM golang:1.21.1-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o binary

ENTRYPOINT ["/app/binary"]
