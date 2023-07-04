package utils

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ReturnError(c *gin.Context, err error, message string) {
	fmt.Fprint(os.Stderr, err.Error())
	c.HTML(http.StatusOK, "error.html", gin.H{
		"Code":    http.StatusInternalServerError,
		"Message": message,
	})
}
