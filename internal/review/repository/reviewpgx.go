package repository

import (
	"context"
	"database/sql"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type ReviewRepository interface {
	GetReviewsByFilmID(ctx context.Context) ([]models.Review, error)
}

// reviewPostgres is implementation repository of Postgres corresponding to the ReviewRepository interface.
type reviewPostgres struct {
	database *sqltools.Database
}

// NewReviewPostgres is constructor for reviewPostgres.
func NewReviewPostgres(database *sqltools.Database) ReviewRepository {
	return &reviewPostgres{
		database,
	}
}

func (r *reviewPostgres) GetReviewsByFilmID(ctx context.Context) ([]models.Review, error) {
	response := make([]ReviewSQL, 0)

	params, _ := ctx.Value(innerPKG.GetReviewsParamsKey).(innerPKG.GetReviewsFilmParamsCtx)

	// Reviews - Main
	errTX := sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, r.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowsReviews, err := tx.QueryContext(ctx, getReviewsByFilmID, params.FilmID, params.Count, params.Offset)
		if err != nil {
			return err
		}
		defer rowsReviews.Close()

		for rowsReviews.Next() {
			var review ReviewSQL

			err = rowsReviews.Scan(
				&review.Name,
				&review.Type,
				&review.Body,
				&review.CountLikes,
				&review.CreateTime,
				&review.Author.ID,
				&review.Author.Nickname,
				&review.Author.Avatar,
				&review.Author.CountReviews)
			if err != nil {
				return err
			}

			response = append(response, review)
		}

		return nil
	})

	// the main entity is not found
	if stdErrors.Is(errTX, sql.ErrNoRows) || len(response) == 0 {
		return []models.Review{}, errors.ErrNotFoundInDB
	}

	if errTX != nil {
		return []models.Review{}, errors.ErrPostgresRequest
	}

	res := make([]models.Review, len(response))

	for idx, value := range response {
		res[idx] = value.Convert()
	}

	return res, nil
}