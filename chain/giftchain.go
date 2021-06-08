package chain

import (
	"context"
	"github.com/beautifultools/gift-chain/grpc/neighbor"
	"google.golang.org/grpc"
	"log"
	"time"
)

type GiftChain struct {
	giftBoxes []GiftBox
	gifts     []Gift
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
		giftBoxes: []GiftBox{},
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

func (g *GiftChain) AddGift(gift *Gift) {
	g.gifts = append(g.gifts, *gift)
}

func (g *GiftChain) MakeGiftBox() {
	g.giftBoxes = append(g.giftBoxes, GiftBox{
		gifts:g.gifts,

	})
	g.gifts = []Gift{}
}

func (g *GiftChain) AddGiftBox(giftbox *GiftBox) {
	g.giftBoxes = append(g.giftBoxes, *giftbox)
}
