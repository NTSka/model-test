FROM alpine:latest

WORKDIR /app

COPY . .

RUN apk update --no-cache && apk add --no-cache tzdata bash

ENTRYPOINT ./migrate.sh ./migrations $DSN up