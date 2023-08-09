package main

import (
	"github.com/rherlt/reval/internal/api/evaluationapi"
	"github.com/rherlt/reval/internal/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//TODO: remove for production builds.
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))
	//TODO: https://github.com/deepmap/oapi-codegen/blob/master/examples/petstore-expanded/gin/petstore.go#L21C1-L48C1

	r.GET("/openapi.json", controller.GetSwagger)

	si := new(controller.EvaluationApiServerInterface)
	evaluationapi.RegisterHandlersWithOptions(r, si, si.GetServerOptions())

	r.Run(":8080")
}
