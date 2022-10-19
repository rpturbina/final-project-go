package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	engine "github.com/rpturbina/final-project-go/config/gin"
	"github.com/rpturbina/final-project-go/config/postgres"
	userrepo "github.com/rpturbina/final-project-go/pkg/repository/user"
	userhandler "github.com/rpturbina/final-project-go/pkg/server/http/handler/user"
	router "github.com/rpturbina/final-project-go/pkg/server/http/router/v1"
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

	router.NewUserRouter(ginEngine, userHandler).Routers()

	ginEngine.Serve()
}
