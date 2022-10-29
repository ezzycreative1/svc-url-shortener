package usecase

import (
	"context"
	"os"
	"time"

	"github.com/ezzycreative1/svc-url-shortener/internal/core/domain"
	"github.com/ezzycreative1/svc-url-shortener/internal/core/ports"
	"github.com/teris-io/shortid"
)

type shortUsecase struct {
	Repo ports.IShortenerRepository
}

func NewUrlShortUsecase(
	repo ports.IShortenerRepository,
) shortUsecase {
	return shortUsecase{
		Repo: repo,
	}
}

func (us *shortUsecase) Store(ctx context.Context, input *domain.UrlShortSpec) error {
	codeKey := shortid.MustGenerate()
	req := &domain.UrlShort{
		Code:      codeKey,
		LongUrl:   input.Url,
		ShortUrl:  os.Getenv("BASE_REDIRECT") + "/" + codeKey,
		CreatedAt: time.Now().UTC().Unix(),
	}

	return us.Repo.Store(ctx, req)
}
