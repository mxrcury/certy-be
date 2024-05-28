package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mxrcury/dragonsage/internal/service"
)

type Handlers struct {
	authHandler Auth

	group *gin.RouterGroup
}

type Deps struct {
	Router *gin.Engine

	Services *service.Services
}

type (
	baseHandler interface {
		group(*gin.RouterGroup)
	}

	Auth interface {
		baseHandler

		signUp(*gin.Context)
		signIn(*gin.Context)
	}
)

func InitHandlers(deps *Deps) *Handlers {
	authHandler := NewAuthHandler("/auth", &AuthHandlerDeps{
		service:       deps.Services.AuthService,
		tokensService: deps.Services.TokensService,
	})

	group := deps.Router.Group("/v1")

	authHandler.group(group)

	return &Handlers{authHandler: authHandler, group: group}
}
