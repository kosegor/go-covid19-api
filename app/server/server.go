package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kosegor/go-covid19-api/app/interface/controller"
	"github.com/kosegor/go-covid19-api/app/interface/persistence/dynamo"
	"github.com/kosegor/go-covid19-api/app/usecase"
)

func CreateServer() *gin.Engine {
	var router = gin.Default()
	var controller = createController()
	inicializeRoutes(router, controller)
	return router
}

func inicializeRoutes(router *gin.Engine, controller *controller.IncidentController) {

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/incident", controller.Post)
	router.GET("/incidents", controller.List)
}

func createController() *controller.IncidentController {
	incidentRepository := dynamo.NewIncidentRepository()
	return &controller.IncidentController{
		Usecase: usecase.NewIncidentUsecase(incidentRepository),
	}
}
