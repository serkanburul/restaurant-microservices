FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY subscription-service /app
COPY db/migrations /app/migrations

CMD [ "/app/subscription-service" ]