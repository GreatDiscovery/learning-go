package gin

import (
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
	"time"
)

func testResponse(c *gin.Context) {
	c.JSON(http.StatusGatewayTimeout, gin.H{
		"code": http.StatusGatewayTimeout,
		"msg":  "timeout",
	})
}

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(3000*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}

func TestTimeout(t *testing.T) {
	r := gin.New()
	r.Use(timeoutMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		time.Sleep(5000 * time.Millisecond)
		c.Status(http.StatusOK)
	})
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
