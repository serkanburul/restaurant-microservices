MAILER_BINARY=mailer
RESERVATION_BINARY=reservation-service
SUBSCRIPTION_BINARY=subscription-service

build_mailer:
	cd mail-service && env GOOS=linux CGO_ENABLED=0 go build -o ${MAILER_BINARY} ./
	@echo "Done!"

build_reservation:
	cd reservation-service && GOOS=linux CGO_ENABLED=0 go build -o ${RESERVATION_BINARY} ./
	@echo "Done!"

build_subscription:
	cd subscription-service && GOOS=linux CGO_ENABLED=0 go build -o ${SUBSCRIPTION_BINARY} ./
	@echo "Done!"

front_end:
	docker-compose down frontend
	docker-compose up --build -d frontend

up_build: build_mailer build_reservation build_subscription
	docker-compose down
	docker-compose up --build -d
	@echo "Done!"

up:
	docker-compose up

down:
	docker-compose down
	@echo "Done!"