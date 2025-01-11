FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY reservation-service /app
COPY db/migrations /app/migrations

CMD [ "/app/reservation-service" ]
