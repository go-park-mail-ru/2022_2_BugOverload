package fillerdb

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"log"
	"os"

	// justifying it
	_ "github.com/jackc/pgx/stdlib"
	"github.com/sirupsen/logrus"

	modelsFilmRepo "go-park-mail-ru/2022_2_BugOverload/internal/film/repository/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	modelsPersonRepo "go-park-mail-ru/2022_2_BugOverload/internal/person/repository/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/dev/generatordatadb"
)

type DBSQL struct {
	Connection *sql.DB
}

func NewPostgreSQLRepository() *DBSQL {
	url := "user=" + os.Getenv("POSTGRES_USER") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" port=" + os.Getenv("POSTGRES_PORT") +
		" sslmode=" + os.Getenv("POSTGRES_SSLMODE")

	connection, err := sql.Open("pgx", url)
	if err != nil {
		log.Fatalln("Can't parse config", err)
	}

	err = connection.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return &DBSQL{Connection: connection}
}

type DBFiller struct {
	Config *Config

	DB *DBSQL

	films    []models.Film
	filmsSQL []modelsFilmRepo.FilmSQL

	persons    []models.Person
	personsSQL []modelsPersonRepo.PersonSQL

	genres      map[string]int
	countries   map[string]int
	companies   map[string]int
	professions map[string]int
	tags        map[string]int

	generator *generatordatadb.DBGenerator

	faceUsers   []generatordatadb.UserFace
	faceReviews []generatordatadb.ReviewFace
}

func NewDBFiller(path string, config *Config) *DBFiller {
	res := &DBFiller{
		Config:      config,
		genres:      make(map[string]int),
		countries:   make(map[string]int),
		companies:   make(map[string]int),
		professions: make(map[string]int),
		tags:        make(map[string]int),

		generator: generatordatadb.NewDBGenerator(),
	}

	res.DB = NewPostgreSQLRepository()

	res.fillGuides(path)
	res.fillStorages(path)

	res.convertStructs()

	return res
}

func (f *DBFiller) fillStorage(path string, someStorage interface{}) {
	file, err := os.ReadFile(path)
	if err != nil {
		logrus.Error("FillStorage: can't get data from file ", err, path)
	}

	err = json.Unmarshal(file, someStorage)
	if err != nil {
		logrus.Error("FillStorage: can't Unmarshal data from file ", err, path)
	}
}

func (f *DBFiller) createGuide(path string, someGuide map[string]int) {
	stream, err := os.Open(path)
	if err != nil {
		logrus.Fatal(path, err)
	}

	defer stream.Close()

	scanner := bufio.NewScanner(stream)

	count := 1

	for scanner.Scan() {
		someGuide[scanner.Text()] = count
		count++
	}

	if err = scanner.Err(); err != nil {
		logrus.Fatal(path, err)
	}
}

func (f *DBFiller) fillGuides(path string) {
	genres := path + "/genres.txt"
	countries := path + "/countries.txt"
	companies := path + "/companies.txt"
	professions := path + "/professions.txt"
	tags := path + "/tags.txt"

	f.createGuide(genres, f.genres)
	f.createGuide(countries, f.countries)
	f.createGuide(companies, f.companies)
	f.createGuide(professions, f.professions)
	f.createGuide(tags, f.tags)
}

func (f *DBFiller) fillStorages(path string) {
	films := path + "/films.json"
	persons := path + "/persons.json"

	f.fillStorage(films, &f.films)
	f.fillStorage(persons, &f.persons)
}

func (f *DBFiller) convertStructs() {
	f.filmsSQL = make([]modelsFilmRepo.FilmSQL, len(f.films))

	for idx, value := range f.films {
		f.filmsSQL[idx] = modelsFilmRepo.NewFilmSQL(value)
	}

	f.personsSQL = make([]modelsPersonRepo.PersonSQL, len(f.persons))

	for idx, value := range f.persons {
		f.personsSQL[idx] = modelsPersonRepo.NewPersonSQL(value)
	}
}

func (f *DBFiller) Action() error {
	count, err := f.uploadFilms()
	if err != nil {
		return err
	}
	logrus.Infof("%d films upload", count)

	count, err = f.linkFilmGenres()
	if err != nil {
		return err
	}
	logrus.Infof("%d films genres link end", count)

	count, err = f.linkFilmCompanies()
	if err != nil {
		return err
	}
	logrus.Infof("%d films companies link end", count)

	count, err = f.linkFilmCountries()
	if err != nil {
		return err
	}
	logrus.Infof("%d films countries link end", count)

	count, err = f.linkFilmTags()
	if err != nil {
		return err
	}
	logrus.Infof("%d film tags link end", count)

	count, err = f.uploadPersons()
	if err != nil {
		return err
	}
	logrus.Infof("%d persons upload", count)

	count, err = f.linkPersonProfession()
	if err != nil {
		return err
	}
	logrus.Infof("%d persons professions link end", count)

	count, err = f.linkPersonGenres()
	if err != nil {
		return err
	}
	logrus.Infof("%d persons genres link end", count)

	f.faceUsers = f.generator.GenerateUsers(f.Config.Volume.CountUser)
	count, err = f.uploadUsers()
	if err != nil {
		return err
	}
	logrus.Infof("%d face users upload", count)

	count, err = f.linkUsersProfiles()
	if err != nil {
		return err
	}
	logrus.Infof("%d face users profiles link end", count)

	count, err = f.linkProfileViews()
	if err != nil {
		return err
	}
	logrus.Infof("%d face users profiles views link end", count)

	count, err = f.linkProfileRatings()
	if err != nil {
		return err
	}
	logrus.Infof("%d face users profiles ratings link end", count)

	f.faceReviews = f.generator.GenerateReviews(f.Config.Volume.CountReviews, f.Config.Volume.MaxLengthReviewsBody)
	count, err = f.uploadReviews()
	if err != nil {
		return err
	}
	logrus.Infof("%d face reviews upload", count)

	count, err = f.linkReviewsLikes()
	if err != nil {
		return err
	}
	logrus.Infof("%d face reviews likes link end", count)

	count, err = f.linkFilmsReviews()
	if err != nil {
		return err
	}
	logrus.Infof("%d film reviews link end", count)

	count, err = f.linkFilmPersons()
	if err != nil {
		return err
	}
	logrus.Infof("%d film persons link end", count)

	return nil
}
