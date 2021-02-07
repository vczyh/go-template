# Build
FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go env -w GOPROXY=https://goproxy.io,direct

WORKDIR /build
COPY . .
RUN go build -o . ./cmd/myapp

# Run
FROM scratch

ENV APP_HOST=zhangeek.com

COPY .env /
COPY --from=builder /build/myapp /
EXPOSE 8080
ENTRYPOINT ["/myapp"]