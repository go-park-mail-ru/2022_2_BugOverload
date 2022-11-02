package fillerdb

import (
	"context"
	"time"

	"github.com/pkg/errors"

	pkgInner "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func (f *DBFiller) uploadReviews() (int, error) {
	countInserts := len(f.faceReviews)

	insertStatement, countAttributes := pkgInner.CreateStatement(insertReviews, countInserts)

	insertStatement += insertReviewsEnd

	values := make([]interface{}, countAttributes*countInserts)

	for idx, value := range f.faceReviews {
		posAttr := 0
		posValue := idx * countAttributes

		values[posValue+posAttr] = value.Name
		posAttr++
		values[posValue+posAttr] = value.Type
		posAttr++
		values[posValue+posAttr] = value.Time
		posAttr++
		values[posValue+posAttr] = value.Body
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := pkgInner.SendQuery(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "uploadReviews")
	}
	defer rows.Close()

	count := 1
	for idx := range f.faceReviews {
		f.faceReviews[idx].ID = count
		count++
	}

	return countInserts, nil
}

func (f *DBFiller) linkReviewsLikes() (int, error) {
	countInserts := f.Config.Volume.CountReviewsLikes

	insertStatement, countAttributes := pkgInner.CreateStatement(insertReviewsLikes, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	appended := 0

	for _, value := range f.faceReviews {
		count := pkg.RandMaxInt(f.Config.Volume.MaxLikesOnReview)
		if (countInserts - appended) < count {
			count = countInserts - appended
		}

		sequence := pkg.CryptoRandSequence(f.faceUsers[len(f.faceUsers)-1].ID+1, f.faceUsers[0].ID)

		for j := 0; j < count; j++ {
			values[pos] = value.ID
			pos++
			values[pos] = sequence[j]
			pos++
		}

		appended += count
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := pkgInner.SendQuery(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkReviewsLikes")
	}
	defer rows.Close()

	return countInserts, nil
}
