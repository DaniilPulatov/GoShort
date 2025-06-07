package urls

import (
	"github.com/gin-gonic/gin"
	"url-shortener/internal/usecase/urls"
)

type UrlHandler interface {
	Shorten(c *gin.Context)
	Redirect(c *gin.Context)
}

type handler struct {
	urlService urls.UrlService
}

func NewUrlHandler(urlService urls.UrlService) UrlHandler {
	return &handler{urlService: urlService}
}

type ShortenUrlRequest struct {
	Url        string `json:"url" binding:"required"`
	Identifier string `json:"identifier"`
}

type UrlResponse struct {
	ShortUrl string `json:"short-url"`
}
