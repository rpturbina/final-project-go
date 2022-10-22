package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	engine "github.com/rpturbina/final-project-go/config/gin"
	"github.com/rpturbina/final-project-go/config/postgres"
	authrepo "github.com/rpturbina/final-project-go/pkg/repository/auth"
	commentrepo "github.com/rpturbina/final-project-go/pkg/repository/comment"
	photorepo "github.com/rpturbina/final-project-go/pkg/repository/photo"
	socialmediarepo "github.com/rpturbina/final-project-go/pkg/repository/socialmedia"
	userrepo "github.com/rpturbina/final-project-go/pkg/repository/user"
	authhandler "github.com/rpturbina/final-project-go/pkg/server/http/handler/auth"
	commenthandler "github.com/rpturbina/final-project-go/pkg/server/http/handler/comment"
	photohandler "github.com/rpturbina/final-project-go/pkg/server/http/handler/photo"
	socialmediahandler "github.com/rpturbina/final-project-go/pkg/server/http/handler/socialmedia"
	userhandler "github.com/rpturbina/final-project-go/pkg/server/http/handler/user"
	"github.com/rpturbina/final-project-go/pkg/server/http/middleware"
	router "github.com/rpturbina/final-project-go/pkg/server/http/router/v1"
	authusecase "github.com/rpturbina/final-project-go/pkg/usecase/auth"
	commentusecase "github.com/rpturbina/final-project-go/pkg/usecase/comment"
	photousecase "github.com/rpturbina/final-project-go/pkg/usecase/photo"
	socialmediausecase "github.com/rpturbina/final-project-go/pkg/usecase/socialmedia"
	userusecase "github.com/rpturbina/final-project-go/pkg/usecase/user"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	postgresHost := os.Getenv("MY_GRAM_POSTGRES_HOST")
	postgresPort := os.Getenv("MY_GRAM_POSTGRES_PORT")
	postgresDatabase := os.Getenv("MY_GRAM_POSTGRES_DATABASE")
	postgresUsername := os.Getenv("MY_GRAM_POSTGRES_USERNAME")
	postgresPassword := os.Getenv("MY_GRAM_POSTGRES_PASSWORD")

	postgresCln := postgres.NewPostgresConnection(postgres.Config{
		Host:         postgresHost,
		Port:         postgresPort,
		User:         postgresUsername,
		Password:     postgresPassword,
		DatabaseName: postgresDatabase,
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
	authUsecase := authusecase.NewAuthUsecase(authRepo, userUsecase)
	authHandler := authhandler.NewAuthHandler(authUsecase)

	photoRepo := photorepo.NewPhotoRepo(postgresCln)
	photoUsecase := photousecase.NewPhotoUsecase(photoRepo, userUsecase)
	photoHandler := photohandler.NewPhotoHandler(photoUsecase)

	commentRepo := commentrepo.NewCommentRepo(postgresCln)
	commentUsecase := commentusecase.NewCommentUsecase(commentRepo, photoUsecase)
	commentHandler := commenthandler.NewCommentHandler(commentUsecase)

	socialMediaRepo := socialmediarepo.NewSocialMediaRepo(postgresCln)
	socialMediaUsecase := socialmediausecase.NewSocialMediaUsecase(socialMediaRepo)
	socialMediaHandler := socialmediahandler.NewSocialMediaHandler(socialMediaUsecase)

	authMiddleware := middleware.NewAuthMiddleware(userUsecase)

	router.NewUserRouter(ginEngine, userHandler, authMiddleware).Routers()
	router.NewAuthRouter(ginEngine, authHandler, authMiddleware).Routers()
	router.NewPhotoRouter(ginEngine, photoHandler, authMiddleware).Routers()
	router.NewCommentRouter(ginEngine, commentHandler, authMiddleware).Routers()
	router.NewSocialMediaRouter(ginEngine, socialMediaHandler, authMiddleware).Routers()

	ginEngine.Serve()
}
