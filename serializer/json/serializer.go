package serializer

import (
	"encoding/json"

	"github.com/ezzycreative1/svc-url-shortener/internal/core/domain"
	"github.com/pkg/errors"
)

// Redirect is an implementation of shortener.Encoder
type Redirect struct{}

// Decode receives json message in bytes and convert to pointer of shortener.Redirect
func (r *Redirect) Decode(input []byte) (*domain.UrlShort, error) {
	redirect := &domain.UrlShort{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return redirect, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

// Encode receives a pointer of shortener.Redirect and returns json message in bytes
func (r *Redirect) Encode(input *domain.UrlShort) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return rawMsg, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
