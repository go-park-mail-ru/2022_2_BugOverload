package pkg

import "time"

const (
	ContentTypeJSON = "application/json"
	ContentTypeJPEG = "image/jpeg"

	BufSizeRequest = 1024 * 1024 * 4
	BufSizeImage   = 1024 * 1024 * 2
	TokenS3Length  = 20

	CookieValueLength = 40
	TimeoutLiveCookie = 10 * time.Hour

	SignupRequest         = "SignupRequest"
	LoginRequest          = "LoginRequest"
	AuthRequest           = "AuthRequest"
	LogoutRequest         = "LogoutRequest"
	InCinemaRequest       = "InCinemaRequest"
	PopularRequest        = "PopularRequest"
	RecommendationRequest = "RecommendationRequest"
	DownloadImageRequest  = "DownloadImageRequest"
	UploadImageRequest    = "UploadImageRequest"
	ChangeImageRequest    = "ChangeImageRequest"
)

type ContextType string

var CookieKey ContextType = "cookie"
var LoggerKey ContextType = "logger"
