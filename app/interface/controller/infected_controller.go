package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kosegor/go-covid19-api/app/domain/model"
	"github.com/kosegor/go-covid19-api/app/usecase"
)

type InfectedController struct {
	Usecase usecase.InfectedUsecase
}

func (i InfectedController) Post(c *gin.Context) {
	var newInfected model.Infected
	err := c.BindJSON(&newInfected)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	newInfected, apiError := i.Usecase.Post(newInfected)

	if apiError != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, apiError.Error())
		return
	}

	c.JSON(http.StatusCreated, newInfected)
	return
}

func (i InfectedController) List(c *gin.Context) {
	infecteds, apiError := i.Usecase.List()

	if apiError != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, apiError.Error())
		return
	}

	c.JSON(http.StatusOK, infecteds)
	return
}
