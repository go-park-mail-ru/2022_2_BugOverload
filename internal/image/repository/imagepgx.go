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

func (i *imageS3WithPostgres) UpdateImageInfo(ctx context.Context, image *models.Image) error {
	err := sqltools.RunTxOnConn(ctx, innerPKG.TxInsertOptions, i.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, updateUserAvatar, image.Key)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return stdErrors.WithMessagef(errors.ErrPostgresRequest,
			"Err: params input: query - [%s], values - [%s]. Special Error [%s]",
			updateUserAvatar, image.Key, err)
	}

	return nil
}
