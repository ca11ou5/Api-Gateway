package main

import (
	"api-gateway/configs"
	"api-gateway/pkg/auth"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	c := configs.InitConfig()

	r := gin.Default()
	authSvc := *auth.RegisterRoutes(r, &c)
	fmt.Println(authSvc)

	r.Run(c.Port)
}
