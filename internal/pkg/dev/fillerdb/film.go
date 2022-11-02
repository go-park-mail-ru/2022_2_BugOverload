package fillerdb

import (
	"github.com/go-faker/faker/v4"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func (f *DBFiller) uploadFilms() (int, error) {
	countInserts := len(f.filmsSQL)

	insertStatement, countAttributes := createStatement(insertFilms, countInserts)

	insertStatement += insertFilmsEnd

	values := make([]interface{}, countAttributes*countInserts)

	for idx, value := range f.filmsSQL {
		posAttr := 0
		posValue := idx * countAttributes

		values[posValue+posAttr] = value.Name
		posAttr++
		values[posValue+posAttr] = value.ProdYear
		posAttr++
		values[posValue+posAttr] = value.PosterVer
		posAttr++
		values[posValue+posAttr] = value.PosterHor
		posAttr++
		values[posValue+posAttr] = value.Description
		posAttr++
		values[posValue+posAttr] = value.ShortDescription
		posAttr++
		values[posValue+posAttr] = value.OriginalName
		posAttr++
		values[posValue+posAttr] = value.Slogan
		posAttr++
		values[posValue+posAttr] = value.AgeLimit
		posAttr++
		values[posValue+posAttr] = value.BoxOffice
		posAttr++
		values[posValue+posAttr] = value.Budget
		posAttr++
		values[posValue+posAttr] = value.Duration
		posAttr++
		values[posValue+posAttr] = value.CurrencyBudget
		posAttr++
		values[posValue+posAttr] = value.Type
		posAttr++
		values[posValue+posAttr] = value.CountSeasons
		posAttr++
		values[posValue+posAttr] = value.EndYear
	}

	target := "films"

	stmt, rows, cancelFunc, err := f.SendQuery(insertStatement, target, values)
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

		f.films[counter].ID = int(insertID)
		counter++
	}

	return countInserts, nil
}

func (f *DBFiller) linkFilmsReviews() (int, error) {
	countInserts := len(f.faceReviews)

	insertStatement, countAttributes := createStatement(insertFilmsReviews, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	appended := 0

	sequenceReviews := pkg.CryptoRandSequence(f.faceReviews[len(f.faceReviews)-1].ID+1, f.faceReviews[0].ID)

	for _, value := range f.films {
		countPartBatch := pkg.RandMaxInt(f.Config.Volume.MaxReviewsOnFilm)
		if (countInserts - appended) < countPartBatch {
			countPartBatch = countInserts - appended
		}

		sequenceUsers := pkg.CryptoRandSequence(f.faceUsers[len(f.faceUsers)-1].ID+1, f.faceUsers[0].ID)

		for j := 0; j < countPartBatch; j++ {
			values[pos] = sequenceReviews[appended:][j]
			pos++
			values[pos] = sequenceUsers[j]
			pos++
			values[pos] = value.ID
			pos++
		}

		appended += countPartBatch
	}

	stmt, rows, cancelFunc, err := f.SendQuery(insertStatement, "film reviews", values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
	defer rows.Close()

	return countInserts, nil
}

func (f *DBFiller) linkFilmGenres() (int, error) {
	countInserts := 0

	for _, value := range f.films {
		countInserts += len(value.Genres)
	}

	insertStatement, countAttributes := createStatement(insertFilmsGenres, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	offset := 0
	posValue := 0

	for _, value := range f.films {
		posValue += offset
		offset = 0
		weight := len(value.Genres)

		for _, genre := range value.Genres {
			values[posValue+offset] = value.ID
			offset++
			values[posValue+offset] = f.genres[genre]
			weight--
			offset++
			values[posValue+offset] = weight
			offset++
		}
	}

	stmt, rows, cancelFunc, err := f.SendQuery(insertStatement, "film genres", values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
	defer rows.Close()

	return countInserts, nil
}

func (f *DBFiller) linkFilmCountries() (int, error) {
	countInserts := 0

	for _, value := range f.films {
		countInserts += len(value.ProdCountries)
	}

	insertStatement, countAttributes := createStatement(insertFilmsCountries, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	offset := 0
	posValue := 0

	for _, value := range f.films {
		posValue += offset
		offset = 0
		weight := len(value.ProdCountries)

		for _, country := range value.ProdCountries {
			values[posValue+offset] = value.ID
			offset++
			_, ok := f.countries[country]
			if !ok {
				logrus.Error(country)
			}
			values[posValue+offset] = f.countries[country]
			weight--
			offset++
			values[posValue+offset] = weight
			offset++
		}
	}

	stmt, rows, cancelFunc, err := f.SendQuery(insertStatement, "film countries", values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
	defer rows.Close()

	return countInserts, nil
}

func (f *DBFiller) linkFilmCompanies() (int, error) {
	countInserts := 0

	for _, value := range f.films {
		countInserts += len(value.ProdCompanies)
	}

	insertStatement, countAttributes := createStatement(insertFilmsCompanies, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	offset := 0
	posValue := 0

	for _, value := range f.films {
		posValue += offset
		offset = 0
		weight := len(value.ProdCompanies)

		for _, company := range value.ProdCompanies {
			values[posValue+offset] = value.ID
			offset++
			values[posValue+offset] = f.companies[company]
			weight--
			offset++
			values[posValue+offset] = weight
			offset++
		}
	}

	stmt, rows, cancelFunc, err := f.SendQuery(insertStatement, "film companies", values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
	defer rows.Close()

	return countInserts, nil
}

func (f *DBFiller) linkFilmPersons() (int, error) {
	values := make([]interface{}, 0)

	countInserts := 0

	for _, value := range f.films {
		countActors := pkg.RandMaxInt(f.Config.Volume.MaxFilmActors) + 1
		weightActors := countActors - 1

		sequenceActors := pkg.CryptoRandSequence(f.persons[len(f.persons)-1].ID+1, f.persons[0].ID)

		for i := 0; i < countActors; i++ {
			values = append(values, sequenceActors[i], value.ID, f.professions["актер"], pkg.NewSQLNullString(faker.Word()), weightActors)
			weightActors--
		}

		for profession := 2; profession < len(f.professions); profession++ {
			countPersons := pkg.RandMaxInt(f.Config.Volume.MaxFilmPersons) + 1
			weightPersons := countPersons - 1

			sequencePersons := pkg.CryptoRandSequence(f.persons[len(f.persons)-1].ID+1, f.persons[0].ID)

			for i := 0; i < countPersons; i++ {
				values = append(values, sequencePersons[i], value.ID, profession, pkg.NewSQLNullString(""), weightPersons)
				weightPersons--
			}

			countInserts += countPersons
		}

		countInserts += countActors
	}

	insertStatement, _ := createStatement(insertFilmsPersons, countInserts)

	stmt, rows, cancelFunc, err := f.SendQuery(insertStatement, "film persons", values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
	defer rows.Close()

	return countInserts, nil
}

func (f *DBFiller) linkFilmTags() (int, error) {
	countInserts := 0

	values := make([]interface{}, 0)

	for _, value := range f.tags {
		count := pkg.RandMaxInt(f.Config.Volume.MaxFilmsInTag) + 1

		sequence := pkg.CryptoRandSequence(f.films[len(f.films)-1].ID+1, f.films[0].ID)

		for i := 0; i < count; i++ {
			values = append(values, sequence[i], value)
		}

		countInserts += count
	}

	insertStatement, _ := createStatement(insertFilmsTags, countInserts)

	stmt, rows, cancelFunc, err := f.SendQuery(insertStatement, "film tags", values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
	defer rows.Close()

	return countInserts, nil
}
