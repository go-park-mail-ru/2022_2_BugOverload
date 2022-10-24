package pkg

import "time"

const (
	// Validation HTTP
	ContentTypeJSON = "application/json"
	ContentTypeJPEG = "image/jpeg"

	// Validattion size Requests
	BufSizeRequest = 1024 * 1024 * 2.5
	BufSizeImage   = 1024 * 1024 * 2

	// Cookie
	CookieValueLength = 40
	TimeoutLiveCookie = 10 * time.Hour

	// Handler factory
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

	// S3 params request
	ImageObjectFilmPosterHor = "film_poster_hor"
	ImageObjectFilmPosterVer = "film_poster_ver"
	ImageObjectDefault       = "default"
	ImageObjectAvatar        = "user_avatar"
)

type ContextType string

var SessionKey ContextType = "cookie"
var LoggerKey ContextType = "logger"
