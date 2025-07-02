package vsession

import (
	"fmt"
	"sync"
	"time"
)

// TODO: clean outdated records with trigger
var MapInfo *sync.Map = new(sync.Map)

type Info struct {
	ShopName       string
	MD5Sum         string
	MusicId        int // id in music_data.bin. does anyone want a music_data.bin parser plus a song mapper?
	Timestamp      int64
	VideoOwnerId   string // optional. video owner's id, for example 8-digit iidxid.
	VideoOwnerName string // optional. video owner's name, this is not considered unique, but adding this makes file more disguishable.
}

// info to filename
func (info *Info) ToFileName() string {
	// TODO: option to get music name from given music_data.bin
	localTime := time.Unix(info.Timestamp, 0).Local().Format("20060102-150405")
	return fmt.Sprintf("%s-%d-%s", localTime, info.MusicId, info.MD5Sum)
}

func (info *Info) ToFileNameWithOwner() string {
	// TODO: option to get music name from given music_data.bin
	localTime := time.Unix(info.Timestamp, 0).Local().Format("20060102-150405")
	videoOwnerId := info.VideoOwnerId
	if videoOwnerId == "" {
		videoOwnerId = "null"
	}
	return fmt.Sprintf("%s-%s-%s-%d-%s", info.VideoOwnerName, info.VideoOwnerId, localTime, info.MusicId, info.MD5Sum)
}
