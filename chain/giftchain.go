package chain

type Gift struct {
	nickname      string
	senderAddress string
	content       string
	startTime     int
	endTime       int
}

type GiftBox struct {
	gifts        []Gift
	timestamp    int
	nonce        string
	previousHash string
}

type GiftChain struct {
	chain     []GiftBox
	neighbors []string
	address   string
}

func CreateGift(nickname, senderAddress, content string, startTime, endTime int) *Gift {
	return &Gift{
		nickname:      nickname,
		senderAddress: senderAddress,
		content:       content,
		startTime:     startTime,
		endTime:       endTime,
	}
}

func CreateGiftBox(gifts []Gift, timestamp int, nonce, previousHash string) *GiftBox {
	return &GiftBox{
		gifts:        gifts,
		timestamp:    timestamp,
		nonce:        nonce,
		previousHash: previousHash,
	}
}

func (g *GiftBox) AddGiftToBox(gift Gift) {
	g.gifts = append(g.gifts, gift)
}

func (g *GiftChain) AddGiftBoxToChain(giftbox GiftBox) {
	g.chain = append(g.chain, giftbox)
}
