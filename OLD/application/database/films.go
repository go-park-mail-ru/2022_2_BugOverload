package database

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"sync"

	"go-park-mail-ru/2022_2_BugOverload/OLD/application/errors"
)

// FilmStorage is TMP impl database for films, where key = film_id
type FilmStorage struct {
	storage map[uint]models.Film
	mu      *sync.Mutex
}

// NewFilmStorage is constructor for FilmStorage
func NewFilmStorage() *FilmStorage {
	return &FilmStorage{
		storage: map[uint]models.Film{
			0: {
				ID:        0,
				Name:      "Дюна",
				YearProd:  "2021",
				Rating:    "9.9",
				PosterVer: "asserts/img/posters/dune_poster.jpg",
				Genres:    []string{"Фэнтези", "приключения"},
			},
			1: {
				ID:        1,
				Name:      "Убить Билла",
				YearProd:  "2021",
				Rating:    "9.0",
				PosterVer: "asserts/img/posters/8.png",
				Genres:    []string{"Боевик", "приключения"},
			},
			2: {
				ID:        2,
				Name:      "Головокружение",
				YearProd:  "2021",
				Rating:    "7.1",
				PosterVer: "asserts/img/posters/9.png",
				Genres:    []string{"Триллер", "приключения"},
			},
			3: {
				ID:        3,
				Name:      "Петровы в гриппе",
				YearProd:  "2020",
				Rating:    "6.8",
				PosterVer: "asserts/img/posters/30.png",
				Genres:    []string{"Драма", "фантастика"},
			},
			4: {
				ID:        4,
				Name:      "Она",
				YearProd:  "2013",
				Rating:    "7.6",
				PosterVer: "asserts/img/posters/31.png",
				Genres:    []string{"Мелодрама", "фантастика"},
			},
			5: {
				ID:        5,
				Name:      "Плата за страх",
				YearProd:  "1952",
				Rating:    "8.0",
				PosterVer: "asserts/img/posters/32.png",
				Genres:    []string{"Мелодрама", "фантастика"},
			},
			6: {
				ID:        6,
				Name:      "Черное зеркало",
				YearProd:  "2011-2019",
				Rating:    "8.5",
				PosterVer: "asserts/img/posters/33.png",
				Genres:    []string{"Фантастика", "драма"},
			},
			7: {
				ID:        7,
				Name:      "Олдбой",
				YearProd:  "2003",
				Rating:    "8.1",
				PosterVer: "asserts/img/posters/34.png",
				Genres:    []string{"Триллер", "детектив"},
			},
			8: {
				ID:        8,
				Name:      "Доказательство смерти",
				YearProd:  "2021",
				Rating:    "3.3",
				PosterVer: "asserts/img/posters/5.png",
				Genres:    []string{"Фэнтези", "приключения"},
			},
			9: {
				ID:        9,
				Name:      "Чунгингский экспресс",
				YearProd:  "2021",
				Rating:    "7.1",
				PosterVer: "asserts/img/posters/7.png",
				Genres:    []string{"Фэнтези", "приключения"},
			},
			10: {
				ID:        10,
				Name:      "Девушка с татуировкой дракона",
				YearProd:  "2021",
				Rating:    "5.7",
				PosterVer: "asserts/img/posters/6.png",
				Genres:    []string{"Триллер", "приключения"},
			},
			11: {
				ID:        11,
				Name:      "Человек",
				YearProd:  "2021",
				Rating:    "7.1",
				PosterVer: "asserts/img/posters/1.png",
				Genres:    []string{"Документальный", "драма"},
			},
			12: {
				ID:        12,
				Name:      "Люси",
				YearProd:  "2021",
				Rating:    "8.9",
				PosterVer: "asserts/img/posters/2.png",
				Genres:    []string{"Фэнтези", "приключения"},
			},
			13: {
				ID:        13,
				Name:      "Властелин колец. Братство кольца",
				YearProd:  "2021",
				Rating:    "8.4",
				PosterVer: "asserts/img/posters/3.png",
				Genres:    []string{"Фэнтези", "приключения"},
			},
			14: {
				ID:        14,
				Name:      "Дом, который построил Джек",
				YearProd:  "2021",
				Rating:    "7.2",
				PosterVer: "asserts/img/posters/4.png",
				Genres:    []string{"Фэнтези", "приключения"},
			},
			15: {
				ID:        15,
				Name:      "Интерстеллар",
				YearProd:  "2014",
				Rating:    "8.6",
				PosterVer: "asserts/img/posters/10.png",
				Genres:    []string{"Фантастика", "драма"},
			},
			16: {
				ID:        16,
				Name:      "1+1",
				YearProd:  "2011",
				Rating:    "8.8",
				PosterVer: "asserts/img/posters/11.png",
				Genres:    []string{"Комедия", "драма"},
			},
			17: {
				ID:        17,
				Name:      "Темный рыцарь",
				YearProd:  "2008",
				Rating:    "8.5",
				PosterVer: "asserts/img/posters/12.png",
				Genres:    []string{"Фантастика", "боевик"},
			},
			18: {
				ID:        18,
				Name:      "Бойцовский клуб",
				YearProd:  "1999",
				Rating:    "8.6",
				PosterVer: "asserts/img/posters/13.png",
				Genres:    []string{"Триллер", "драма"},
			},
			19: {
				ID:        19,
				Name:      "Титаник",
				YearProd:  "1997",
				Rating:    "8.6",
				PosterVer: "asserts/img/posters/14.png",
				Genres:    []string{"Мелодрама", "история"},
			},
			20: {
				ID:        20,
				Name:      "Ходячий замок",
				YearProd:  "2004",
				Rating:    "8.3",
				PosterVer: "asserts/img/posters/15.png",
				Genres:    []string{"Мультфильм", "аниме"},
			},
			21: {
				ID:        21,
				Name:      "Волк с Уолл-стрит",
				YearProd:  "2013",
				Rating:    "7.9",
				PosterVer: "asserts/img/posters/16.png",
				Genres:    []string{"Криминал", "комедия"},
			},
			22: {
				ID:        22,
				Name:      "Три тысячи лет желаний",
				YearProd:  "2022",
				Rating:    "7.0",
				PosterVer: "asserts/img/posters/17.png",
				Genres:    []string{"Фэнтези", "драма"},
			},
			23: {
				ID:        23,
				Name:      "Джиперс Криперс: Возрожденный",
				YearProd:  "2022",
				Rating:    "3.5",
				PosterVer: "asserts/img/posters/18.png",
				Genres:    []string{"Ужасы"},
			},
			24: {
				ID:        24,
				Name:      "Эра выживания ",
				YearProd:  "2022",
				Rating:    "5.5",
				PosterVer: "asserts/img/posters/19.png",
				Genres:    []string{"Фантастика"},
			},
			25: {
				ID:        25,
				Name:      "Проклятие Плачущей. Возвращение",
				YearProd:  "2022",
				Rating:    "4.2",
				PosterVer: "asserts/img/posters/20.png",
				Genres:    []string{"Ужасы"},
			},
			26: {
				ID:               26,
				Name:             "Звёздные войны. Эпизод IV: Новая надежда",
				ShortDescription: "Галактическая империя порабощает целые системы, повстанцы будут мешать этому",
				YearProd:         "2021",
				PosterHor:        "asserts/img/StarWars.jpeg",
				Genres:           []string{"фэнтези", "приключения"},
			},
			27: {
				ID:               27,
				Name:             "Дюна",
				ShortDescription: "Наследник дома Атрейдесов отправляется на одну из самых опасных планет во Вселенной — Арракис",
				YearProd:         "2021",
				PosterHor:        "asserts/img/dune.jpg",
				Genres:           []string{"фэнтези", "приключения"},
			},
			28: {
				ID:               28,
				Name:             "Джокер",
				ShortDescription: "Пытаясь нести в мир хорошее и дарить людям радость, Артур сталкивается с человеческой жестокостью",
				YearProd:         "2021",
				PosterHor:        "asserts/img/joker_hor.jpg",
				Genres:           []string{"драма"},
			},
			29: {
				ID:               29,
				Name:             "2001: Космическая одисея",
				ShortDescription: "Экипаж космического корабля «Дискавери» должен исследовать район галактики",
				YearProd:         "1968",
				PosterHor:        "asserts/img/space_odyssey_hor.jpg",
				Genres:           []string{"фэнтези", "приключения"},
			},
		},
		mu: &sync.Mutex{},
	}
}

// CheckExist is method to check the existence of such a film in the database
func (fs *FilmStorage) CheckExist(filmID uint) bool {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	_, ok := fs.storage[filmID]
	return ok
}

// AddFilm is method for creating a film in database
func (fs *FilmStorage) AddFilm(f models.Film) {
	if !fs.CheckExist(f.ID) {
		fs.mu.Lock()
		defer fs.mu.Unlock()

		fs.storage[f.ID] = f
	}
}

// GetFilm return film using film_id (primary key)
func (fs *FilmStorage) GetFilm(filmID uint) (models.Film, error) {
	if !fs.CheckExist(filmID) {
		return models.Film{}, errors.ErrFilmNotFound
	}

	fs.mu.Lock()
	defer fs.mu.Unlock()

	return fs.storage[filmID], nil
}

// GetStorageLen return films count in storage
func (fs *FilmStorage) GetStorageLen() int {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	return len(fs.storage)
}

// ClearStorage will delete all films from storage
func (fs *FilmStorage) ClearStorage() {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	if len(fs.storage) != 0 {
		fs.storage = make(map[uint]models.Film)
	}
}
