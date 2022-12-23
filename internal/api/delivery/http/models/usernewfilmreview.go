package models

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

//go:generate easyjson  -disallow_unknown_fields usernewfilmreview.go

//easyjson:json
type NewFilmReviewRequest struct {
	FilmID     int
	ReviewName string `json:"name,omitempty" example:"Почему игра престолов это всего лишь пустое насилие?"`
	ReviewType string `json:"type,omitempty" example:"negative"`
	ReviewBody string `json:"body,omitempty" example:"много много текса"`
}

func NewNewFilmReviewRequest() NewFilmReviewRequest {
	return NewFilmReviewRequest{}
}

func (f *NewFilmReviewRequest) Bind(r *http.Request) error {
	var err error

	vars := mux.Vars(r)

	f.FilmID, err = strconv.Atoi(vars["id"])
	if err != nil {
		return errors.ErrConvertQueryType
	}

	if r.Header.Get("Content-Type") == "" {
		return errors.ErrContentTypeUndefined
	}

	if r.Header.Get("Content-Type") != constparams.ContentTypeJSON {
		return errors.ErrUnsupportedMediaType
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.ErrBadBodyRequest
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	if len(body) == 0 {
		return errors.ErrEmptyBody
	}

	err = easyjson.Unmarshal(body, f)
	if err != nil {
		return err
		//  return errors.ErrJSONUnexpectedEnd
	}

	if f.ReviewName == "" || f.ReviewType == "" || f.ReviewBody == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	return nil
}

func (f *NewFilmReviewRequest) GetReview() *models.Review {
	return &models.Review{
		Name: f.ReviewName,
		Type: f.ReviewType,
		Body: f.ReviewBody,
	}
}

func (f *NewFilmReviewRequest) GetParams() *constparams.NewFilmReviewParams {
	return &constparams.NewFilmReviewParams{
		FilmID: f.FilmID,
	}
}

//easyjson:json
type FilmReviewResponse struct {
	UserID int `json:"user_id,omitempty" example:"13"`
}

func NewFilmReviewResponse(user *models.User) *FilmReviewResponse {
	return &FilmReviewResponse{
		UserID: user.ID,
	}
}
