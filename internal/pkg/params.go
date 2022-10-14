package pkg

import "time"

type ContextType string

var CookieKey ContextType = "cookie"
var LoggerKey ContextType = "logger"

const (
	CookieValueLength = 40
	TimeoutLiveCookie = 10 * time.Hour
)
