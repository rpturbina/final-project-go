package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/rpturbina/final-project-go/config/gin"
	"github.com/rpturbina/final-project-go/pkg/domain/photo"
	"github.com/rpturbina/final-project-go/pkg/server/http/middleware"
	"github.com/rpturbina/final-project-go/pkg/server/http/router"
)

type PhotoRouterImpl struct {
	ginEngine      engine.HttpServer
	routerGroup    *gin.RouterGroup
	photoHandler   photo.PhotoHandler
	authMiddleware middleware.AuthMiddleware
}

func (p *PhotoRouterImpl) get() {
	p.routerGroup.GET("", p.authMiddleware.CheckJWTAuth, p.photoHandler.GetPhotosHdl)
}

func (p *PhotoRouterImpl) post() {
	p.routerGroup.POST("", p.authMiddleware.CheckJWTAuth, p.photoHandler.CreatePhotoHdl)
}

func (p *PhotoRouterImpl) put() {
	p.routerGroup.PUT("/:photoId", p.authMiddleware.CheckJWTAuth, p.photoHandler.UpdatePhotoHdl)
}

func (p *PhotoRouterImpl) delete() {
	p.routerGroup.DELETE("/:photoId", p.authMiddleware.CheckJWTAuth, p.photoHandler.DeletePhotoHdl)
}

func (p *PhotoRouterImpl) Routers() {
	p.get()
	p.post()
	p.put()
	p.delete()
}

func NewPhotoRouter(ginEngine engine.HttpServer, photoHandler photo.PhotoHandler, authMiddleware middleware.AuthMiddleware) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/mygram/v1/photos")
	return &PhotoRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, photoHandler: photoHandler, authMiddleware: authMiddleware}
}
