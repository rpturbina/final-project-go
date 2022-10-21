package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	engine "github.com/rpturbina/final-project-go/config/gin"
	"github.com/rpturbina/final-project-go/config/postgres"
	authrepo "github.com/rpturbina/final-project-go/pkg/repository/auth"
	photorepo "github.com/rpturbina/final-project-go/pkg/repository/photo"
	userrepo "github.com/rpturbina/final-project-go/pkg/repository/user"
	authhandler "github.com/rpturbina/final-project-go/pkg/server/http/handler/auth"
	photohandler "github.com/rpturbina/final-project-go/pkg/server/http/handler/photo"
	userhandler "github.com/rpturbina/final-project-go/pkg/server/http/handler/user"
	"github.com/rpturbina/final-project-go/pkg/server/http/middleware"
	router "github.com/rpturbina/final-project-go/pkg/server/http/router/v1"
	authusecase "github.com/rpturbina/final-project-go/pkg/usecase/auth"
	photousecase "github.com/rpturbina/final-project-go/pkg/usecase/photo"
	userusecase "github.com/rpturbina/final-project-go/pkg/usecase/user"
)

func main() {
	postgresCln := postgres.NewPostgresConnection(postgres.Config{
		Host:         "localhost",
		Port:         "5432",
		User:         "postgres",
		Password:     "mysecretpassword",
		DatabaseName: "mygram",
	})

	ginEngine := engine.NewGinHttp(engine.Config{
		Port: ":8080",
	})

	ginEngine.GetGin().Use(
		gin.Recovery(),
		gin.Logger(),
	)

	startTime := time.Now()
	ginEngine.GetGin().GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":    "server up and running",
			"start_time": startTime,
		})
	})

	userRepo := userrepo.NewUserRepo(postgresCln)
	userUsecase := userusecase.NewUserUsecase(userRepo)
	userHandler := userhandler.NewUserHandler(userUsecase)

	authRepo := authrepo.NewAuthRepo(postgresCln)
	authUsecase := authusecase.NewAuthUsecase(authRepo, userRepo)
	authHandler := authhandler.NewAuthHandler(authUsecase)

	photoRepo := photorepo.NewPhotoRepo(postgresCln)
	photoUsecase := photousecase.NewPhotoUsecase(photoRepo)
	photoHandler := photohandler.NewPhotoHandler(photoUsecase)

	authMiddleware := middleware.NewAuthMiddleware(userUsecase)

	router.NewUserRouter(ginEngine, userHandler, authMiddleware).Routers()
	router.NewAuthRouter(ginEngine, authHandler, authMiddleware).Routers()
	router.NewPhotoRouter(ginEngine, photoHandler, authMiddleware).Routers()

	ginEngine.Serve()
}
