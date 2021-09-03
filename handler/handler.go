package handler

import (
	tech "git.wildberries.ru/finance/go-infrastructure/tech/v2"
	"github.com/gin-gonic/gin"
	service "github.com/proggcreator/wb-lib/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	tech.New().SetAppInfo("stuff-shk", "1.0.0")

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
