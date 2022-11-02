package models

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type UserLogoutRequest struct{}

func (ulr *UserLogoutRequest) Bind(r *http.Request) error {
	if r.Header.Get("Cookie") == "" {
		return errors.NewErrAuth(errors.ErrNoCookie)
	}

	return nil
}
