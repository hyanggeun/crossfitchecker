package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)


func main(){
    router := gin.Default()

    router.GET("/hello",func (c *gin.Context) {
        c.JSON(http.StatusOK,map[string]string{
            "hello": "world",
        })
    })
    router.Run(":8080")
}
