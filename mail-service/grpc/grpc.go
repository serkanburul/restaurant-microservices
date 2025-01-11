package grpc_server

import (
	"google.golang.org/grpc"
	"log"
	"mail-service/proto"
	"net"
)

type MailServer struct {
	proto.UnimplementedMailServiceServer
}

func GRPCListen() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterMailServiceServer(grpcServer, &MailServer{})
	log.Println("grpc server listening on :50051")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
