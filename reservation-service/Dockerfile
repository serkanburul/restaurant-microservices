FROM golang:1.23.1-alpine AS builder

WORKDIR /build

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN env GOOS=linux CGO_ENABLED=0 go build -o reservation-service ./

FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /build/reservation-service /app/reservation-service

COPY db/migrations /app/migrations

CMD ["/app/reservation-service"]
