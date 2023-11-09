package app

import (
	"absapi/config"
	"absapi/internal/handler"
	"absapi/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func Run() {
	cfg := config.NewConfig()
	db, err := sqlx.Connect("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalln(err)
	}

	repo := repository.NewRepository(db)
	h := handler.NewHandler(repo)

	router := gin.Default()
	router.GET("/:persId", h.GetUserMap)

	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
