package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener/internal/rest/handler/urls"
)

type Server struct {
	mux        *gin.Engine
	urlHandler urls.UrlHandler
}

func NewServer(mux *gin.Engine, urlHandler urls.UrlHandler) *Server {
	mux.Use(gin.Recovery())
	mux.Use(gin.Logger())
	return &Server{mux: mux, urlHandler: urlHandler}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) Init() {
	const basUrl = "/api/v1"

	baseGroup := s.mux.Group(basUrl)
	baseGroup.POST("/shorten", s.urlHandler.Shorten)
	s.mux.GET("/:token", s.urlHandler.Redirect)
}
