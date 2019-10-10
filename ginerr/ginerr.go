package ginerr

import (
	"net/http"

	"github.com/afternoob/afterror"
	"github.com/gin-gonic/gin"
)

func RespWithError(c *gin.Context, err error) {
	if e, ok := err.(*afterror.Error); ok {
		c.JSON(e.Code, gin.H{
			"type":    e.Type,
			"message": e.Message,
		})

		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"type":    "UnknownType",
		"message": err.Error(),
	})
}
