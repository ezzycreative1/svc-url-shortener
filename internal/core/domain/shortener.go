package domain

type UrlShort struct {
	Code      string `json:"code"`
	LongUrl   string `json:"long_url" validate:"required"`
	ShortUrl  string `json:"short_url"`
	CreatedAt int64  `json:"created_at"`
}

type UrlShortSpec struct {
	Url string `validate:"required"`
}

type UrlResponse struct {
	ShortUrl string `json:"short_url"`
}

// Redirect is an implementation of shortener.Encoder
type Redirect struct {
	Code string `json:"code"`
	Url  string `json:"url"`
}
