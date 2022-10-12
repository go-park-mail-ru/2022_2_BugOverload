package params

import "time"

type ContextType string

var CookieKey ContextType = "cookie"

const countSeconds = 10

var ContextTimeout = time.Duration(countSeconds) * time.Second

const (
	CookieValueLength = 40
	TimeoutLiveCookie = 10 * time.Hour
)
