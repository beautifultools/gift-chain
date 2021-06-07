package chain

import (
	"context"
	"github.com/beautifultools/gift-chain/grpc/neighbor"
	"google.golang.org/grpc"
	"log"
	"time"
)

type GiftChain struct {
	chain     []GiftBox
	neighbors []string
	address   string
}

type GiftBox struct {
	gifts        []Gift
	timestamp    int64
	nonce        string
	previousHash string
}

type Gift struct {
	Nickname      string
	SenderAddress string
	Content       string
	StartTime     int64
	EndTime       int64
}

func CreateGiftChain(serverAddress string) *GiftChain{
	getNeighborAddressesByGivenServer(serverAddress)

	return &GiftChain{
		chain:     []GiftBox{},
		neighbors: []string{},
	}
}

func getNeighborAddressesByGivenServer(serverAddress string) []string {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := neighbor.NewNeighborGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	neighbors, err := c.GetNeighbors(ctx, &neighbor.NeighborOptions{})

	if err != nil {
		log.Fatalf("fail to get neighbors: %v", err)
	}
	
	return neighbors.Addresses
}


func CreateGiftBox(gifts []Gift, timestamp int64, nonce, previousHash string) *GiftBox {
	return &GiftBox{
		gifts:        gifts,
		timestamp:    timestamp,
		nonce:        nonce,
		previousHash: previousHash,
	}
}

func (g *GiftBox) AddGiftToBox(gift *Gift) {
	g.gifts = append(g.gifts, *gift)
}

func (g *GiftChain) AddGiftBoxToChain(giftbox GiftBox) {
	g.chain = append(g.chain, giftbox)
}
