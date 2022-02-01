FROM golang:1.17-alpine AS builder

COPY . /github.com/goboden/lf-notifier/
WORKDIR /github.com/goboden/lf-notifier/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/bot ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/goboden/lf-notifier/.bin/bot .

EXPOSE 80

CMD ["./bot"]