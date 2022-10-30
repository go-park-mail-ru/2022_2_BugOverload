package models

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type UserPutSettingsRequest struct {
	Nickname string `json:"nickname,omitempty" example:"StepByyyy"`
	Password string `json:"password,omitempty" example:"Widget Adapter"`
}

func NewPutUserSettingsRequest() *UserPutSettingsRequest {
	return &UserPutSettingsRequest{}
}

func (u *UserPutSettingsRequest) Bind(r *http.Request) (context.Context, error) {
	if r.Header.Get("Cookie") == "" {
		return nil, errors.NewErrAuth(errors.ErrNoCookie)
	}

	if r.Header.Get("Content-Type") == "" {
		return nil, errors.NewErrValidation(errors.ErrContentTypeUndefined)
	}

	if r.Header.Get("Content-Type") != pkg.ContentTypeJSON {
		return nil, errors.NewErrValidation(errors.ErrUnsupportedMediaType)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	if len(body) == 0 {
		return nil, errors.NewErrValidation(errors.ErrEmptyBody)
	}

	err = json.Unmarshal(body, u)
	if err != nil {
		return nil, errors.NewErrValidation(errors.ErrCJSONUnexpectedEnd)
	}

	cookie := r.Cookies()[0]

	ctx := context.WithValue(r.Context(), pkg.SessionKey, cookie.Value)

	return ctx, nil
}
