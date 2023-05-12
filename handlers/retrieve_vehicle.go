package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RetrieveVehicle(c *gin.Context) {
	vin := c.Param("vin")

	log.Println(vin)
	c.JSON(http.StatusOK, nil)
}
