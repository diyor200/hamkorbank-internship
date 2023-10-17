package main

import (
	"github.com/diyor200/gin-middleware-blogpost/internal/app"
)

//	@title			Blog website
//	@version		1.0
//	@description	This is a sample server that manages tasks.

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	app.Run()
}
