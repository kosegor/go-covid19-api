package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kosegor/go-covid19-api/app/interface/controller"
	"github.com/kosegor/go-covid19-api/app/interface/persistence/memory"
	"github.com/kosegor/go-covid19-api/app/usecase"
)

func CreateServer() *gin.Engine {
	var router = gin.Default()
	var controller = createController()
	inicializeRoutes(router, controller)
	return router
}

func inicializeRoutes(router *gin.Engine, controller *controller.InfectedController) {

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/infected", controller.Post)
	router.GET("/infecteds", controller.List)
}

func createController() *controller.InfectedController {
	infectedRepository := memory.NewInfectedRepository()
	return &controller.InfectedController{
		Usecase: usecase.NewInfectedUsecase(infectedRepository),
	}
}
