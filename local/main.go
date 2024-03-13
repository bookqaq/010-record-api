package local

import (
	"010-record-api/config"
	"010-record-api/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() http.Handler {
	// TODO: replace from gin to go1.22 ServerMux, avoid jit assembler usage github.com/chenzhuoyu/iasm
	// f**k bytedance for adding this sh**.
	// local := http.NewServeMux()
	// local.Handle("POST /movie-upload/{filename}", MovieUploadContext)
	local := gin.New()
	local.Use(gin.RecoveryWithWriter(logger.GetWriter()))

	// record api group
	groupMovie := local.Group("/movie")
	initRouterGroupMovie(groupMovie)

	// handler for embedded web patcher
	local.StaticFS("/patcher", http.FS(MustGetPatcher()))

	// tell user way to access patcher
	fmt.Printf("\tto access embedded patcher, go http://%s/patcher/\n\n", config.Config.ListenAddress)

	// not setting this handler into group /movie
	local.PUT("/movie-upload/:filename", MovieUploadContext)

	return local
}

func initRouterGroupMovie(group *gin.RouterGroup) {
	group.GET("/server/status", MovieServerStatus)
	group.POST("/sessions/new", MovieSessionNew)
	group.POST("/sessions/:sid/videos/:vid/:operation", MovieUploadManagement)
}
