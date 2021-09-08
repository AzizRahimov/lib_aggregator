package handler

import (
	"fmt"
	"github.com/AzizRahimov/lib_aggregator/models"
	"github.com/AzizRahimov/lib_aggregator/pkg/gateways"
	"github.com/AzizRahimov/lib_aggregator/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type Inter interface {
	InitRoutes() *gin.Engine
	getFromProc(c *gin.Context)
}

type handler struct {
	port     string
	token    string
	dataproc func(payload models.RequestFromProc) models.Response
}

//TODO: Поменяй название
func NewHandler(port string, token string, dataproc func(payload models.RequestFromProc) models.Response) Inter {
	return &handler{port: port, token: token, dataproc: dataproc}
}

func (h *handler) InitRoutes() *gin.Engine {
	r := gin.New()
	logging.Setup()
	//setting.Setup()
	r.Use(logging.Logger(logrus.StandardLogger()), gin.Recovery())
	r.Use(logging.Logger(logrus.StandardLogger()))
	r.GET("/ping", h.routetest)
	r.POST("/", h.getFromProc)
	//просто новый ком

	//r.Run(":" + setting.AppConfig.Server.Port)
	// и вот это в main'e вынести
	if err := r.Run(":" + h.port); err != nil {
		logrus.Fatalf("error occired while running server %s", err.Error())
	}

	return r

}

func (h *handler) getFromProc(c *gin.Context) {

	var payload models.RequestFromProc //
	var authenticated bool
	if values, _ := c.Request.Header["Authorization"]; len(values) > 0 {
		//pretty.Logln("[EPHandler Check token ]:", values[0]+" "+setting.AppConfig.Server.Token)
		pretty.Logln("[EPHandler Check token ]:", values[0]+" "+h.token)

		//	authenticated = gateways.CheckToken(values[0], setting.AppConfig.Server.Token)
		authenticated = gateways.CheckToken(values[0], h.token)
	}
	if !authenticated {
		m := "authentication error"
		pretty.Logln("[EPHandler]:", m)
		c.XML(http.StatusNetworkAuthenticationRequired, models.Response{Message: m})
		return
	}
	log.Println("begin!!!!!!")

	err := c.ShouldBindXML(&payload)
	if err != nil {
		//
		c.XML(http.StatusBadRequest, gin.H{
			"error2": err,
		})
		return
	}
	//c.XML(http.StatusOK, gateways.Handler(payload))

	c.XML(http.StatusOK, h.dataproc(payload))
}

func (h *handler) routetest(c *gin.Context) {
	var payload models.RequestFromProc
	fmt.Println("Это пустой payload", payload)
	fmt.Println(payload.Command, "ЭТО КОМАНДА, не должна быть пустой")
	c.ShouldBindJSON(&payload)
	fmt.Println("Тут должны уже быть данные", payload)

	//c.XML(http.StatusOK, h.dataproc(payload))

	c.JSON(200, gin.H{
		"message": "Pongggg",
	})
}


