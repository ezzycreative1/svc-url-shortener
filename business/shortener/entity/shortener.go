package entity

type UrlShort struct {
	Url      string `json:"url"`
	ShortUrl string `json:"short_url"`
}

type UrlShortSpec struct {
	Url string `validate:"required"`
}
