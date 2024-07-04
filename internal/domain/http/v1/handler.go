package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mxrcury/certy/internal/config"
	"github.com/mxrcury/certy/internal/service"
)

type Handlers struct {
	authHandler Auth

	group *gin.RouterGroup
}

type Deps struct {
	Router *gin.Engine

	Services *service.Services

	Config *config.Config
}

type (
	baseHandler interface {
		group(*gin.RouterGroup)
	}

	Auth interface {
		baseHandler

		signUp(*gin.Context)
		signIn(*gin.Context)
		sendVerificationCode(*gin.Context)
		verifyCode(*gin.Context)
		getProfile(*gin.Context)
	}
)

func InitHandlers(deps *Deps) *Handlers {
	authHandler := NewAuthHandler("/auth", &AuthHandlerDeps{
		service:       deps.Services.AuthService,
		tokensService: deps.Services.TokensService,
		domainName:    deps.Config.ServerConfig.Domain,
	})

	group := deps.Router.Group("/v1")

	authHandler.group(group)

	return &Handlers{authHandler: authHandler, group: group}
}
