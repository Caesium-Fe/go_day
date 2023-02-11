package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RedirectTest(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://10.175.94.58:7002/#/")
}

func RedirectURL2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Hello": "NMSL"})
}
