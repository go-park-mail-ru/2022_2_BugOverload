package errorshandlers

import "errors"

var (
	ErrContentTypeUndefined     = errors.New("content-type undefined")
	ErrUnsupportedMediaType     = errors.New("unsupported media type")
	ErrEmptyFieldAuth           = errors.New("request has empty fields (nickname | email | password)")
	ErrLoginCombinationNotFound = errors.New("no such combination of user and password")
	ErrUserExist                = errors.New("such user exist")
	ErrUserNotExist             = errors.New("such user doesn't exist")
	ErrSignupUserExist          = errors.New("a user with such a mail already exists")
	ErrNoCookie                 = errors.New("request not have cookie")
	ErrCookieNotExist           = errors.New("no such cookie")
)
