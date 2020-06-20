FROM golang:alpine AS builder

# download git so 'go get' works, then copy src to builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/diffuse/gloss
COPY . .

# get dependencies and build
RUN go get -d ./... && CGO_ENABLED=0 GOOS=linux go install ./...

# copy binary from builder
FROM alpine
COPY --from=builder /go/bin/gloss .

ENTRYPOINT ["/gloss"]