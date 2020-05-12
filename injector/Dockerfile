# build stage
FROM golang:1.12-alpine AS build-env

RUN apk add git

RUN mkdir -p /go/src/github.com/jicowan/secret-sidecar/injector
WORKDIR /go/src/github.com/jicowan/secret-sidecar/injector

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY  . .
RUN adduser -D -u 10001 webhook
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' \
    -o secret-sidecar ./cmd/secret-sidecar/*.go

FROM scratch
COPY --from=build-env /go/src/github.com/aws/aws-app-mesh-inject/secret-sidecar .
COPY --from=build-env /etc/passwd /etc/passwd
USER webhook
ENTRYPOINT ["/secret-sidecar"]
