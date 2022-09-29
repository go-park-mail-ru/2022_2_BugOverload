package errorshandlers

import "errors"

var (
	ContentTypeUndefined     = errors.New("content-type undefined")
	UnsupportedMediaType     = errors.New("unsupported media type")
	EmptyFieldAuth           = errors.New("request has empty fields (nickname | email | password)")
	LoginCombinationNotFound = errors.New("no such combination of user and password")
)
