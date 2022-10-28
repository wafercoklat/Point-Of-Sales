package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, err error) bool {
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return true // signal that there was an error and the caller should return
	}
	return false // no error, can continue
}
