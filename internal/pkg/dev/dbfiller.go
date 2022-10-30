package dev

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"log"
	"os"

	// justifying it
	_ "github.com/jackc/pgx/stdlib"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
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

	films   []models.Film
	persons []models.Person

	genres      map[string]int
	countries   map[string]int
	companies   map[string]int
	professions map[string]int
}

func NewDBFiller(path string, config *Config) *DBFiller {
	res := &DBFiller{
		Config:      config,
		genres:      make(map[string]int),
		countries:   make(map[string]int),
		companies:   make(map[string]int),
		professions: make(map[string]int),
	}

	res.DB = NewPostgreSQLRepository(res.Config.Database.URL)

	res.fillGuides(path)
	res.fillStorages(path)

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

func (f *DBFiller) Action() error {
	count, err := f.UploadFilms()
	if err != nil {
		return err
	}

	logrus.Infof("%d films created", count)

	return nil
}
