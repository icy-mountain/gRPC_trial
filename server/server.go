package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/icy-mountain/gRPC_trial/arithQuestioner"
	"google.golang.org/grpc"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
)

type arithQuestionerServer struct {
	pb.UnimplementedArithQuestionerServer
}

func (s *arithQuestionerServer) QuestionChat(stream pb.ArithQuestioner_QuestionChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		sendMessage := in.String() + "from server!"
		qm := &pb.QuestionMessage{Message: sendMessage}
		if err := stream.Send(qm); err != nil {
			return err
		}
	}
}

func newServer() *arithQuestionerServer {
	s := &arithQuestionerServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterArithQuestionerServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
