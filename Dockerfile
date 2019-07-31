FROM golang AS builder
WORKDIR /go/src/app
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GO111MODULE=on GOPROXY=https://goproxy.io
RUN go build -a -installsuffix cgo -o github-missing-api .

FROM scratch
COPY --from=builder /go/src/app .
# ENV GIN_MODE=release
EXPOSE 8000
ENTRYPOINT ["./github-missing-api"]