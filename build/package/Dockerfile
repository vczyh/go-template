# Build
FROM golang:1.15.8 AS builder

WORKDIR /myapp
COPY . .

RUN chmod -R +x scripts
RUN scripts/app_build.sh

# Run
FROM scratch

COPY config.yml /etc/myapp/

WORKDIR /
COPY --from=builder /myapp/bin/myapp .

EXPOSE 8080

ENTRYPOINT ["/myapp"]
CMD ["--config","/etc/myapp/config.yml"]