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

// game-related API (WebAPI2)
const (
	APIMovie = "/movie"

	APIServerStatus = APIMovie + "/server/status"

	APIMovieSession       = APIMovie + "/sessions"
	APIMovieSessionNew    = APIMovieSession + "/new"
	APIMovieSessionManage = APIMovieSession + "/{sid}/videos/{vid}/{operation}"
)

// dedicated movie upload entry. since you can set video upload address in
// APIMovieSessionManage, this entry is not considered as a part of WebAPI2
const (
	APIDedicatedMovieUpload = "/movie-upload/{key}"
)

// feature API that exposed to potential users. these might not necessary for
// singleplay, but useful when running on a shared arcade or you just want to
// implement your own feature.
const (
	APIFeature = "/feature"

	APIFeatureXrpcIIDXMovieInfo = APIFeature + "/xrpcIIDXMovieInfo"
)
