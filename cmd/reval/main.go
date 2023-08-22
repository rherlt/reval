package main

import (
	"log"
	"net/http"

	"github.com/rherlt/reval/internal/api/evaluationapi"
	"github.com/rherlt/reval/internal/config"
	"github.com/rherlt/reval/internal/controller"
	"github.com/rherlt/reval/internal/data"
	"github.com/rherlt/reval/internal/oidc"
	"github.com/rherlt/reval/internal/persistence"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config.Configure()

	//setup gin webserver
	r := gin.Default()

	//setup http redirect from app root / to web base url, e.g. /ui/
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, config.Current.Gin_Web_BaseUrl)
	})

	//register folder for static web deployment
	r.Static("/ui/", config.Current.Gin_Web_Path)

	//setup CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, config.Current.Gin_Cors_AdditionalAllowedHeaders...)
	corsConfig.AllowAllOrigins = config.Current.Gin_Cors_AllowAllOrigins
	r.Use(cors.New(corsConfig))
	r.Use(oidc.OidcAuthMiddleware(config.Current.Oidc_Authority, config.Current.Oidc_Audience))

	//register HTTP handlers for evaluatio api
	si := new(controller.EvaluationApiServerInterface)
	evaluationapi.RegisterHandlersWithOptions(r, si, si.GetServerOptions())

	importRequired := !persistence.DbExistis()

	//setup database
	err := persistence.SetupDb()
	defer persistence.CloseClient()

	if err != nil {
		log.Fatal("cannot setup database: ", err)
	}

	if importRequired {
		err = data.ImportData()
		if err != nil {
			log.Fatal("cannot import data: ", err)
		}
	}

	//run webserver
	r.Run(config.Current.Gin_WebServerAddress)
}
