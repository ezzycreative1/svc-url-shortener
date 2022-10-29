package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/ezzycreative1/svc-url-shortener/internal/core/domain"
	"github.com/go-redis/redis/v8"
)

type shortRepo struct {
	Client *redis.Client
}

func NewShortenerRepo(client *redis.Client) shortRepo {
	return shortRepo{
		Client: client,
	}
}

func (sr *shortRepo) generateKey(code string) string {
	return fmt.Sprintf("redirect:%s", code)
}

func (sr *shortRepo) Store(ctx context.Context, input *domain.UrlShort) error {
	key := sr.generateKey(input.Code)

	data := map[string]interface{}{
		"code":       input.Code,
		"url":        input.LongUrl,
		"short_url":  input.ShortUrl,
		"created_at": time.Now(),
	}
	_, err := sr.Client.HMSet(ctx, key, data).Result()
	if err != nil {
		return err
	}
	return nil
}
