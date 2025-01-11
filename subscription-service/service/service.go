package service

import (
	"context"
	"subscription-service/db"
	"subscription-service/infrastructure/mail"
)

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateSubscription(email string) error {
	ctx := context.Background()
	err := s.db.CreateSubscription(ctx, email)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SendMail(subject string, body string) error {
	ctx := context.Background()
	emailList, err := s.db.ListSubscriptionEmail(ctx)
	if err != nil {
		return err
	}
	mail.SendMail(mail.Mail{
		To:      emailList,
		Subject: subject,
		Body:    body},
	)

	return nil
}
