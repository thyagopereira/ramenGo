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

func (w *WebServer) AddNewCreateOrderRoute(path string, uc *usecases.CreateOrderUseCase) {
	w.Router.POST(path, func(c *gin.Context) {
		var requestDTO usecases.CreateOrderRequestDTO

		if err := c.BindJSON(&requestDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if requestDTO.BrothId == "" || requestDTO.ProteinId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "both brothId and proteinId are required"})
			return
		}

		response, err := uc.Execute(&requestDTO)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not place order"})
			return
		}

		c.JSON(http.StatusCreated, response)

	})
}
