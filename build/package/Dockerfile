# Build
FROM golang:1.15.8 AS builder

#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64
#
#RUN go env -w GOPROXY=https://goproxy.io,direct

WORKDIR /build
COPY . .

RUN chmod -R +x scripts
RUN scripts/app_build.sh

# Run
FROM scratch

WORKDIR /

COPY config.yml .
COPY --from=builder /build/bin/myapp .

EXPOSE 8080

ENTRYPOINT ["/myapp"]
CMD ["--config","config.yml"]