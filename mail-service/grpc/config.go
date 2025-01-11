package grpc_server

import (
	"os"
	"strconv"
)

type SMTP struct {
	HOST string
	PORT int
	USER string
	PASS string
	FROM string
}

func SMTPConfig() (SMTP, error) {
	PORT, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return SMTP{}, err
	}

	return SMTP{
		HOST: os.Getenv("SMTP_HOST"),
		PORT: PORT,
		USER: os.Getenv("SMTP_USER"),
		PASS: os.Getenv("SMTP_PASS"),
		FROM: os.Getenv("FROM_EMAIL"),
	}, nil
}
