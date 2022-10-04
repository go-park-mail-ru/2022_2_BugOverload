package database

import (
	"sync"

	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// FilmStorage is TMP impl database for films, where key = film_id
type FilmStorage struct {
	storage map[uint]structs.Film
	mu      *sync.Mutex
}

// NewFilmStorage is constructor for FilmStorage
func NewFilmStorage() *FilmStorage {
	return &FilmStorage{
		make(map[uint]structs.Film),
		&sync.Mutex{},
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
func (fs *FilmStorage) AddFilm(f structs.Film) {
	if !fs.CheckExist(f.ID) {
		fs.mu.Lock()
		defer fs.mu.Unlock()

		fs.storage[f.ID] = f
	}
}

// GetFilm return film using film_id (primary key)
func (fs *FilmStorage) GetFilm(filmID uint) (structs.Film, error) {
	if !fs.CheckExist(filmID) {
		return structs.Film{}, errorshandlers.ErrFilmNotFound
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

// FillFilmStorage is temporary function, filling local storage
func (fs *FilmStorage) FillFilmStorage() {
	// First collection
	var currentID uint
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Дюна",
		YearProd:  "2021",
		Rating:    "9.9",
		PosterVer: "asserts/img/posters/dune_poster.jpg",
		Genres:    []string{"Фэнтези", "приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Убить Билла",
		YearProd:  "2021",
		Rating:    "9.0",
		PosterVer: "asserts/img/posters/8.png",
		Genres:    []string{"Боевик", "приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Головокружение",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/9.png",
		Genres:    []string{"Триллер", "приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Петровы в гриппе",
		YearProd:  "2020",
		Rating:    "6.8",
		PosterVer: "asserts/img/posters/30.png",
		Genres:    []string{"Драма", "фантастика"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Она",
		YearProd:  "2013",
		Rating:    "7.6",
		PosterVer: "asserts/img/posters/31.png",
		Genres:    []string{"Мелодрама", "фантастика"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Плата за страх",
		YearProd:  "1952",
		Rating:    "8.0",
		PosterVer: "asserts/img/posters/32.png",
		Genres:    []string{"Мелодрама", "фантастика"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Черное зеркало",
		YearProd:  "2011-2019",
		Rating:    "8.5",
		PosterVer: "asserts/img/posters/33.png",
		Genres:    []string{"Фантастика", "драма"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Олдбой",
		YearProd:  "2003",
		Rating:    "8.1",
		PosterVer: "asserts/img/posters/34.png",
		Genres:    []string{"Триллер", "детектив"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Доказательство смерти",
		YearProd:  "2021",
		Rating:    "3.3",
		PosterVer: "asserts/img/posters/5.png",
		Genres:    []string{"Фэнтези", "приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Чунгингский экспресс",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/7.png",
		Genres:    []string{"Фэнтези", "приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Девушка с татуировкой дракона",
		YearProd:  "2021",
		Rating:    "5.7",
		PosterVer: "asserts/img/posters/6.png",
		Genres:    []string{"Триллер", "приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Человек",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/1.png",
		Genres:    []string{"Документальный", "драма"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Люси",
		YearProd:  "2021",
		Rating:    "8.9",
		PosterVer: "asserts/img/posters/2.png",
		Genres:    []string{"Фэнтези", "приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Властелин колец. Братство кольца",
		YearProd:  "2021",
		Rating:    "8.4",
		PosterVer: "asserts/img/posters/3.png",
		Genres:    []string{"Фэнтези", "приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Дом, который построил Джек",
		YearProd:  "2021",
		Rating:    "7.2",
		PosterVer: "asserts/img/posters/4.png",
		Genres:    []string{"Фэнтези", "приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Интерстеллар",
		YearProd:  "2014",
		Rating:    "8.6",
		PosterVer: "asserts/img/posters/10.png",
		Genres:    []string{"Фантастика", "драма"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "1+1",
		YearProd:  "2011",
		Rating:    "8.8",
		PosterVer: "asserts/img/posters/11.png",
		Genres:    []string{"Комедия", "драма"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Темный рыцарь",
		YearProd:  "2008",
		Rating:    "8.5",
		PosterVer: "asserts/img/posters/12.png",
		Genres:    []string{"Фантастика", "боевик"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Бойцовский клуб",
		YearProd:  "1999",
		Rating:    "8.6",
		PosterVer: "asserts/img/posters/13.png",
		Genres:    []string{"Триллер", "драма"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Титаник",
		YearProd:  "1997",
		Rating:    "8.6",
		PosterVer: "asserts/img/posters/14.png",
		Genres:    []string{"Мелодрама", "история"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Ходячий замок",
		YearProd:  "2004",
		Rating:    "8.3",
		PosterVer: "asserts/img/posters/15.png",
		Genres:    []string{"Мультфильм", "аниме"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Волк с Уолл-стрит",
		YearProd:  "2013",
		Rating:    "7.9",
		PosterVer: "asserts/img/posters/16.png",
		Genres:    []string{"Криминал", "комедия"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Волк с Уолл-стрит",
		YearProd:  "2013",
		Rating:    "7.9",
		PosterVer: "asserts/img/posters/16.png",
		Genres:    []string{"Криминал", "комедия"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Три тысячи лет желаний",
		YearProd:  "2022",
		Rating:    "7.0",
		PosterVer: "asserts/img/posters/17.png",
		Genres:    []string{"Фэнтези", "драма"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Джиперс Криперс: Возрожденный",
		YearProd:  "2022",
		Rating:    "3.5",
		PosterVer: "asserts/img/posters/18.png",
		Genres:    []string{"Ужасы"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Эра выживания ",
		YearProd:  "2022",
		Rating:    "5.5",
		PosterVer: "asserts/img/posters/19.png",
		Genres:    []string{"фантастика"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Проклятие Плачущей. Возвращение",
		YearProd:  "2022",
		Rating:    "4.2",
		PosterVer: "asserts/img/posters/20.png",
		Genres:    []string{"Ужасы"},
	})
	currentID++
}

// FillFilmStorageSpecial is temporary function, filling local storage
func (fs *FilmStorage) FillFilmStorageSpecial() {
	var currentID uint = uint(len(fs.storage))
	// Third collection (poster)
	fs.AddFilm(structs.Film{
		ID:               currentID,
		Name:             "Звёздные войны. Эпизод IV: Новая надежда",
		ShortDescription: "Галактическая империя порабощает целые системы, повстанцы будут мешать этому",
		YearProd:         "2021",
		PosterHor:        "asserts/img/StarWars.jpeg",
		Genres:           []string{"фэнтези", "приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:               currentID,
		Name:             "Дюна",
		ShortDescription: "Наследник дома Атрейдесов отправляется на одну из самых опасных планет во Вселенной — Арракис",
		YearProd:         "2021",
		PosterHor:        "asserts/img/dune.jpg",
		Genres:           []string{"фэнтези", "приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:               currentID,
		Name:             "Джокер",
		ShortDescription: "Пытаясь нести в мир хорошее и дарить людям радость, Артур сталкивается с человеческой жестокостью",
		YearProd:         "2021",
		PosterHor:        "asserts/img/joker_hor.jpg",
		Genres:           []string{"драма"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:               currentID,
		Name:             "2001: Космическая одисея",
		ShortDescription: "Экипаж космического корабля «Дискавери» должен исследовать район галактики",
		YearProd:         "1968",
		PosterHor:        "asserts/img/space_odyssey_hor.jpg",
		Genres:           []string{"фэнтези", "приключения"},
	})
}
