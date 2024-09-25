// main.go

package main

import (
	"${APP_NAME}/config"
	"context"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	logger "github.com/jfcarter2358/go-logger"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func run(ctx context.Context, channel chan struct{}) {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	config.LoadConfig()
	logger.SetLevel(config.Config.LogLevel)
	logger.SetFormat(config.Config.LogFormat)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: config.Config.TLSSkipVerify}

	router = gin.New()
	router.Use(gin.LoggerWithFormatter(logger.ConsoleLogFormatter))
	router.Use(gin.Recovery())

	logger.Infof("", "Running with port: %d", config.Config.Port)

	initializeRoutes()

	rand.Seed(time.Now().UnixNano())

	routerPort := fmt.Sprintf(":%d", config.Config.Port)
	router.Run(routerPort)
}

//	@title			${APP_NAME_CAPITAL} Swagger API
//	@version		0.0.1
//	@description	${APP_DESCRIPTION}
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	${AUTHOR_NAME}
//	@contact.url	https://github.com/${GITHUB_ORG}/${APP_NAME}/issues
//	@contact.email	${AUTHOR_EMAIL}

// @license.name	MIT
// @license.url	https://opensource.org/license/mit/
func main() {
	channel := make(chan struct{})
	ctx, _ := context.WithCancel(context.Background())
	run(ctx, channel)
}
