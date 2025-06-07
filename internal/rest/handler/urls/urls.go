package urls

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"url-shortener/internal/domain/entites"
)

func (h *handler) Shorten(c *gin.Context) {
	var req ShortenUrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("url: %v\nidentifier: %v\n", req.Url, req.Identifier)

	newUrl, err := h.urlService.Shorten(c.Request.Context(), &entites.InputUrl{RealUrl: req.Url, Identifier: req.Identifier})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UrlResponse{ShortUrl: newUrl})
}

func (h *handler) Redirect(c *gin.Context) {
	urlObj, err := h.urlService.Redirect(c.Request.Context(), c.Param("token"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// log.Printf("Real-url of token %v is %v\n", urlObj.Token, urlObj.RealUrl)
	c.Redirect(http.StatusPermanentRedirect, urlObj.RealUrl)
}
