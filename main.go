package main

import (
    "github.com/gin-gonic/gin"
)

func index (c *gin.Context){

    content := gin.H{"url": "http://www.flipkart.com/search?q=pendrive&sid=6bo/jdy&as=on&as-show=on&otracker=start&as-pos=1_1_ic_pen"}
    c.JSON(200, content)
}

func main(){
  app := gin.Default()
  app.GET("/", index)
  app.Run(":3000")
}