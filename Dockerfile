FROM golang:1.12.7-alpine

WORKDIR /go/src/app
COPY . .

RUN go build main.go
ENV GIN_MODE=release
EXPOSE 8080

ENTRYPOINT ["./main"]