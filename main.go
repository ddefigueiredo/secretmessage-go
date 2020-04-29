package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/neufeldtech/smsg-go/pkg/middleware"
	"github.com/neufeldtech/smsg-go/pkg/redis"
	"github.com/neufeldtech/smsg-go/pkg/secretmessage"

	"os"
)

func main() {

	redis.Init()
	secretmessage.InitSlackClient()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(middleware.ValidateSignature())

	r.POST("/slash", secretmessage.HandleSlash)
	r.POST("/interactive", secretmessage.HandleInteractive)

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 64)
	if err != nil {
		port = 8080
	}
	r.Run(fmt.Sprintf("0.0.0.0:%v", port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
