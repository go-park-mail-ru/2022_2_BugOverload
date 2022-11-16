package pkg

import (
	"database/sql"
	"time"
)

// ImagesParams
const (
	// Images params request
	ImageObjectFilmPosterHor = "film_poster_hor" // key - 1 to 1 from prev request
	ImageObjectFilmPosterVer = "film_poster_ver" // key - 1 to 1 from prev request
	ImageObjectFilmImage     = "film_image"      // key - filmID/filmImageKey - Example 1/2
	ImageObjectDefault       = "default"         // key - login or signup
	ImageObjectUserAvatar    = "user_avatar"     // key - 1 to 1 from prev request
	ImageObjectPersonAvatar  = "person_avatar"   // key - 1 to 1 from prev request
	ImageObjectPersonImage   = "person_image"    // key - personID/personImageKey - Example 5/12

	// Image Def Images
	DefFilmPosterHor = "hor"
	DefFilmPosterVer = "ver"
	DefUserAvatar    = "avatar"
	DefPersonAvatar  = "avatar"

	ImageCountSignupLogin = 12

	// S3
	FilmsBucket   = "films/"
	DefBucket     = "default/"
	PersonsBucket = "persons/"
	UsersBucket   = "users/"
)

// Handler factory
const (
	SignupRequest                = "SignupRequest"
	LoginRequest                 = "LoginRequest"
	AuthRequest                  = "AuthRequest"
	LogoutRequest                = "LogoutRequest"
	RecommendationRequest        = "RecommendationRequest"
	DownloadImageRequest         = "DownloadImageRequest"
	UploadImageRequest           = "UploadImageRequest"
	ChangeImageRequest           = "ChangeImageRequest"
	GetUserProfileRequest        = "GetUserProfileRequest"
	TagCollectionRequest         = "TagCollectionRequest"
	PersonRequest                = "PersonRequest"
	FilmRequest                  = "FilmRequest"
	ReviewsFilmRequest           = "ReviewsFilmRequest"
	GetUserSettingsRequest       = "GetUserSettingsRequest"
	PutUserSettingsRequest       = "PutUserSettingsRequest"
	FilmRateRequest              = "FilmRateRequest"
	FilmRateDropRequest          = "FilmRateDropRequest"
	NewFilmReviewRequest         = "NewFilmReviewRequest"
	GetUserActivityOnFilmRequest = "GetUserActivityOnFilmRequest"

	DateFormat = "2006.01.02"
	TimeFormat = "15:04:05"
)

// DB
const (
	// PersonProfessions
	Actor    = 1
	Artist   = 7
	Director = 2
	Writer   = 3
	Producer = 4
	Operator = 5
	Montage  = 8
	Composer = 6

	TagFromPopular  = "popular"
	TagInPopular    = "популярное"
	TagFromInCinema = "in_cinema"
	TagInInCinema   = "сейчас в кино"

	DefTypeFilm   = "film"
	DefTypeSerial = "serial"

	TypeReviewPositive = "positive"
	TypeReviewNegative = "negative"
	TypeReviewNeutral  = "neutral"

	DefGender = "male"
	OnlyDate  = "2006"
)

// TxDefaultOptions for Postgres
var TxDefaultOptions = &sql.TxOptions{
	Isolation: sql.LevelDefault,
	ReadOnly:  true,
}

// TxInsertOptions for Postgres
var TxInsertOptions = &sql.TxOptions{
	Isolation: sql.LevelDefault,
	ReadOnly:  false,
}

// Global
const (
	// HTTPMode
	HTTPS = "https"

	// Validation HTTP
	ContentTypeJSON              = "application/json"
	ContentTypeMultipartFormData = "multipart/form-data"
	ContentTypeWEBP              = "image/webp"
	ContentTypeJPEG              = "image/jpeg"
	ContentTypePNG               = "image/png"

	// Validattion size Requests
	BufSizeRequest = 1024 * 1024 * 2.5
	BufSizeImage   = 1024 * 1024 * 2

	// Cookie
	SessionCookieName = "session_id"
	CookieValueLength = 40
	TimeoutLiveCookie = 10 * time.Hour
	GlobalCookiePath  = "/"

	// Crypt
	SaltLength     = 16
	ArgonTime      = 1
	ArgonMemory    = 16 * 1024
	ArgonThreads   = 4
	ArgonKeyLength = 16

	// csrf
	CsrfSecretDefault = "J25qeHRobmpkc2NyZmN0cmh0biEhIQ=="
	CsrfSecretLength  = 32
)

type ContextKeyType string

// SessionKey LoggerKey for ctx
var SessionKey ContextKeyType = "cookie"
var LoggerKey ContextKeyType = "logger"

// CurrentUserKey is key for ctx in auth middleware
var CurrentUserKey ContextKeyType = "current-user"
