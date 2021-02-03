package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func jsonOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func jsonError(c *gin.Context, status int, err interface{}) {
	var message interface{}

	switch v := err.(type) {
	case error:
		message = v.Error()
	case nil:
		message = http.StatusText(status)
	default:
		message = v
	}

	c.AbortWithStatusJSON(status, gin.H{"error": message})
}

func notFound(c *gin.Context, err interface{}) {
	jsonError(c, http.StatusNotFound, err)
}

func badRequest(c *gin.Context, err interface{}) {
	jsonError(c, http.StatusBadRequest, err)
}

func serverError(c *gin.Context, err interface{}) {
	jsonError(c, http.StatusInternalServerError, err)
}
