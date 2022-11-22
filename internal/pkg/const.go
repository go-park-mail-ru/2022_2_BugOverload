package pkg

import (
	"database/sql"
	"time"
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
	UserCollectionsRequest       = "UserCollectionsRequest"

	DateFormat = "2006.01.02"
	TimeFormat = "15:04:05"
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

// Service
const (
	CollectionTargetTag         = "tag"
	CollectionTargetGenre       = "genre"
	CollectionTargetProdCompany = "prod_company"
	CollectionTargetProdCountry = "prod_country"

	CollectionSortParamDate       = "date"
	CollectionSortParamFilmRating = "rating"

	UserCollectionsSortParamCreateDate = "create_time"
	UserCollectionsSortParamUpdateDate = "update_time"
	UserCollectionsDelimiter           = "now"

	MaxCountAttrInCollection = 2
)

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
