package mail

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"reservation-service/infrastructure/mail/proto"
	pb "reservation-service/infrastructure/mail/proto"
	"time"
)

type Mail struct {
	Token           string
	Type            string
	Name            string
	Email           string
	Capacity        int
	ReservationDate string
	TimeSlot        string
}

func SendMail(mail Mail) error {
	conn, err := grpc.NewClient("mail-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer conn.Close()

	client := pb.NewMailServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &pb.SendReservationMailRequest{
		ReservationMailEntry: &proto.ReservationMail{
			To:              mail.Email,
			Token:           mail.Token,
			Name:            mail.Name,
			ReservationDate: mail.ReservationDate,
			TimeSlot:        mail.TimeSlot,
			Type:            mail.Type,
		},
	}

	res, err := client.SendReservationMail(ctx, req)
	if err != nil {
		log.Fatalf("Error while sending mail: %v", err)
	}

	log.Printf("Status: %v, Message: %s", res.Error, res.Message)

	return nil
}
