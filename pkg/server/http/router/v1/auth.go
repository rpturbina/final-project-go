package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/rpturbina/final-project-go/config/gin"
	"github.com/rpturbina/final-project-go/pkg/domain/auth"
	"github.com/rpturbina/final-project-go/pkg/server/http/middleware"
	"github.com/rpturbina/final-project-go/pkg/server/http/router"
)

type AuthRouterImpl struct {
	ginEngine      engine.HttpServer
	routerGroup    *gin.RouterGroup
	authHandler    auth.AuthHandler
	authMiddleware middleware.AuthMiddleware
}

func (a *AuthRouterImpl) post() {
	a.routerGroup.POST("/login", a.authHandler.LoginUserHdl)
	a.routerGroup.POST("/refresh", a.authMiddleware.CheckJWTAuth, a.authHandler.GetRefreshTokenHdl)
}

func (a *AuthRouterImpl) Routers() {
	a.post()
}

func NewAuthRouter(ginEngine engine.HttpServer, authHandler auth.AuthHandler, auhtMiddleware middleware.AuthMiddleware) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/mygram/v1/auth")
	return &AuthRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, authHandler: authHandler, authMiddleware: auhtMiddleware}
}
