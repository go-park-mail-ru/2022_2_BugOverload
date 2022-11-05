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
	RecommendationRequest = "RecommendationRequest"
	DownloadImageRequest  = "DownloadImageRequest"
	UploadImageRequest    = "UploadImageRequest"
	ChangeImageRequest    = "ChangeImageRequest"
	GetUserProfileRequest = "GetUserProfileRequest"
	TagCollectionRequest  = "TagCollectionRequest"
	GetPersonRequest      = "GetPersonRequest"
	FilmRequest           = "FilmRequest"

	// Images params request
	ImageObjectFilmPosterHor = "film_poster_hor"
	ImageObjectFilmPosterVer = "film_poster_ver"
	ImageObjectDefault       = "default"
	ImageObjectAvatar        = "user_avatar"

	// ParamsContextKeys
	GetReviewsParams       = "GetReviewsRequestParams"
	GetPersonParams        = "GetPersonParams"
	GetCollectionTagParams = "GetPersonParams"

	// Crypt
	SaltLength     = 16
	ArgonTime      = 1
	ArgonMemory    = 32 * 1024
	ArgonThreads   = 4
	ArgonKeyLength = 32
)

type ContextKeyType string

// SessionKey LoggerKey for ctx
var SessionKey ContextKeyType = "cookie"
var LoggerKey ContextKeyType = "logger"

// GetReviewsParamsKey for RequestParams
var GetReviewsParamsKey ContextKeyType = GetPersonParams

// GetCollectionTagParamsKey for RequestParams
var GetCollectionTagParamsKey ContextKeyType = GetCollectionTagParams

// TxDefaultOptions for Postgres
var TxDefaultOptions = &sql.TxOptions{
	Isolation: sql.LevelDefault,
	ReadOnly:  true,
}

// GetPersonParamsCtx in struct for GetPersonParams in personHandler
type GetPersonParamsCtx struct {
	CountFilms int
}

// GetCollectionTagParamsCtx in struct for GetPersonParams in personHandler
type GetCollectionTagParamsCtx struct {
	Tag        string
	CountFilms int
	Delimiter  string
}
