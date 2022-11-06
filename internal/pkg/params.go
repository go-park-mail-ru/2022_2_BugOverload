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
	PersonRequest         = "PersonRequest"
	FilmRequest           = "FilmRequest"
	ReviewsFilmRequest    = "ReviewsFilmRequest"

	// Images params request
	ImageObjectFilmPosterHor = "film_poster_hor"
	ImageObjectFilmPosterVer = "film_poster_ver"
	ImageObjectDefault       = "default"
	ImageObjectAvatar        = "user_avatar"

	// ParamsContextKeys
	GetReviewsParams       = "GetReviewsParamsKey"
	GetPersonParams        = "GetPersonParams"
	GetCollectionTagParams = "GetPersonParams"
	GetFilmParams          = "GetFilmParams"

	// Crypt
	SaltLength     = 16
	ArgonTime      = 1
	ArgonMemory    = 32 * 1024
	ArgonThreads   = 4
	ArgonKeyLength = 32

	// PersonProfessions
	Actor    = 1
	Artist   = 7
	Director = 2
	Writer   = 3
	Producer = 4
	Operator = 5
	Montage  = 8
	Composer = 6
)

type ContextKeyType string

// SessionKey LoggerKey for ctx
var SessionKey ContextKeyType = "cookie"
var LoggerKey ContextKeyType = "logger"

// GetPersonParamsKey for RequestParams
var GetPersonParamsKey ContextKeyType = GetPersonParams

// GetPersonParamsCtx in struct for GetPersonParams in personHandler
type GetPersonParamsCtx struct {
	CountFilms  int
	CountImages int
}

// GetCollectionTagParamsKey for RequestParams
var GetCollectionTagParamsKey ContextKeyType = GetCollectionTagParams

// GetCollectionTagParamsCtx in struct for GetPersonParams in tagCollectionHandler
type GetCollectionTagParamsCtx struct {
	Tag        string
	CountFilms int
	Delimiter  string
}

// GetReviewsParamsKey for RequestParams
var GetReviewsParamsKey ContextKeyType = GetReviewsParams

// GetReviewsFilmParamsCtx in struct for GetReviewsParamsKey in reviewHandler
type GetReviewsFilmParamsCtx struct {
	FilmID int
	Count  int
	Offset int
}

// GetFilmParamsKey for RequestParams
var GetFilmParamsKey ContextKeyType = GetFilmParams

// GetFilmParamsCtx in struct for GetFilmParams in filmHandler
type GetFilmParamsCtx struct {
	CountImages int
}

// TxDefaultOptions for Postgres
var TxDefaultOptions = &sql.TxOptions{
	Isolation: sql.LevelDefault,
	ReadOnly:  true,
}
