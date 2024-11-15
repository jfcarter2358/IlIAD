// main.go

package main

import (
	"fmt"
	"iad/config"
	"iad/provider"
	"log"
	"math/rand"
	"net/rpc"
	"time"

	logger "github.com/jfcarter2358/go-logger"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	config.LoadConfig()
	logger.SetLevel(config.Config.LogLevel)
	logger.SetFormat(config.Config.LogFormat)

	// http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: config.Config.TLSSkipVerify}

	router = gin.New()
	router.Use(gin.LoggerWithFormatter(logger.ConsoleLogFormatter))
	router.Use(gin.Recovery())

	logger.Infof("", "Running with port: %d", config.Config.Port)

	initializeRoutes()

	rand.Seed(time.Now().UnixNano())

	provider.Providers = make(map[string]*provider.Provider)

	for name, conf := range config.Config.Providers {
		logger.Infof("", "Connecting to %s at %s", name, conf.URL)
		provider.Providers[name] = &provider.Provider{URL: conf.URL}
		var err error
		provider.Providers[name].Client, err = rpc.DialHTTP("tcp", conf.URL)
		if err != nil {
			log.Fatal("dialing:", err)
		}

	}

	routerPort := fmt.Sprintf(":%d", config.Config.Port)
	router.Run(routerPort)
}
