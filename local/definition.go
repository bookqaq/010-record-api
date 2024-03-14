package local

type requestMovieSessionUploadBegin struct {
	EA3LocationId string
	EA3ShopName   string
	EA3SystemId   string // a pcbid
	MD5Sum        string
	MusicId       int // id in music_data.bin
	NoteType      int
	PlayScore     int // old style score
	PlayerNames   []string
	// PlayerRefIds []int // not sure about the type
	Timestamp int64
	Timezone  string
	VideoKbps int
	VideoType string // seems like only value "short"
}
