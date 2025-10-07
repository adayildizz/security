package main

import (
	"context"
	"crypto/x509"
	"flag"
	"log"
	"os"
	"time"

	pb "assignment1/code/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	addr := flag.String("addr", "localhost:7007", "server host:port")
	msg := flag.String("msg", "hello from client", "message to send")
	caPath := flag.String("ca", "../server/server.crt", "path to server certificate (PEM) to trust")
	serverName := flag.String("servername", "localhost", "expected TLS server name (CN/SAN)")
	flag.Parse()

	// 1) Load server cert (self-signed) into a Root CA pool we trust
	caBytes, err := os.ReadFile(*caPath)
	if err != nil {
		log.Fatalf("read ca cert: %v", err)
	}
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(caBytes) {
		log.Fatalf("failed to append CA cert: %s", *caPath)
	}

	
	creds := credentials.NewClientTLSFromCert(cp, *serverName)

	conn, err := grpc.Dial(
		*addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(), // optional: wait until the connection is up
	)
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer conn.Close()

	
	client := pb.NewMessengerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.Send(ctx, &pb.SendRequest{Message: *msg})
	if err != nil {
		log.Fatalf("rpc error: %v", err)
	}
	log.Printf("Server replied: %s", res.GetStatus())
}
