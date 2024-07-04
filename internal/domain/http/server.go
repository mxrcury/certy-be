package http

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mxrcury/certy/internal/config"
)

type Server struct {
	Router *gin.Engine
	port   string
}

var allowedMethods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPatch,
	http.MethodPut,
	http.MethodDelete,
	http.MethodOptions,
}

func NewServer(cfg *config.ServerConfig) *Server {
	router := gin.Default()

	router.Use(applyCors(cfg.ClientURL))
	router.Use(gin.Logger())

	return &Server{Router: router, port: cfg.Port}
}

func (s *Server) Run() error {
	return s.Router.Run(":" + s.port)
}

func applyCors(clientURL string) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{clientURL}
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = allowedMethods

	return cors.New(corsConfig)
}
