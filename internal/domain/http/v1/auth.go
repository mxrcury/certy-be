package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mxrcury/certy/internal/service"
)

type AuthHandler struct {
	service       service.Auth
	tokensService service.Tokens
	path          string
}

type AuthHandlerDeps struct {
	service       service.Auth
	tokensService service.Tokens
}

func NewAuthHandler(path string, deps *AuthHandlerDeps) Auth {
	return &AuthHandler{
		path:          path,
		service:       deps.service,
		tokensService: deps.tokensService,
	}
}

func (h *AuthHandler) group(group *gin.RouterGroup) {
	usersGroup := group.Group(h.path)
	{
		usersGroup.POST("/sign-up", h.signUp)
		usersGroup.POST("/sign-in", h.signIn)
	}
}

func (h *AuthHandler) signUp(c *gin.Context) {
	var body service.SignUpInput

	if err := c.ShouldBindJSON(&body); err != nil {
		sendResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.service.SignUp(&body); err != nil {
		sendResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	c.Status(201)
}

func (h *AuthHandler) signIn(c *gin.Context) {
	var body service.SignInInput

	if err := c.ShouldBindJSON(&body); err != nil {
		sendResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	resp, err := h.service.SignIn(&body)

	if err != nil {
		sendResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	accessToken, err := h.tokensService.GenerateJWT(resp.ID.String())
	if err != nil {
		sendResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	c.JSON(201, &service.JWTTokens{AccessToken: accessToken})
}
