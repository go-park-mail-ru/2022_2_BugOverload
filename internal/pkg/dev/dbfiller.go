package dev

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	//   justifying it
	_ "github.com/jackc/pgx/stdlib"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

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

func (f *DBFiller) Action() {
	f.UploadFilms()

	logrus.Info("SUCCESS")
}

const countAttributesFilms = 6

func (f *DBFiller) UploadFilms() {
	var values []interface{}

	placeholders := CreatePlaceholders(countAttributesFilms, len(f.films))

	for _, value := range f.films {
		values = append(values,
			value.Name,
			value.ProdYear,
			value.PosterVer,
			value.PosterHor,
			value.Description,
			value.ShortDescription)
	}

	query := "INSERT INTO films(name, prod_year, poster_ver, poster_hor, description, short_description) VALUES"

	insertStatement := fmt.Sprintf("%s %s", query, placeholders)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout))
	defer cancelFunc()

	stmt, err := f.DB.Connection.PrepareContext(ctx, insertStatement)
	if err != nil {
		logrus.Fatalf("Error %s when preparing SQL statement", err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, values...)
	if err != nil {
		logrus.Fatalf("Error %s when inserting row into films table", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		logrus.Fatalf("Error %s when finding rows affected", err)
	}

	logrus.Infof("%d films created", rows)
}

func CreatePlaceholders(countAttributes int, countValues int) string {
	values := make([]string, countAttributes*countValues)

	for i := 0; i < countAttributes*countValues; i++ {
		values[i] = fmt.Sprintf("$%d", i+1)
	}

	valuesRow := make([]string, countValues)

	for i := 0; i < countValues; i++ {
		valuesRow[i] = "(" + strings.Join(values[i*countAttributes:countAttributes*(i+1)], ",") + ")"
	}

	return strings.Join(valuesRow, ",\n")
}
