package app

import (
	"github.com/diyor200/learn-sqlx/config"
	"github.com/diyor200/learn-sqlx/crud"
	"github.com/diyor200/learn-sqlx/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)

func Run() {
	cfg := config.NewConfig()
	db, err := sqlx.Connect("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalln(err)
	}
	cr := crud.NewCrud(db)
	c := controller.NewController(cr)

	router := gin.Default()
	router.GET("/users", c.GetUsers)
	router.POST("/insert", c.InsertUser)
	router.POST("/delete", c.Delete)
	router.POST("/update", c.UpdateUser)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
