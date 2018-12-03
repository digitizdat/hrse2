package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/static"
    "os"
)

func main() {
    port := ":" + os.Getenv("PORT")

    r := gin.Default()
    r.Use(static.Serve("/", static.LocalFile("./public", true)))
    r.Use(static.Serve("/favicon.ico", static.LocalFile("/img/hrse.ico", true)))

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "title": "HRSE",
        })
    })
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    r.Run(port)
}

