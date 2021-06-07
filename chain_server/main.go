package main

import (
	"context"
	pb "github.com/beautifultools/gift-chain/grpc/gift"
	chain "github.com/beautifultools/gift-chain/chain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGiftChainServer
}

func (s *server) AddGift(ctx context.Context, gift *pb.Gift) (*pb.Reply, error) {
	giftbox := chain.CreateGiftBox([]chain.Gift{}, time.Now().Unix(), "", "")

	giftbox.AddGiftToBox(&chain.Gift{
		Nickname:gift.Nickname,
		SenderAddress:gift.SenderAddress,
		Content:gift.Content,
		StartTime:gift.StartTime,
		EndTime:gift.EndTime,
	})

	return nil, status.Errorf(codes.Unimplemented, "method AddGift not implemented")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("failed to listen : %v", err)
	}
	s:= grpc.NewServer()
	pb.RegisterGiftChainServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to add a gift to chain")
	}
}
