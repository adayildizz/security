package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"

	pb "assignment1/code/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)


type server struct {
	pb.UnimplementedMessengerServer
}

func (s *server) Send( ctx context.Context, req *pb.SendRequest) (*pb.SendReply, error) {
	log.Printf("Message received. %q", req.GetMessage())
	return &pb.SendReply{Status: "Message received."}, nil

}

func main(){

	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")

	if err != nil {
		log.Fatalf("Error in loading certificate: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion: tls.VersionTLS13,
	}

	crds := credentials.NewTLS(tlsConfig)
	grpcServer := grpc.NewServer(grpc.Creds(crds))

	pb.RegisterMessengerServer(grpcServer, &server{})
	lis, err := net.Listen("tcp", ":7007")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}

	log.Println("TLS gRPC server listening on :7007")

	if err:= grpcServer.Serve(lis); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
