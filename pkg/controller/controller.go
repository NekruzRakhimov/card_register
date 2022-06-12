package controller

import (
	"card_register/models"
	"card_register/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddNewOrderInfo(c *gin.Context) {
	var info models.Info
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.AddNewOrderInfo(info); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "данные успешно сохранены"})
}

func GetAllInfo(c *gin.Context) {
	info, err := service.GetAllInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, info)
}
