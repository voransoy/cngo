package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ "status" : "OKAY" })
	})

	r.GET("/t/:host", testWebsite)

	// It detects PORT env
	r.Run()
}

func testWebsite(c *gin.Context) {
	host := "https://" + c.Param("host")
	resp, err := http.Get(host)
	if err != nil {
		log.Printf("Couldn't test website %s: %v", host, err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "ERROR", "error" : err.Error()})
	}
	if resp.StatusCode != 200 {
		log.Printf("Website %s didin't respond with OK", host)
		c.JSON(http.StatusOK, gin.H{"website": host, "status": "FAIL"})
	} else {
		log.Printf("Website %s is OK", host)
		c.JSON(http.StatusOK, gin.H{"website": host, "status": "OK"})
	}
}
