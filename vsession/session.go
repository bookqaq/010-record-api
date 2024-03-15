package vsession

import (
	"fmt"
	"sync"
	"time"
)

// TODO: clean outdated records with trigger
var MapInfo *sync.Map = new(sync.Map)

type Info struct {
	ShopName  string
	MD5Sum    string
	MusicId   int // id in music_data.bin
	Timestamp int64
}

// info to filename
func (info *Info) ToFileName() string {
	// TODO: option to get music name from given music_data.bin
	localTime := time.Unix(info.Timestamp, 0).Local().Format("20060102-150405")
	return fmt.Sprintf("%s-%d-%s", localTime, info.MusicId, info.MD5Sum)
}
