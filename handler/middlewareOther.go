package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) checkJsonType(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-type")
	if contentType != "application/json" {
		logrus.Fatalf("Invalid content-type")

	}

}

func (h *Handler) acceptJsonOrXml(c *gin.Context) {
	listAccept := c.Request.Header["Accept"]
	for _, cur := range listAccept {
		if (cur == "application/json") || (cur == "application/xml") {
			return
		}
	}
	logrus.Fatalf("Invalid accept header")
}
