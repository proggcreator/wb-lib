package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/proggcreator/wb-Restful/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/")
	{
		mylists := api.Group("/api/v1", h.checkClaim)
		{
			mylists.POST("/employee_add", h.checkJsonType, h.employee_add)
			mylists.DELETE("/employee_remove", h.employee_remove)
			mylists.PUT("/employee_upd", h.checkJsonType, h.employee_upd)
			mylists.GET("/get_all", h.acceptJsonOrXml, h.get_all)
			mylists.GET("/employee_get/:id", h.employee_get)
			mylists.GET("/tech/info", h.employee_tech)

		}
	}

	return router
}
