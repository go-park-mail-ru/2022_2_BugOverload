package pkg

import (
	"database/sql"
	"time"
)

const (
	// Validation HTTP
	ContentTypeJSON = "application/json"
	ContentTypeJPEG = "image/webp"

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

	// Images params request
	ImageObjectFilmPosterHor = "film_poster_hor"
	ImageObjectFilmPosterVer = "film_poster_ver"
	ImageObjectDefault       = "default"
	ImageObjectAvatar        = "user_avatar"

	// GetReviewsParams
	GetReviewsParams = "GetReviewsRequestParams"

	// User
	GetUserProfile = "GetUserProfile"

	// Crypt
	SaltLength     = 16
	ArgonTime      = 1
	ArgonMemory    = 32 * 1024
	ArgonThreads   = 4
	ArgonKeyLength = 32

	// csrf
	CsrfSecretDefault = "J25qeHRobmpkc2NyZmN0cmh0biEhIQ=="
	CsrfSecretLength  = 32
)

type ContextType string

var SessionKey ContextType = "cookie"
var LoggerKey ContextType = "logger"

var GetReviewsParamsKey ContextType = GetReviewsParams

// TxDefaultOptions for Postgres
var TxDefaultOptions = &sql.TxOptions{
	Isolation: sql.LevelDefault,
	ReadOnly:  true,
}
