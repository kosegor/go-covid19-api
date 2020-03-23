package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kosegor/go-covid19-api/app/domain/model"
	"github.com/kosegor/go-covid19-api/app/usecase"
)

type IncidentController struct {
	Usecase usecase.IncidentUsecase
}

func (i *IncidentController) Post(c *gin.Context) {
	var newIncident model.Incident
	err := c.BindJSON(&newIncident)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = newIncident.Validate()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	newIncident, apiError := i.Usecase.Post(newIncident)

	if apiError != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, apiError.Error())
		return
	}

	c.JSON(http.StatusCreated, newIncident)
	return
}

func (i *IncidentController) List(c *gin.Context) {
	incidents, apiError := i.Usecase.List()

	if apiError != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, apiError.Error())
		return
	}

	c.JSON(http.StatusOK, incidents)
	return
}
