package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"os"

	pb "github.com/icy-mountain/gRPC_trial/arithQuestioner"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func runQuestionChat(client pb.ArithQuestionerClient) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := client.QuestionChat(ctx)
	if err != nil {
		log.Fatalf("client.QuestionChat failed: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("runQuestionChat failed: %v", err)
			}
			log.Printf("Got message %s\n", in.Message)
		}
	}()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		qMessage := &pb.QuestionMessage{Message: s.Text()}
		if err := stream.Send(qMessage); err != nil {
			log.Fatalf("client.QuestionMessage: stream.Send(%v) failed: %v", qMessage, err)
		}
	}
	stream.CloseSend()
	<-waitc
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewArithQuestionerClient(conn)

	runQuestionChat(client)
}
