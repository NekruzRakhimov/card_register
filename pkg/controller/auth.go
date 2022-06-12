package controller

import (
	"card_register/models"
	"card_register/pkg/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func SingIn(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	token, err := service.GenerateToken(user.Login, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": "empty auth header"})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": "invalid auth header"})
		return
	}

	userId, err := service.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int64, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "user id not found"})
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "user id is of invalid type"})
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}
