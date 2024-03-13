package local

import (
	"010-record-api/config"
	"010-record-api/utils"
	"fmt"
	"net/http"
)

func New() http.Handler {
	// replace from gin to go1.22 ServerMux, avoid jit assembler usage github.com/chenzhuoyu/iasm
	// f**k bytedance for adding this sh**.

	local := http.NewServeMux()

	// record api group
	initRouterGroupMovie(local)

	// tell user way to access patcher
	fmt.Printf("\tto access embedded patcher, go http://%s/patcher/\n\n", config.Config.ListenAddress)

	// movie upload put request, not setting this handler into group /movie
	local.HandleFunc(utils.RequestURL(http.MethodPut, APIDedicatedMovieUpload), MovieUploadContext)

	// handler for embedded web patcher, refer to MustGetPatcher() for routing details
	local.Handle(APIPatcher+"/", http.FileServerFS(MustGetPatcher()))

	return local
}

func initRouterGroupMovie(group *http.ServeMux) {
	group.HandleFunc(utils.RequestURL(http.MethodGet, APIServerStatus), MovieServerStatus)
	group.HandleFunc(utils.RequestURL(http.MethodPost, APIMovieSessionNew), MovieSessionNew)
	group.HandleFunc(utils.RequestURL(http.MethodPost, APIMovieSessionManage), MovieUploadManagement)
}
