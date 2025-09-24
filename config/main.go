package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/bookqaq/010-record-api/utils"

	"github.com/pelletier/go-toml/v2"
)

var Config struct {
	ListenAddress                 string `toml:"listen_address" comment:"proxy listen address(ip:port), should NOT match upstream_server_url"`
	UploadServiceAddress          string `toml:"upload_service_address" comment:"upload service address(ip:port), where video will be truly uploaded to. \nShould be a domain:port(if needed) or ip:port(if needed) as above, \nbut with real computer ip from command like ipconfig, etc. "`
	VideoSaveDirectory            string `toml:"video_save_dir" comment:"directory to save uploaded video"`
	LogPath                       string `toml:"log_path" comment:"log file path"`
	FeatureXrpcIIDXMusicMovieInfo *bool  `toml:"feature_xrpcIIDXMusicMovieInfo" comment:"whether to add video owner in the filename. \nIf true, the filename will be like 'video_owner-something_else.mp4\nExtra development on xrpc server is needed\n\t - sending IIDX00music.movieinfo request data to POST /feature/xrpcIIDXMusicMovieInfo\nFor more info, see server/local/feature/xrpc_iidx_movieinfo.go."`
}

func MustParse() {
	data, err := os.ReadFile("./config.toml")
	if err != nil {
		utils.SimulatedPanic(fmt.Errorf("config read failed: %w", err))
	}
	if err := toml.Unmarshal(data, &Config); err != nil {
		utils.SimulatedPanic(fmt.Errorf("config parse failed: %w", err))
	}

	Config.ListenAddress, _ = strings.CutPrefix(Config.ListenAddress, "http://")
	Config.ListenAddress, _ = strings.CutSuffix(Config.ListenAddress, "/")
	Config.UploadServiceAddress, _ = strings.CutPrefix(Config.UploadServiceAddress, "http://")
	Config.UploadServiceAddress, _ = strings.CutSuffix(Config.UploadServiceAddress, "/")
	Config.VideoSaveDirectory, _ = strings.CutSuffix(Config.VideoSaveDirectory, "/")
}

// if config.toml is not exist, create a default one and exit
func CheckFile() {
	if _, err := os.Stat("./config.toml"); os.IsNotExist(err) {
		Config.ListenAddress = "127.0.0.1:4399"
		Config.UploadServiceAddress = "127.0.0.1:4399"
		Config.VideoSaveDirectory = "./video"
		Config.LogPath = "./log.txt"
		Config.FeatureXrpcIIDXMusicMovieInfo = new(bool)
		*Config.FeatureXrpcIIDXMusicMovieInfo = false

		data, err := toml.Marshal(Config)
		if err != nil {
			utils.SimulatedPanic(fmt.Errorf("create config failed: %w", err))
		}
		if err := os.WriteFile("./config.toml", data, 0644); err != nil {
			utils.SimulatedPanic(fmt.Errorf("create config failed: %w", err))
		}

		fmt.Println("config.toml created, edit with notepad and then reopen executable...")
		fmt.Scanln()
		os.Exit(0)
	}
}
