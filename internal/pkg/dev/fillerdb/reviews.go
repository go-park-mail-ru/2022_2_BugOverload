package fillerdb

import (
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func (f *DBFiller) uploadReviews() (int, error) {
	countInserts := len(f.faceReviews)

	insertStatement, countAttributes := pkg.CreateStatement(insertReviews, countInserts)

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

	target := "reviews"

	stmt, rows, cancelFunc, err := pkg.SendQuery(f.DB.Connection, f.Config.Database.Timeout, insertStatement, target, values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
	defer rows.Close()

	counter := 0
	var insertID int64
	for rows.Next() {
		err = rows.Scan(&insertID)
		if err != nil {
			logrus.Errorf("Error [%s] when getting insertID [%s]", err, target)
			return 0, err
		}

		f.faceReviews[counter].ID = int(insertID)
		counter++
	}

	return countInserts, nil
}

func (f *DBFiller) linkReviewsLikes() (int, error) {
	countInserts := f.Config.Volume.CountReviewsLikes

	insertStatement, countAttributes := pkg.CreateStatement(insertReviewsLikes, countInserts)

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

	stmt, rows, cancelFunc, err := pkg.SendQuery(f.DB.Connection, f.Config.Database.Timeout, insertStatement, "reviews likes", values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
	defer rows.Close()

	return countInserts, nil
}
