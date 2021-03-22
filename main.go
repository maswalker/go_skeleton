package main

import (
  "time"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  grace "github.com/facebookgo/grace/gracehttp"
  mw "skeleton/middleware"
)

func pong(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "pong",
  })
}

func main() {
  gin.SetMode(gin.ReleaseMode)

  router := gin.New()

  router.Use(mw.Logger())

  router.Use(gin.Logger())

  router.Use(gin.Recovery())

  router.Use(cors.Default())

  router.GET("/ping", pong)

  srv := &http.Server {
    Addr: ":8080",
    Handler: router,
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }

  grace.Serve(srv)
}

