package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ezzycreative1/svc-url-shortener/config"
	"github.com/ezzycreative1/svc-url-shortener/internal/core/ports"
	"github.com/ezzycreative1/svc-url-shortener/pkg/mid"
	"github.com/ezzycreative1/svc-url-shortener/pkg/mlog"
	"github.com/ezzycreative1/svc-url-shortener/pkg/mvalidator"
	js "github.com/ezzycreative1/svc-url-shortener/serializer/json"
	ms "github.com/ezzycreative1/svc-url-shortener/serializer/msgpack"

	//shortener "github.com/ezzycreative1/svc-url-shortener/shortener"
	"github.com/pkg/errors"
)

func makeResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

type UrlShortHandler struct {
	UseCase   ports.IShortenerUsecase
	Validator mvalidator.Validator
	Logger    mlog.Logger
	Cfg       config.Group
}

func NewCogsHandler(
	usecase ports.IShortenerUsecase,
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

func (us *UrlShortHandler) serializer(contentType string) shortener.RedirectSerializer {
	if contentType == "application/x-msgpack" {
		return &ms.Redirect{}
	}
	return &js.Redirect{}
}

func (us *UrlShortHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := mid.GetID(ctx)
	userContext := mid.SetIDx(ctx, requestID)

	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	redirect, err := us.serializer(contentType).Decode(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = us.UseCase.Store(userContext, redirect)
	if err != nil {
		if errors.Cause(err) == shortener.ErrRedirectInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := us.serializer(contentType).Encode(redirect)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	makeResponse(w, contentType, responseBody, http.StatusCreated)
}
