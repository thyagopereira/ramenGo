package api

import (
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

func (w *WebServer) AddNewGetBrothsRoute(path string, uc *usecases.ListBrothsUseCase) {
	w.Router.GET(path, func(c *gin.Context) {
		result, err := uc.Execute()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, result.Broths)
		return
	})
}

func (w *WebServer) AddNewGetProteinsRoute(path string, uc *usecases.ListProteinsUseCase) {
	w.Router.GET(path, func(c *gin.Context) {
		result, err := uc.Execute()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, result.Proteins)
		return
	})
}
