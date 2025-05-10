MAILER_BINARY=mailer
RESERVATION_BINARY=reservation-service
SUBSCRIPTION_BINARY=subscription-service

front_end:
	docker-compose down frontend
	docker-compose up --build -d frontend

build_up:
	docker-compose down
	docker-compose up --build -d
	@echo "Done!"

up:
	docker-compose up

down:
	docker-compose down
	@echo "Done!"
