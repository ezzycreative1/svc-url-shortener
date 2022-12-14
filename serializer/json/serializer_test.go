package serializer

import (
	"encoding/json"
	"testing"

	"github.com/ezzycreative1/svc-url-shortener/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestRedirect(t *testing.T) {
	serializer := Redirect{}

	t.Run("Decode with invalid input", func(t *testing.T) {
		invalidRawMsg := []byte("msg=strogonoficamente-sensivel-message")
		_, err := serializer.Decode(invalidRawMsg)
		assert.Equal(t, "serializer.Redirect.Decode: invalid character 'm' looking for beginning of value", err.Error())
	})

	t.Run("Decode with valid input", func(t *testing.T) {
		redirect := &domain.UrlShort{
			Code:    "github-ezzycreative1",
			LongUrl: "https://github.com/ezzycreative1",
		}
		rawMsg, err := json.Marshal(redirect)
		assert.Nil(t, err)
		redirectResult, err := serializer.Decode(rawMsg)
		assert.Nil(t, err)
		assert.Equal(t, redirect, redirectResult)
	})

	t.Run("Encode with valid input", func(t *testing.T) {
		redirect := &domain.UrlShort{
			Code:      "github-allisson",
			LongUrl:   "https://github.com/allisson",
			CreatedAt: 949407194000,
		}
		rawMsg, err := json.Marshal(redirect)
		assert.Nil(t, err)
		rawMsgResult, err := serializer.Encode(redirect)
		assert.Nil(t, err)
		assert.Equal(t, rawMsgResult, rawMsg)
	})
}
