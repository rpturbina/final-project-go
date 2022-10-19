package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/rpturbina/final-project-go/config/gin"
	"github.com/rpturbina/final-project-go/pkg/domain/user"
	"github.com/rpturbina/final-project-go/pkg/server/http/router"
)

type UserRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	userHandler user.UserHandler
}

func (u *UserRouterImpl) post() {
	u.routerGroup.POST("/register", u.userHandler.RegisterUserHdl)
}

func (u *UserRouterImpl) get() {
	u.routerGroup.GET("/:user_id", u.userHandler.GetUserByIdHdl)
}

func (u *UserRouterImpl) Routers() {
	u.post()
	u.get()
}

func NewUserRouter(ginEngine engine.HttpServer, userHandler user.UserHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/mygram/v1/users")
	return &UserRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, userHandler: userHandler}
}
