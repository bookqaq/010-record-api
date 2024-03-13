package main

import (
	"010-record-api/config"
	"010-record-api/initialize"
	"010-record-api/logger"
	"010-record-api/proxy"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
)

func main() {
	isDebug := flag.Bool("debug", false, "enable debug mode for http server and log")
	loglevel := flag.String("loglevel", logger.LOGLEVELWARNING, "enable debug mode for gin and log")
	flag.Parse()

	if !(*isDebug) {
		gin.SetMode(gin.ReleaseMode)
	}
	banner()
	config.CheckFile()                   // generate config file and exit if not exist
	config.MustParse()                   // read config
	logger.New(*isDebug, *loglevel)      // init logger
	initialize.CreateRequiredDirectory() // check video directory
	go handleInterrupt()
	proxy.Start() // start proxy
}

func handleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		logger.Warning.Printf("receive ^C, bye!\n")
		os.Exit(0)
	}()
}

func banner() {
	fmt.Println("\n010 record api -- lightning model video upload handler")
	fmt.Println()
}
