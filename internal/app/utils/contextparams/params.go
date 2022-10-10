package contextparams

import "time"

type ContextType string

var CookieKey ContextType = "cookie"

const countSeconds = 2

var ContextTimeout = time.Duration(countSeconds) * time.Second
