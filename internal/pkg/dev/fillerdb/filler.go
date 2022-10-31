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

func NewPostgreSQLRepository(url string) *DBSQL {
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

		generator: generatordatadb.NewDBGenerator(),
	}

	res.DB = NewPostgreSQLRepository(res.Config.Database.URL)

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

	count := 0

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

	f.createGuide(genres, f.genres)
	f.createGuide(countries, f.countries)
	f.createGuide(companies, f.companies)
	f.createGuide(professions, f.professions)
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
	count, err := f.UploadFilms()
	if err != nil {
		return err
	}
	logrus.Infof("%d films upload", count)

	count, err = f.UploadPersons()
	if err != nil {
		return err
	}
	logrus.Infof("%d persons upload", count)

	f.faceUsers = f.generator.GenerateUsers(f.Config.Volume.CountUser)
	count, err = f.UploadUsers()
	if err != nil {
		return err
	}
	logrus.Infof("%d face users upload", count)

	count, err = f.LinkUsersProfiles()
	if err != nil {
		return err
	}
	logrus.Infof("%d face users profiles link end", count)

	f.faceReviews = f.generator.GenerateReviews(f.Config.Volume.CountReviews, f.Config.Volume.MaxLengthReviewsBody)
	count, err = f.UploadReviews()
	if err != nil {
		return err
	}
	logrus.Infof("%d face reviews upload", count)

	return nil
}
