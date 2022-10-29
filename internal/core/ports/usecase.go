package ports

import (
	"context"

	"github.com/ezzycreative1/svc-url-shortener/internal/core/domain"
)

type IShortenerUsecase interface {
	Store(context.Context, domain.UrlShortSpec) error
}

type RedirectSerializer interface {
	Decode(input []byte) (*Redirect, error)
	Encode(input *Redirect) ([]byte, error)
}
