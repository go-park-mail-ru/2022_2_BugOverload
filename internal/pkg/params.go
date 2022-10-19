package pkg

import "time"

const (
	ContentTypeJSON  = "application/json"
	ContentTypeImage = "image/jpeg"

	BufSizeImage  = 1024 * 1024 * 24
	TokenS3Length = 20

	CookieValueLength = 40
	TimeoutLiveCookie = 10 * time.Hour
)

type ContextType string

var CookieKey ContextType = "cookie"
var LoggerKey ContextType = "logger"
