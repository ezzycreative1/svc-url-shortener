package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ezzycreative1/svc-url-shortener/app/v1/handler/request"
	urlPort "github.com/ezzycreative1/svc-url-shortener/business/shortener/ports"
	"github.com/ezzycreative1/svc-url-shortener/config"
	"github.com/ezzycreative1/svc-url-shortener/pkg/mid"
	"github.com/ezzycreative1/svc-url-shortener/pkg/mlog"
	"github.com/ezzycreative1/svc-url-shortener/pkg/mvalidator"
)

type UrlShortHandler struct {
	UseCase   *urlPort.IUrlShortenerUsecase
	Validator mvalidator.Validator
	Logger    mlog.Logger
	Cfg       config.Group
}

func NewCogsHandler(
	usecase *urlPort.IUrlShortenerUsecase,
	validator mvalidator.Validator,
	logger mlog.Logger,
	config config.Group,
) UrlShortHandler {
	return UrlShortHandler{
		UseCase:   usecase,
		Validator: validator,
		Logger:    logger,
		Cfg:       config,
	}
}

func (us *UrlShortHandler) Create(w http.ResponseWriter, r *http.Request) error {

	ctx := r.Context()
	requestID := mid.GetID(ctx)
	usecaseContext := mid.SetIDx(ctx, requestID)

	if r.ContentLength == 0 {
		return errors.New("empty body")
	}

	//var plan *model.Plan
	req := request.UrlShortReq{
		Url: r.PostFormValue("url"),
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		us.Logger.ErrorT(requestID, "create url shortener payload", err, mlog.Any("payload", req))
		return err
	}

	data, err := us.UseCase.CreateShortUrl(usecaseContext, req)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(data)
}
