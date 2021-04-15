# build
FROM            golang:1.16.3-alpine as builder
RUN             apk add --no-cache git gcc musl-dev make
WORKDIR         /go/src/github.com/moul/gotty-client
COPY            go.* ./
RUN             go mod download
COPY            . ./
RUN             make install

# minimal runtime
FROM            alpine:3.13.5
COPY            --from=builder /go/bin/gotty-client /bin/
ENTRYPOINT      ["/bin/gotty-client"]
