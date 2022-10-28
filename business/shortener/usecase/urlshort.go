package services

import (
	"context"

	"github.com/ezzycreative1/svc-url-shortener/business/shortener/entity"
	"github.com/ezzycreative1/svc-url-shortener/business/shortener/ports"
	"github.com/ezzycreative1/svc-url-shortener/config"
	"github.com/ezzycreative1/svc-url-shortener/pkg/mlog"
	"github.com/ezzycreative1/svc-url-shortener/pkg/mvalidator"
)

type UrlShortUsecase struct {
	Repo      *ports.Repository
	Validator mvalidator.Validator
	Logger    mlog.Logger
	Cfg       config.Group
}

func NewUrlShortUsecase(
	repo *ports.Repository,
	validator mvalidator.Validator,
	logger mlog.Logger,
	config config.Group,
) UrlShortUsecase {
	return UrlShortUsecase{
		Repo:      repo,
		Validator: validator,
		Logger:    logger,
		Cfg:       config,
	}
}

func (uss *UrlShortUsecase) CreateShortUrl(ctx context.Context, entity.UrlShortSpec) (*entity.UrlShort, error) {
	
	return nil, nil
}
