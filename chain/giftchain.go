package chain

type Gift struct {
	Nickname      string
	SenderAddress string
	Content       string
	StartTime     int64
	EndTime       int64
}

type GiftBox struct {
	gifts        []Gift
	timestamp    int64
	nonce        string
	previousHash string
}

type GiftChain struct {
	chain     []GiftBox
	neighbors []string
	address   string
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
