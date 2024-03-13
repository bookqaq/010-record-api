package local

const constUploadStatusBegin = "begin-upload"
const constUploadStatusEnd = "end-upload"

// mock session id with a fixed(maybe) length 5 (v5 = *v4 == 5;)
//
// TODO: remove this if detailed session and file management is needed
const mockSID = "12345"

const (
	APIPatcher = "/patcher"
)

const (
	APIMovie = "/movie"

	APIServerStatus = APIMovie + "/server/status"

	APIMovieSession       = APIMovie + "/sessions"
	APIMovieSessionNew    = APIMovieSession + "/new"
	APIMovieSessionManage = APIMovieSession + "/{sid}/videos/{vid}/{operation}"
)

const (
	APIDedicatedMovieUpload = "/movie-upload/{filename}"
)
