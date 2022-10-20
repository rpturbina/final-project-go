package postgres

import (
	"fmt"
	"log"

	"github.com/rpturbina/final-project-go/pkg/domain/comment"
	"github.com/rpturbina/final-project-go/pkg/domain/photo"
	"github.com/rpturbina/final-project-go/pkg/domain/socialmedia"
	"github.com/rpturbina/final-project-go/pkg/domain/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	DatabaseName string `json:"database_name"`
	User         string `json:"user"`
	Password     string `json:"password"`
}

type PostgresClient interface {
	GetClient() *gorm.DB
}

type PostgresClientImpl struct {
	cln    *gorm.DB
	config Config
}

func (p *PostgresClientImpl) GetClient() *gorm.DB {
	return p.cln
}

func NewPostgresConnection(config Config) PostgresClient {
	connectionConfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", config.Host, config.Port, config.User, config.Password, config.DatabaseName)

	db, err := gorm.Open(postgres.Open(connectionConfig), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("error connection to database:", err)
	}

	fmt.Println("database connection is successfully connected")
	db.AutoMigrate(&user.User{}, &photo.Photo{}, &comment.Comment{}, &socialmedia.SocialMedia{})

	return &PostgresClientImpl{cln: db, config: config}
}
