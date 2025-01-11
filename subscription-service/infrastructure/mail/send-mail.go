package mail

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"subscription-service/infrastructure/mail/proto"
	"time"
)

type Mail struct {
	To      []string
	Subject string
	Body    string
}

func SendMail(mail Mail) error {
	conn, err := grpc.NewClient("mail-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer conn.Close()

	client := proto.NewMailServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &proto.SendSubscriptionMailRequest{
		SubscriptionMailEntry: &proto.SubscriptionMail{
			To:      mail.To,
			Body:    mail.Body,
			Subject: mail.Subject,
		},
	}

	res, err := client.SendSubscriptionMail(ctx, req)
	if err != nil {
		log.Fatalf("Error while sending mail: %v", err)
	}

	log.Printf("Status: %v, Message: %s", res.Error, res.Message)

	return nil
}
