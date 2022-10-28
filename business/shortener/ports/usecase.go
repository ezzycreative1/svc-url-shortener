package ports

import (
	"context"
	"github.com/ezzycreative1/svc-url-shortener/business/shortener/entity"
)

type IUrlShortenerUsecase interface {
	CreateShortUrl(ctx *context.Context, entity.UrlShortSpec) (*entity.UrlShort, error)
}
