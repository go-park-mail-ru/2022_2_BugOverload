package models

import (
	"context"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type UserLogoutRequest struct{}

func (ulr *UserLogoutRequest) Bind(r *http.Request) (context.Context, error) {
	if r.Header.Get("Cookie") == "" {
		return nil, errors.NewErrAuth(errors.ErrNoCookie)
	}

	cookie := r.Cookies()[0]

	ctx := context.WithValue(r.Context(), pkg.SessionKey, cookie.Value)

	return ctx, nil
}
