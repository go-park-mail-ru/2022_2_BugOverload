package fillerdb

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/dev/devpkg"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/dev/generatordatadb"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type Guides struct {
	Genres      map[string]int
	Countries   map[string]int
	Companies   map[string]int
	Professions map[string]int
	Tags        map[string]int
}

func NewGuides(path string) (*Guides, error) {
	res := &Guides{
		Genres:      make(map[string]int),
		Countries:   make(map[string]int),
		Companies:   make(map[string]int),
		Professions: make(map[string]int),
		Tags:        make(map[string]int),
	}

	err := res.fillGuides(path)
	if err != nil {
		return nil, errors.Wrap(err, "NewGuides")
	}

	return res, nil
}

func (g *Guides) createGuide(path string, someGuide map[string]int) error {
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

func (g *Guides) fillGuides(path string) error {
	genres := path + "/genres.txt"
	err := g.createGuide(genres, g.Genres)
	if err != nil {
		return errors.Wrap(err, "fillGuides")
	}

	countries := path + "/countries.txt"
	err = g.createGuide(countries, g.Countries)
	if err != nil {
		return errors.Wrap(err, "fillGuides")
	}

	companies := path + "/companies.txt"
	err = g.createGuide(companies, g.Companies)
	if err != nil {
		return errors.Wrap(err, "fillGuides")
	}

	professions := path + "/professions.txt"
	err = g.createGuide(professions, g.Professions)
	if err != nil {
		return errors.Wrap(err, "fillGuides")
	}

	tags := path + "/tags.txt"

	err = g.createGuide(tags, g.Tags)
	if err != nil {
		return errors.Wrap(err, "fillGuides")
	}

	return nil
}

type DBFiller struct {
	Config *devpkg.Config

	DB *sqltools.Database

	films    []devpkg.FilmFiller
	filmsSQL []devpkg.FilmSQLFiller

	persons    []devpkg.PersonFiller
	personsSQL []devpkg.PersonSQLFiller
	mapPersons map[string]int

	collections    []devpkg.CollectionFiller
	collectionsSQL []devpkg.CollectionSQLFiller

	guides *Guides

	generator *generatordatadb.DBGenerator

	Users   []generatordatadb.UserFace
	Reviews []generatordatadb.ReviewFace
}

func NewDBFiller(path string, config *devpkg.Config) (*DBFiller, error) {
	res := &DBFiller{
		Config:     config,
		mapPersons: make(map[string]int),
		generator:  generatordatadb.NewDBGenerator(),
	}

	res.DB = sqltools.NewPostgresRepository(&config.DatabaseParams)

	var err error

	res.guides, err = NewGuides(path)
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

func (f *DBFiller) fillStorages(path string) error {
	if f.Config.Volume.TypeFilms == devpkg.TypeDataReal {
		films := path + "/films.json"
		err := f.fillStorage(films, &f.films)
		if err != nil {
			return errors.Wrap(err, "fillStorages")
		}
	}

	if f.Config.Volume.TypeFilms == devpkg.TypeDataRandom {
		f.films = f.generator.GenerateFilms(&f.Config.Volume)
	}

	persons := path + "/persons.json"
	err := f.fillStorage(persons, &f.persons)
	if err != nil {
		return errors.Wrap(err, "fillStorages")
	}

	collections := path + "/collections.json"
	err = f.fillStorage(collections, &f.collections)
	if err != nil {
		return errors.Wrap(err, "fillStorages")
	}

	if f.Config.Volume.CountReviews > 0 {
		f.Reviews = f.generator.GenerateReviews(&f.Config.Volume)
	}

	return nil
}

func (f *DBFiller) convertStructs() {
	f.filmsSQL = make([]devpkg.FilmSQLFiller, len(f.films))

	for idx, value := range f.films {
		f.filmsSQL[idx] = devpkg.NewFilmSQLFillerOnFilm(value)
	}

	f.personsSQL = make([]devpkg.PersonSQLFiller, len(f.persons))

	for idx, value := range f.persons {
		f.personsSQL[idx] = devpkg.NewPersonSQLFillerOnPerson(value)
	}

	f.collectionsSQL = make([]devpkg.CollectionSQLFiller, len(f.collections))

	for idx, value := range f.collections {
		f.collectionsSQL[idx] = devpkg.NewCollectionSQLFilmOnCollection(value)
	}
}

func (f *DBFiller) Action() error {
	count, err := f.uploadFilms()
	if err != nil {
		return errors.Wrap(err, "Action")
	}

	if f.Config.Volume.TypeFilms == devpkg.TypeDataReal {
		logrus.Infof("%d real films upload", count)
	} else if f.Config.Volume.TypeFilms == devpkg.TypeDataRandom {
		logrus.Infof("%d random films upload", count)
	}

	count, err = f.uploadFilmsMedia()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d films media upload", count)

	count, err = f.uploadSerials()
	if err != nil {
		return errors.Wrap(err, "Action")
	}
	logrus.Infof("%d serials upload", count)

	if f.Config.Volume.TypeFilms == devpkg.TypeDataReal {
		count, err = f.linkFilmGenresReal()
		if err != nil {
			return errors.Wrap(err, "Action")
		}
		logrus.Infof("%d real films genres link end", count)

		count, err = f.linkFilmCompaniesReal()
		if err != nil {
			return errors.Wrap(err, "Action")
		}
		logrus.Infof("%d real films companies link end", count)

		count, err = f.linkFilmCountriesReal()
		if err != nil {
			return errors.Wrap(err, "Action")
		}
		logrus.Infof("%d real films countries link end", count)
	} else if f.Config.Volume.TypeFilms == devpkg.TypeDataRandom {
		count, err = f.linkFilmGenresRandom()
		if err != nil {
			return errors.Wrap(err, "Action")
		}
		logrus.Infof("%d random films genres link end", count)

		count, err = f.linkFilmCompaniesRandom()
		if err != nil {
			return errors.Wrap(err, "Action")
		}
		logrus.Infof("%d random films companies link end", count)

		count, err = f.linkFilmCountriesRandom()
		if err != nil {
			return errors.Wrap(err, "Action")
		}
		logrus.Infof("%d random films countries link end", count)
	}

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

	f.Users = f.generator.GenerateUsers(f.Config.Volume.CountUser)
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

	if f.Config.Volume.CountReviews > 0 {
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
	}

	if f.Config.Volume.TypeFilmsPersonLinks == devpkg.TypeDataRandom {
		count, err = f.linkFilmPersonsRandom()
		if err != nil {
			return errors.Wrap(err, "Action")
		}
		logrus.Infof("%d random film persons link end", count)
	} else if f.Config.Volume.TypeFilmsPersonLinks == devpkg.TypeDataReal {
		count, err = f.linkFilmPersonsReal()
		if err != nil {
			return errors.Wrap(err, "Action")
		}
		logrus.Infof("%d real film persons link end", count)
	}

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

	if f.Config.Volume.CountReviews > 0 {
		count, err = f.UpdateReviews()
		if err != nil {
			return errors.Wrap(err, "Action")
		}
		logrus.Infof("%d reviews denormal fields updated", count)
	}

	if f.Config.Volume.TypeFilms == devpkg.TypeDataReal {
		count, err = f.UpdateFilms()
		if err != nil {
			return errors.Wrap(err, "Action")
		}
		logrus.Infof("%d films denormal fields updated", count)
	}

	return nil
}
