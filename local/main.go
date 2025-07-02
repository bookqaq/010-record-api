package local

import (
	"fmt"
	"net/http"

	"github.com/bookqaq/010-record-api/config"
	"github.com/bookqaq/010-record-api/local/feature"
	"github.com/bookqaq/010-record-api/utils"
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

	// handler for feature api, for example receive some xrpc request data(1.1.0+)
	// refer to github.com/bookqaq/010-record-api/local/feature
	initRouterGroupFeature(local)

	return local
}

func initRouterGroupMovie(group *http.ServeMux) {
	group.HandleFunc(utils.RequestURL(http.MethodGet, APIServerStatus), MovieServerStatus)
	group.HandleFunc(utils.RequestURL(http.MethodPost, APIMovieSessionNew), MovieSessionNew)
	group.HandleFunc(utils.RequestURL(http.MethodPost, APIMovieSessionManage), MovieUploadManagement)
}

func initRouterGroupFeature(group *http.ServeMux) {
	group.HandleFunc(utils.RequestURL(http.MethodPost, APIFeatureXrpcIIDXMovieInfo), feature.FeatureXrpcIIDXMusicMovieInfo)
}
