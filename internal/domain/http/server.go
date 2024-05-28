package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mxrcury/dragonsage/internal/config"
)

type Server struct {
	Router *gin.Engine
	port   string
}

func NewServer(cfg *config.ServerConfig) *Server {
	router := gin.Default()

	router.Use(gin.Logger())

	return &Server{Router: router, port: cfg.Port}
}

func (s *Server) Run() error {
	return s.Router.Run(":" + s.port)
}
