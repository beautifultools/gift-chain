package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddGiftToBox(t *testing.T) {
	gifts := []Gift{}
	gifts = append(gifts, *CreateGift("john", "abcd", "create design of main screen", 000000, 111111))
	giftBox := CreateGiftBox(gifts, 100000, "123", "bbcc")

	assert.Equal(t, len(giftBox.gifts), 1)
}
