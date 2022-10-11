package models

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
)

type UserLogoutRequest struct {
	user models.User
}

func (ulr *UserLogoutRequest) Bind(w http.ResponseWriter, r *http.Request) error {
	if r.Header.Get("Cookie") == "" {
		return errors.NewErrAuth(errors.ErrNoCookie)
	}

	return nil
}

func (ulr *UserLogoutRequest) GetUser() *models.User {
	return &ulr.user
}
