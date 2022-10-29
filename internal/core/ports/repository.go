package ports

import (
	"context"

	"github.com/ezzycreative1/svc-url-shortener/internal/core/domain"
)

type IShortenerRepository interface {
	Find(code string) (*domain.UrlShort, error)
	Store(context.Context, *domain.UrlShort) error
}
