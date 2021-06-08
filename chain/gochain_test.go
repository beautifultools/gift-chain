package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const serverAddress = "127.0.0.1:50051"

func TestCreateGiftChain(t *testing.T) {
	giftChain := CreateGiftChain(serverAddress)

	assert.GreaterOrEqual(t, len(giftChain.neighbors), 1)
}

func TestGiftChain_AddGift(t *testing.T) {
	giftChain := CreateGiftChain(serverAddress)

	giftChain.AddGift(&Gift{})
	assert.Equal(t, len(giftChain.gifts), 1)

	giftChain.AddGift(&Gift{})
	assert.Equal(t, len(giftChain.gifts), 2)
}

func TestGiftChain_AddGiftBox(t *testing.T) {
	giftChain := CreateGiftChain(serverAddress)

	giftChain.AddGiftBox(&GiftBox{})
	assert.Equal(t, len(giftChain.giftBoxes), 1)
}

func TestGiftChain_MakeGiftBox(t *testing.T) {
	giftChain := CreateGiftChain(serverAddress)

	giftChain.MakeGiftBox()
	assert.Equal(t, len(giftChain.giftBoxes), 1)

	giftChain.MakeGiftBox()
	assert.Equal(t, len(giftChain.giftBoxes), 1)
}
