FROM golang:1.22.5

ARG APP

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg
CMD sleep 10000
RUN CGO_ENABLED=0 GOOS=linux go build -o /svc cmd/$APP/main.go

COPY docker.config.yaml ./local.yaml

CMD ["/svc"]