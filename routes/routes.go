package routes

import (
	"card_register/pkg/controller"
	"card_register/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func RunAllRoutes() {
	r := gin.Default()

	// Использование CORS
	r.Use(controller.CORSMiddleware())

	// Установка Logger-а
	utils.SetLogger()

	// Форматирование логов
	utils.FormatLogs(r)

	// Статус код 500, при любых panic()
	r.Use(gin.Recovery())

	//r.Use(limits.RequestSizeLimiter(100))

	// Запуск end-point'ов
	runAllRoutes(r)

	// Запуск сервера
	runServer(r)

}

func runServer(r *gin.Engine) {
	var (
		port string
		host string
	)
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = utils.AppSettings.AppParams.PortRun
		host = utils.AppSettings.AppParams.ServerURL
	} else {
		host = "0.0.0.0"
	}
	err := r.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Println(err)
	}
}

func runAllRoutes(r *gin.Engine) {
	r.GET("/", Check)
	r.GET("/ping", Ping)

	r.POST("/info", controller.AddNewOrderInfo)

	r.POST("/auth", controller.SingIn)
	api := r.Group("/api", controller.UserIdentity)
	api.GET("/info", controller.GetAllInfo)
}

func Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"reason": "up and working"})
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"reason": "pong"})
}
