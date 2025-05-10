FROM golang:1.23.1-alpine AS builder

WORKDIR /build

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN env GOOS=linux CGO_ENABLED=0 go build -o mailer ./

FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /build/mailer /app/mailer

CMD ["/app/mailer"]
