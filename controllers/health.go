package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthStatus struct {
	Status string `json:"status"`
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, HealthStatus{Status: "OK"})
}
