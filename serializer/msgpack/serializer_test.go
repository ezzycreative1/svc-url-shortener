package serializer

import (
	"testing"

	"github.com/ezzycreative1/svc-url-shortener/internal/core/domain"
	"github.com/vmihailenco/msgpack"

	"github.com/stretchr/testify/assert"
)

func TestRedirect(t *testing.T) {
	serializer := Redirect{}

	t.Run("Decode with invalid input", func(t *testing.T) {
		invalidRawMsg := []byte("msg=strogonoficamente-sensivel-message")
		_, err := serializer.Decode(invalidRawMsg)
		assert.Equal(t, "serializer.Redirect.Decode: msgpack: invalid code=6d decoding map length", err.Error())
	})

	t.Run("Decode with valid input", func(t *testing.T) {
		redirect := &domain.UrlShort{
			Code:    "github-ezzycreative1",
			LongUrl: "https://github.com/ezzycreative1",
		}
		rawMsg, err := msgpack.Marshal(redirect)
		assert.Nil(t, err)
		redirectResult, err := serializer.Decode(rawMsg)
		assert.Nil(t, err)
		assert.Equal(t, redirect, redirectResult)
	})

	t.Run("Encode with valid input", func(t *testing.T) {
		redirect := &domain.UrlShort{
			Code:      "github-ezzycreative1",
			LongUrl:   "https://github.com/ezzycreative1",
			CreatedAt: 949407194000,
		}
		rawMsg, err := msgpack.Marshal(redirect)
		assert.Nil(t, err)
		rawMsgResult, err := serializer.Encode(redirect)
		assert.Nil(t, err)
		assert.Equal(t, rawMsgResult, rawMsg)
	})
}
