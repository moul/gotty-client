# build
FROM            golang:1.15-alpine as builder
RUN             apk add --no-cache git gcc musl-dev make
WORKDIR         /go/src/github.com/moul/gotty-client
COPY            go.* ./
RUN             go mod download
COPY            . ./
RUN             make install

# minimal runtime
FROM            alpine:3.12
COPY            --from=builder /go/bin/gotty-client /bin/
ENTRYPOINT      ["/bin/gotty-client"]
