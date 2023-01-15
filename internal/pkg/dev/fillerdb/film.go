package fillerdb

import (
	"context"
	"math"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/pkg/errors"

	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/dev/devpkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func (f *DBFiller) uploadFilms() (int, error) {
	// Defining sending parameters
	query := insertFilms
	message := "uploadFilms"

	globalCountInserts := len(f.filmsSQL)

	countAttributes := strings.Count(query, ",") + 1

	maxInsertValues := devpkg.MaxInsertValuesSQL / countAttributes

	action := func(countInserts int, values []interface{}) error {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
		defer cancelFunc()

		insertStatement := sqltools.CreateFullQuery(query, countInserts, countAttributes)

		_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
		if err != nil {
			return errors.Wrap(err, message)
		}

		return nil
	}

	pos := 0

	values := make([]interface{}, maxInsertValues*countAttributes)

	countInserts := 0

	for i := 0; i < len(f.films); i++ {
		values[pos] = f.filmsSQL[i].Name
		pos++
		values[pos] = f.filmsSQL[i].ProdYear
		pos++
		values[pos] = f.filmsSQL[i].PosterVer
		pos++
		values[pos] = f.filmsSQL[i].PosterHor
		pos++
		values[pos] = f.filmsSQL[i].Description
		pos++
		values[pos] = f.filmsSQL[i].ShortDescription
		pos++
		values[pos] = f.filmsSQL[i].OriginalName
		pos++
		values[pos] = f.filmsSQL[i].Slogan
		pos++
		values[pos] = f.filmsSQL[i].AgeLimit
		pos++
		values[pos] = f.filmsSQL[i].BoxOfficeDollars
		pos++
		values[pos] = f.filmsSQL[i].Budget
		pos++
		values[pos] = f.filmsSQL[i].DurationMinutes
		pos++
		values[pos] = f.filmsSQL[i].CurrencyBudget
		pos++
		values[pos] = f.filmsSQL[i].Type
		pos++

		countInserts++

		if devpkg.MaxInsertValuesSQL-pos < 20 || i == len(f.films)-1 {
			values = values[:pos]

			err := action(countInserts, values)
			if err != nil {
				return 0, errors.Wrap(err, message)
			}

			pos = 0

			countInserts = 0
		}
	}

	for j := 0; j < globalCountInserts; j++ {
		f.films[j].ID = j + 1
	}

	return globalCountInserts, nil
}

func (f *DBFiller) uploadFilmsMedia() (int, error) {
	globalCountInserts := 0

	ids := make([]int, 0)

	for idx := range f.films {
		if f.films[idx].Ticket != "" || f.films[idx].Trailer != "" {
			globalCountInserts++

			ids = append(ids, idx)
		}
	}

	// Defining sending parameters
	query := insertFilmsMedia
	message := "uploadFilmsMedia"

	countAttributes := strings.Count(query, ",") + 1

	maxInsertValues := devpkg.MaxInsertValuesSQL / countAttributes

	action := func(countInserts int, values []interface{}) error {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
		defer cancelFunc()

		insertStatement := sqltools.CreateFullQuery(query, countInserts, countAttributes)

		_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
		if err != nil {
			return errors.Wrap(err, message)
		}

		return nil
	}

	pos := 0

	values := make([]interface{}, maxInsertValues*countAttributes)

	countInserts := 0

	for i := 0; i < len(ids); i++ {
		values[pos] = f.films[ids[i]].ID
		pos++
		values[pos] = f.filmsSQL[ids[i]].Ticket
		pos++
		values[pos] = f.filmsSQL[ids[i]].Trailer
		pos++

		countInserts++

		if devpkg.MaxInsertValuesSQL-pos < 20 || i == len(ids)-1 {
			values = values[:pos]

			err := action(countInserts, values)
			if err != nil {
				return 0, errors.Wrap(err, message)
			}

			pos = 0

			countInserts = 0
		}
	}

	return globalCountInserts, nil
}

func (f *DBFiller) uploadSerials() (int, error) {
	globalCountInserts := 0

	ids := make([]int, 0)

	for idx := range f.films {
		if f.films[idx].Type == innerPKG.DefTypeSerial {
			globalCountInserts++

			ids = append(ids, idx)
		}
	}

	// Defining sending parameters
	query := insertSerials
	message := "uploadSerials"

	countAttributes := strings.Count(query, ",") + 1

	maxInsertValues := devpkg.MaxInsertValuesSQL / countAttributes

	action := func(countInserts int, values []interface{}) error {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
		defer cancelFunc()

		insertStatement := sqltools.CreateFullQuery(query, countInserts, countAttributes)

		_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
		if err != nil {
			return errors.Wrap(err, message)
		}

		return nil
	}

	pos := 0

	values := make([]interface{}, maxInsertValues*countAttributes)

	countInserts := 0

	for i := 0; i < len(ids); i++ {
		values[pos] = f.films[ids[i]].ID
		pos++
		values[pos] = f.filmsSQL[ids[i]].CountSeasons
		pos++
		values[pos] = f.filmsSQL[ids[i]].EndYear
		pos++

		countInserts++

		if devpkg.MaxInsertValuesSQL-pos < 20 || i == len(ids)-1 {
			values = values[:pos]

			err := action(countInserts, values)
			if err != nil {
				return 0, errors.Wrap(err, message)
			}

			pos = 0

			countInserts = 0
		}
	}

	return globalCountInserts, nil
}

func (f *DBFiller) linkFilmsReviews() (int, error) {
	countInserts := len(f.Reviews)

	countAttributes := strings.Count(insertFilmsReviews, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertFilmsReviews, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	appended := 0

	sequenceReviewsID := pkg.CryptoRandSequence(len(f.Reviews)+1, 1)

	for _, value := range f.films {
		count := pkg.RandMaxInt(f.Config.Volume.MaxReviewsOnFilm)
		if (countInserts - appended) < count {
			count = countInserts - appended
		}

		if count == 0 {
			continue
		}

		sequenceUsersID := pkg.CryptoRandSequence(count+1, 1)

		for j := 0; j < count; j++ {
			curID := sequenceReviewsID[appended:][j]

			values[pos] = curID
			pos++
			values[pos] = sequenceUsersID[j]
			pos++
			values[pos] = value.ID
			pos++

			// For update denormal fields
			switch f.Reviews[curID-1].Type {
			case devpkg.TypeReviewNegative:
				f.films[value.ID-1].CountNegativeReviews++
			case devpkg.TypeReviewNeutral:
				f.films[value.ID-1].CountNeutralReviews++
			case devpkg.TypeReviewPositive:
				f.films[value.ID-1].CountPositiveReviews++
			}
		}

		appended += count
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkFilmsReviews")
	}

	return countInserts, nil
}

func (f *DBFiller) linkFilmGenresReal() (int, error) {
	countInserts := 0

	for _, value := range f.films {
		countInserts += len(value.Genres)
	}

	countAttributes := strings.Count(insertFilmsGenres, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertFilmsGenres, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0

	for _, value := range f.films {
		weight := len(value.Genres)
		for _, genre := range value.Genres {
			values[pos] = value.ID
			pos++
			values[pos] = f.guides.Genres[genre]
			weight--
			pos++
			values[pos] = weight
			pos++
		}
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkFilmGenresReal")
	}

	return countInserts, nil
}

func (f *DBFiller) linkFilmCountriesReal() (int, error) {
	countInserts := 0

	for _, value := range f.films {
		countInserts += len(value.ProdCountries)
	}

	countAttributes := strings.Count(insertFilmsCountries, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertFilmsCountries, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0

	for _, value := range f.films {
		weight := len(value.ProdCountries)

		for _, country := range value.ProdCountries {
			values[pos] = value.ID
			pos++
			values[pos] = f.guides.Countries[country]
			weight--
			pos++
			values[pos] = weight
			pos++
		}
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkFilmCountriesReal")
	}

	return countInserts, nil
}

func (f *DBFiller) linkFilmCompaniesReal() (int, error) {
	countInserts := 0

	for _, value := range f.films {
		countInserts += len(value.ProdCompanies)
	}

	countAttributes := strings.Count(insertFilmsCompanies, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertFilmsCompanies, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0

	for _, value := range f.films {
		weight := len(value.ProdCompanies)

		for _, company := range value.ProdCompanies {
			values[pos] = value.ID
			pos++
			values[pos] = f.guides.Companies[company]
			weight--
			pos++
			values[pos] = weight
			pos++
		}
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkFilmCompaniesReal")
	}

	return countInserts, nil
}

func (f *DBFiller) linkFilmPersonsRandom() (int, error) {
	// Defining sending parameters
	query := insertFilmsPersons
	message := "linkFilmPersonsRandom"

	globalCountInserts := 0

	countAttributes := strings.Count(query, ",") + 1

	maxInsertValues := devpkg.MaxInsertValuesSQL / countAttributes

	action := func(countInserts int, values []interface{}) error {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
		defer cancelFunc()

		insertStatement := sqltools.CreateFullQuery(query, countInserts, countAttributes)

		_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
		if err != nil {
			return errors.Wrap(err, message)
		}

		return nil
	}

	pos := 0

	values := make([]interface{}, maxInsertValues*countAttributes)

	countInserts := 0

	maxPersons := int(math.Max(float64(f.Config.Volume.MaxFilmsActors), float64(f.Config.Volume.MaxFilmsPersons)))

	for i := 0; i < len(f.films); i++ {
		countActors := pkg.RandMaxInt(f.Config.Volume.MaxFilmsActors) + 1

		sequencePersons := pkg.CryptoRandSequence(maxPersons+1, 1)

		if countActors != 0 {
			weightActors := countActors - 1

			for j := 0; j < countActors; j++ {
				values[pos] = sequencePersons[j]
				pos++
				values[pos] = f.films[i].ID
				pos++
				values[pos] = f.guides.Professions["актер"]
				pos++
				values[pos] = sqltools.NewSQLNullString(faker.Word())
				weightActors--
				pos++
				values[pos] = weightActors
				pos++

				f.films[i].CountActors = countActors
			}

			countInserts += countActors
		}

		countPersons := pkg.RandMaxInt(f.Config.Volume.MaxFilmsPersons) + 1
		if countPersons == 0 {
			continue
		}

		weightPersons := countPersons - 1

		end := 2

		sequenceProfessions := pkg.CryptoRandSequence(len(f.guides.Professions)+1, end)

		for j := 0; j < countPersons; j++ {
			values[pos] = sequencePersons[j]
			pos++
			values[pos] = f.films[i].ID
			pos++
			values[pos] = sequenceProfessions[j]
			pos++
			values[pos] = sqltools.NewSQLNullString("")
			weightPersons--
			pos++
			values[pos] = weightPersons
			pos++
		}

		countInserts += countPersons

		if devpkg.MaxInsertValuesSQL-pos < 50 || i == len(f.films)-1 {
			err := action(countInserts, values[:pos])
			if err != nil {
				return 0, errors.Wrap(err, message)
			}

			pos = 0

			globalCountInserts += countInserts

			countInserts = 0
		}
	}

	return globalCountInserts, nil
}

func (f *DBFiller) linkFilmPersonsReal() (int, error) {
	values := make([]interface{}, 0)

	countInserts := 0

	for _, value := range f.films {
		countActors := len(value.Actors)
		if countActors > 0 {
			weightActors := countActors - 1

			for i := 0; i < countActors; i++ {
				values = append(
					values,
					f.mapPersons[value.Actors[i].Name],
					value.ID,
					f.guides.Professions["актер"],
					sqltools.NewSQLNullString(value.Actors[i].Name),
					weightActors)

				weightActors--
			}
		}

		countDirectors := len(value.Directors)
		if countDirectors > 0 {
			weightDirectors := countDirectors - 1

			for i := 0; i < countDirectors; i++ {
				values = append(values,
					f.mapPersons[value.Directors[i]],
					value.ID,
					f.guides.Professions["режиссер"],
					sqltools.NewSQLNullString(""),
					weightDirectors)

				weightDirectors--
			}
		}

		countInserts += countActors + countDirectors
	}

	countAttributes := strings.Count(insertFilmsPersons, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertFilmsPersons, countInserts, countAttributes)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkFilmPersonsReal")
	}

	return countInserts, nil
}

func (f *DBFiller) linkFilmTags() (int, error) {
	countInserts := 0

	values := make([]interface{}, 0)

	for _, value := range f.guides.Tags {
		count := pkg.RandMaxInt(f.Config.Volume.MaxFilmsInTag) + 1
		if count == 0 {
			continue
		}

		sequence := pkg.CryptoRandSequence(count+1, 1)

		for i := 0; i < count; i++ {
			values = append(values, sequence[i], value)
		}

		countInserts += count
	}

	countAttributes := strings.Count(insertFilmsTags, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertFilmsTags, countInserts, countAttributes)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkFilmTags")
	}

	return countInserts, nil
}

func (f *DBFiller) linkFilmImages() (int, error) {
	countInserts := 0

	for _, value := range f.films {
		countInserts += len(value.Images)
	}

	if countInserts == 0 {
		return 0, nil
	}

	countAttributes := strings.Count(insertFilmsImages, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertFilmsImages, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0

	for _, value := range f.persons {
		weight := len(value.Images)

		for _, image := range value.Images {
			values[pos] = value.ID
			pos++
			values[pos] = image
			weight--
			pos++
			values[pos] = weight
			pos++
		}
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkFilmImages")
	}

	return countInserts, nil
}

func (f *DBFiller) UpdateFilms() (int, error) {
	// Defining sending parameters
	query := updateFilmsSimple
	message := "UpdateFilms"

	globalCountInserts := len(f.filmsSQL)

	countAttributes := strings.Count(query, "$")

	action := func(values []interface{}) error {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
		defer cancelFunc()

		_, err := sqltools.InsertBatch(ctx, f.DB.Connection, query, values)
		if err != nil {
			return errors.Wrap(err, message)
		}

		return nil
	}

	values := make([]interface{}, countAttributes)

	for i := 0; i < len(f.films); i++ {
		pos := 0

		countScores := f.films[i].CountScores

		rating := sqltools.NewSQLNullFloat64(float32(math.Round(f.films[i].Rating/float64(countScores)*10*float64(1)) / 10 * float64(1)))

		values[pos] = rating
		pos++
		values[pos] = countScores
		pos++
		values[pos] = f.films[i].CountNegativeReviews
		pos++
		values[pos] = f.films[i].CountNeutralReviews
		pos++
		values[pos] = f.films[i].CountPositiveReviews
		pos++
		values[pos] = f.films[i].CountActors
		pos++
		values[pos] = f.films[i].ID
		pos++

		err := action(values)
		if err != nil {
			return 0, errors.Wrap(err, message)
		}
	}

	return globalCountInserts, nil
}

func (f *DBFiller) linkFilmGenresRandom() (int, error) {
	globalCountInserts := 0

	for _, value := range f.films {
		globalCountInserts += len(value.Genres)
	}

	// Defining sending parameters
	query := insertFilmsGenres
	message := "insertFilmsGenres"

	countAttributes := strings.Count(query, ",") + 1

	maxInsertValues := devpkg.MaxInsertValuesSQL / countAttributes

	action := func(countInserts int, values []interface{}) error {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
		defer cancelFunc()

		insertStatement := sqltools.CreateFullQuery(query, countInserts, countAttributes)

		_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
		if err != nil {
			return errors.Wrap(err, message)
		}

		return nil
	}

	pos := 0

	values := make([]interface{}, maxInsertValues*countAttributes)

	countInserts := 0

	for i := 0; i < len(f.films); i++ {
		sequence := pkg.CryptoRandSequence(len(f.guides.Genres)+1, 1)

		for j := 0; j < len(f.films[i].Genres); j++ {
			values[pos] = f.films[i].ID
			pos++
			values[pos] = sequence[j]
			pos++
			values[pos] = j
			pos++
		}

		countInserts += len(f.films[i].Genres)

		if devpkg.MaxInsertValuesSQL-pos < 20 || i == len(f.films)-1 {
			values = values[:pos]

			err := action(countInserts, values)
			if err != nil {
				return 0, errors.Wrap(err, message)
			}

			pos = 0

			countInserts = 0
		}
	}

	return globalCountInserts, nil
}

func (f *DBFiller) linkFilmCountriesRandom() (int, error) {
	globalCountInserts := 0

	for _, value := range f.films {
		globalCountInserts += len(value.ProdCountries)
	}

	// Defining sending parameters
	query := insertFilmsCountries
	message := "insertFilmsCountries"

	countAttributes := strings.Count(query, ",") + 1

	maxInsertValues := devpkg.MaxInsertValuesSQL / countAttributes

	action := func(countInserts int, values []interface{}) error {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
		defer cancelFunc()

		insertStatement := sqltools.CreateFullQuery(query, countInserts, countAttributes)

		_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
		if err != nil {
			return errors.Wrap(err, message)
		}

		return nil
	}

	pos := 0

	values := make([]interface{}, maxInsertValues*countAttributes)

	countInserts := 0

	for i := 0; i < len(f.films); i++ {
		sequence := pkg.CryptoRandSequence(len(f.guides.Countries)+1, 1)

		for j := 0; j < len(f.films[i].ProdCountries); j++ {
			values[pos] = f.films[i].ID
			pos++
			values[pos] = sequence[j]
			pos++
			values[pos] = j
			pos++
		}

		countInserts += len(f.films[i].ProdCountries)

		if devpkg.MaxInsertValuesSQL-pos < 20 || i == len(f.films)-1 {
			values = values[:pos]

			err := action(countInserts, values)
			if err != nil {
				return 0, errors.Wrap(err, message)
			}

			pos = 0

			countInserts = 0
		}
	}

	return globalCountInserts, nil
}

func (f *DBFiller) linkFilmCompaniesRandom() (int, error) {
	globalCountInserts := 0

	for _, value := range f.films {
		globalCountInserts += len(value.ProdCountries)
	}

	// Defining sending parameters
	query := insertFilmsCompanies
	message := "insertFilmsCompanies"

	countAttributes := strings.Count(query, ",") + 1

	maxInsertValues := devpkg.MaxInsertValuesSQL / countAttributes

	action := func(countInserts int, values []interface{}) error {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
		defer cancelFunc()

		insertStatement := sqltools.CreateFullQuery(query, countInserts, countAttributes)

		_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
		if err != nil {
			return errors.Wrap(err, message)
		}

		return nil
	}

	pos := 0

	values := make([]interface{}, maxInsertValues*countAttributes)

	countInserts := 0

	for i := 0; i < len(f.films); i++ {
		sequence := pkg.CryptoRandSequence(len(f.guides.Companies)+1, 1)

		for j := 0; j < len(f.films[i].ProdCompanies); j++ {
			values[pos] = f.films[i].ID
			pos++
			values[pos] = sequence[j]
			pos++
			values[pos] = j
			pos++
		}

		countInserts += len(f.films[i].ProdCompanies)

		if devpkg.MaxInsertValuesSQL-pos < 20 || i == len(f.films)-1 {
			values = values[:pos]

			err := action(countInserts, values)
			if err != nil {
				return 0, errors.Wrap(err, message)
			}

			pos = 0

			countInserts = 0
		}
	}

	return globalCountInserts, nil
}
