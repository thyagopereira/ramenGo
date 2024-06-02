package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramenGo/domain/usecases"
)

type WebServer struct {
	Router *gin.Engine
}

func NewWebServer() *WebServer {
	return &WebServer{
		Router: gin.Default(),
	}
}

func (w *WebServer) AddNewGetRoute(path string, uc usecases.Usecase) {
	w.Router.GET(path, func(c *gin.Context) {
		result, err := uc.Execute()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		json, err := json.Marshal(result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		c.JSON(http.StatusAccepted, json)
		return
	})
}
