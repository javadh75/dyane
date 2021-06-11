package device

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/javadh75/dyane/device/modules/network"
	pb "github.com/javadh75/dyane/protoc"
	"google.golang.org/grpc"
)

const (
	port = ":5000"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func RunServer() {
	fmt.Println("Running Device Agent")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func Run(args []string) {
}

func GetLinkCmd(dev string) {
	nllink, err := network.GetLink(dev)
	if err != nil {
		log.Println("Error:", err)
	}
	dlink := network.DyaneLink{}
	dlink.Set(nllink)
	linkJson, err := json.MarshalIndent(dlink, "", "  ")
	if err != nil {
		log.Println("Error:", err)
	}
	fmt.Print(string(linkJson))
}

func GetAllLinksCmd() {
	nllinks, err := network.GetAllLinks()
	if err != nil {
		log.Println("Error:", err)
	}
	dlinks := []network.DyaneLink{}
	for _, nllink := range nllinks {
		dlink := network.DyaneLink{}
		dlink.Set(nllink)
		dlinks = append(dlinks, dlink)
	}
	linksJson, err := json.MarshalIndent(dlinks, "", "  ")
	if err != nil {
		log.Println("Error:", err)
	}
	fmt.Print(string(linksJson))
}

func InitCurrentConfigsCMD() {
	network.InitCurrentConfigs()
}
