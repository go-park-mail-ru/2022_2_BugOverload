package errorshandlers

import "errors"

var (
	ErrContentTypeUndefined     = errors.New("content-type undefined")
	ErrUnsupportedMediaType     = errors.New("unsupported media type")
	ErrEmptyFieldAuth           = errors.New("request has empty fields (nickname | email | password)")
	ErrLoginCombinationNotFound = errors.New("no such combination of user and password")
)
