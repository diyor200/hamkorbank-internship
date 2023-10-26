package app

import (
	"github.com/diyor200/gin-middleware-blogpost/config"
	"github.com/diyor200/gin-middleware-blogpost/docs"
	"github.com/diyor200/gin-middleware-blogpost/internal/controller"
	"github.com/diyor200/gin-middleware-blogpost/internal/middleware"
	"github.com/diyor200/gin-middleware-blogpost/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Version = "1.0"
	cfg := config.NewConfig()

	db, err := sqlx.Connect("postgres", cfg.DBUrl)
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepo(db)
	c := controller.NewController(repo)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/users", c.GetUsers)
	router.GET("/posts", c.GetPosts)
	router.GET("/posts/:post_id", c.GetPost)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", c.SignUp)
		auth.POST("/sign-in", c.SignIn)
	}

	actions := router.Group("/action")
	actions.Use(middleware.CheckUser())
	actions.POST("/create/post", c.CreatePost)
	d := actions.Group("/delete")
	{
		d.POST("/user", c.DeleteUser)
		d.POST("/post/:post_id", c.DeletePost)
	}
	actions.POST("/edit/post", c.EditPost)

	if err := router.Run(":8083"); err != nil {
		panic(err)
	}
}
