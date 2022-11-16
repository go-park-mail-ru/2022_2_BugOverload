package fillerdb

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/dev/generatordatadb"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type DBFiller struct {
	Config *Config

	DB *sqltools.Database

	films    []FilmFiller
	filmsSQL []FilmSQLFiller

	persons    []PersonFiller
	personsSQL []PersonSQLFiller

	collections    []CollectionFiller
	collectionsSQL []CollectionSQLFiller

	genres      map[string]int
	countries   map[string]int
	companies   map[string]int
	professions map[string]int
	tags        map[string]int

	generator *generatordatadb.DBGenerator

	faceUsers   []generatordatadb.UserFace
	faceReviews []generatordatadb.ReviewFace
}

func NewDBFiller(path string, config *Config) (*DBFiller, error) {
	res := &DBFiller{
		Config:      config,
		genres:      make(map[string]int),
		countries:   make(map[string]int),
		companies:   make(map[string]int),
		professions: make(map[string]int),
		tags:        make(map[string]int),

		generator: generatordatadb.NewDBGenerator(),
	}

	res.DB = sqltools.NewPostgresRepository()

	err := res.fillGuides(path)
	if err != nil {
		return nil, errors.Wrap(err, "NewDBFiller")
	}

	err = res.fillStorages(path)
	if err != nil {
		return nil, errors.Wrap(err, "NewDBFiller")
	}

	res.convertStructs()

	return res, nil
}

func (f *DBFiller) fillStorage(path string, someStorage interface{}) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("fillStorage: can't get data from file [%s] - [%w]", path, err)
	}

	err = json.Unmarshal(file, someStorage)
	if err != nil {
		return fmt.Errorf("fillStorage: can't Unmarshal data from file [%s] - [%w]", path, err)
	}

	return nil
}

func (f *DBFiller) createGuide(path string, someGuide map[string]int) error {
	stream, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("createGuide: err open file [%s] - [%w]", path, err)
	}

	defer stream.Close()

	scanner := bufio.NewScanner(stream)

	count := 1

	for scanner.Scan() {
		someGuide[scanner.Text()] = count
		count++
	}

	if err = scanner.Err(); err != nil {
		return fmt.Errorf("createGuide: err read from file [%s] - [%w]", path, err)
	}

	return nil
}

func (f *DBFiller) fillGuides(path string) error {
	genres := path + "/genres.txt"
	countries := path + "/countries.txt"
	companies := path + "/companies.txt"
	professions := path + "/professions.txt"
	tags := path + "/tags.txt"

	err := f.createGuide(genres, f.genres)
	if err != nil {
		return errors.Wrap(err, "fillGuides")
	}

	err = f.createGuide(countries, f.countries)
	if err != nil {
		return errors.Wrap(err, "fillGuides")
	}

	err = f.createGuide(companies, f.companies)
	if err != nil {
		return errors.Wrap(err, "fillGuides")
	}

	err = f.createGuide(professions, f.professions)
	if err != nil {
		return errors.Wrap(err, "fillGuides")
	}

	err = f.createGuide(tags, f.tags)
	if err != nil {
		return errors.Wrap(err, "fillGuides")
	}

	return nil
}

func (f *DBFiller) fillStorages(path string) error {
	films := path + "/films.json"
	persons := path + "/persons.json"
	collections := path + "/collections.json"

	err := f.fillStorage(films, &f.films)
	if err != nil {
		return errors.Wrap(err, "fillStorages")
	}

	err = f.fillStorage(persons, &f.persons)
	if err != nil {
		return errors.Wrap(err, "fillStorages")
	}

	err = f.fillStorage(collections, &f.collections)
	if err != nil {
		return errors.Wrap(err, "fillStorages")
	}

	return nil
}

func (f *DBFiller) convertStructs() {
	f.filmsSQL = make([]FilmSQLFiller, len(f.films))

	for idx, value := range f.films {
		f.filmsSQL[idx] = NewFilmSQLFillerOnFilm(value)
	}

	f.personsSQL = make([]PersonSQLFiller, len(f.persons))

	for idx, value := range f.persons {
		f.personsSQL[idx] = NewPersonSQLFillerOnPerson(value)
	}

	f.collectionsSQL = make([]CollectionSQLFiller, len(f.collections))

	for idx, value := range f.collections {
		f.collectionsSQL[idx] = NewCollectionSQLFilmOnCollection(value)
	}
}

func (f *DBFiller) Action() error {
	count, err := f.uploadFilms()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d films upload", count)

	count, err = f.uploadSerials()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d serials upload", count)

	count, err = f.linkFilmGenres()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d films genres link end", count)

	count, err = f.linkFilmCompanies()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d films companies link end", count)

	count, err = f.linkFilmCountries()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d films countries link end", count)

	count, err = f.linkFilmTags()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d film tags link end", count)

	count, err = f.linkFilmImages()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d films images link end", count)

	count, err = f.uploadPersons()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d persons upload", count)

	count, err = f.linkPersonProfession()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d persons professions link end", count)

	count, err = f.linkPersonImages()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d persons images link end", count)

	count, err = f.linkPersonGenres()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d persons genres link end", count)

	f.faceUsers = f.generator.GenerateUsers(f.Config.Volume.CountUser)
	count, err = f.uploadUsers()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d face users upload", count)

	count, err = f.linkProfileViews()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d face users profiles views link end", count)

	count, err = f.linkProfileRatings()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d face users profiles ratings link end", count)

	f.faceReviews = f.generator.GenerateReviews(f.Config.Volume.CountReviews, f.Config.Volume.MaxLengthReviewsBody)
	count, err = f.uploadReviews()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d face reviews upload", count)

	count, err = f.linkReviewsLikes()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d face reviews likes link end", count)

	count, err = f.linkFilmsReviews()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d film reviews link end", count)

	count, err = f.linkFilmPersons()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d film persons link end", count)

	count, err = f.uploadCollections()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d collections upload", count)

	count, err = f.linkCollectionProfile()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d collections profiles link end", count)

	count, err = f.UpdateFilms()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d films denormal fields updated", count)

	count, err = f.UpdatePersons()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d persons denormal fields updated", count)

	count, err = f.UpdateProfiles()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d profiles denormal fields updated", count)

	count, err = f.UpdateReviews()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d reviews denormal fields updated", count)

	return nil
}
