package main

import (
	"log"
	"os"

	"github.com/rherlt/reval/internal/api/evaluationapi"
	"github.com/rherlt/reval/internal/config"
	"github.com/rherlt/reval/internal/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//load config path from command line or use "." (current application path)
	var configPath *string = nil
	if len(os.Args) > 0 {
		configPath = &os.Args[1:][0]
	}

	err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, config.Current.Gin_Cors_AdditionalAllowedHeaders...)
	corsConfig.AllowAllOrigins = config.Current.Gin_Cors_AllowAllOrigins
	r.Use(cors.New(corsConfig))
	//TODO: https://github.com/deepmap/oapi-codegen/blob/master/examples/petstore-expanded/gin/petstore.go#L21C1-L48C1

	r.GET("/openapi.json", controller.GetSwagger)

	//register HTTP handlers
	si := new(controller.EvaluationApiServerInterface)
	evaluationapi.RegisterHandlersWithOptions(r, si, si.GetServerOptions())

	//register folder for static web deployment
	r.Static("/ui", config.Current.Gin_Web_Path)

	//run webserver
	r.Run(config.Current.WebServerAddress)
}
