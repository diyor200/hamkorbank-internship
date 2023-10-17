package app

import (
	"github.com/diyor200/learn-gorm/config"
	"github.com/diyor200/learn-gorm/internal/controller"
	"github.com/diyor200/learn-gorm/internal/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Run() {
	cfg := config.NewConfig()
	db, err := gorm.Open(postgres.Open(cfg.DBUrl))
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatalf("error automigrating model: %v", err)
	}

	ctrl := controller.NewController(db)

	router := gin.Default()

	router.GET("/users", ctrl.Read)
	router.POST("/create", ctrl.Create)
	router.POST("/delete/:id", ctrl.Delete)
	router.POST("/update", ctrl.Update)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}
