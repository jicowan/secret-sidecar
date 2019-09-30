# Use base golang image from Docker Hub
FROM golang:1.12
WORKDIR /src/aws-secrets-manager
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /app -v ./cmd/aws-secrets-manager
ENTRYPOINT ["/app"]
