package constparams

import "database/sql"

// DB
const (
	// PersonProfessions
	Actor    = 1
	Artist   = 7
	Director = 2
	Writer   = 3
	Producer = 4
	Operator = 5
	Montage  = 8
	Composer = 6

	DefTypeFilm   = "film"
	DefTypeSerial = "serial"

	TypeReviewPositive = "positive"
	TypeReviewNegative = "negative"
	TypeReviewNeutral  = "neutral"

	DefGender = "male"
	OnlyDate  = "2006"
)

// Tags from front
const (
	PopularFrom  = "popular"
	InCinemaFrom = "in_cinema"
)

// Tags in db
const (
	PopularIn  = "популярное"
	InCinemaIn = "сейчас в кино"
)

func NewTagsMap() map[string]string {
	res := make(map[string]string, 0)

	res[PopularFrom] = PopularIn
	res[InCinemaFrom] = InCinemaIn

	return res
}

// Genres from front
const (
	ComedyFrom      = "comedy"
	AnimeFrom       = "anime"
	BiographyFrom   = "biography"
	ActionFrom      = "action"
	WesternFrom     = "western"
	WarFrom         = "war"
	DetectiveFrom   = "detective"
	ChildrenFrom    = "children"
	DocumentaryFrom = "documentary"
	DramaFrom       = "drama"
	TheGameFrom     = "the_game"
	HistoryFrom     = "history"
	ConcertFrom     = "concert"
	ShortFilmFrom   = "short_film"
	CrimeFrom       = "crime"
	MelodramaFrom   = "melodrama"
	MusicFrom       = "music"
	AnimationFrom   = "animation"
	MusicalFrom     = "musical"
	AdventuresFrom  = "adventures"
	FamilyFrom      = "family"
	SportsFrom      = "sports"
	TalkShowFrom    = "talk_show"
	ThrillerFrom    = "thriller"
	HorrorFrom      = "horror"
	SciFiFrom       = "sci-fi"
	NoirFrom        = "noir"
	FantasyFrom     = "fantasy"
	CeremonyFrom    = "ceremony"
	NewsFrom        = "news"
	SeriesFrom      = "series"
	RealityTVFrom   = "reality_TV"
)

// Genres in db
const (
	ComedyIn      = "комедия"
	AnimeIn       = "аниме"
	BiographyIn   = "биография"
	ActionIn      = "боевик"
	WesternIn     = "вестерн"
	WarIn         = "военный"
	DetectiveIn   = "детектив"
	ChildrenIn    = "детский"
	DocumentaryIn = "документальный"
	DramaIn       = "драма"
	TheGameIn     = "игра"
	HistoryIn     = "история"
	ConcertIn     = "концерт"
	ShortFilmIn   = "короткометражка"
	CrimeIn       = "криминал"
	MelodramaIn   = "мелодрама"
	MusicIn       = "музыка"
	AnimationIn   = "мультфильм"
	MusicalIn     = "мюзикл"
	AdventuresIn  = "приключения"
	FamilyIn      = "семейный"
	SportsIn      = "спорт"
	TalkShowIn    = "ток-шоу"
	ThrillerIn    = "триллер"
	HorrorIn      = "ужасы"
	SciFiIn       = "фантастика"
	NoirIn        = "фильм-нуар"
	FantasyIn     = "фэнтези"
	CeremonyIn    = "церемония"
	NewsIn        = "новости"
	SeriesIn      = "сериал"
	RealityTVIn   = "реальное ТВ"
)

func NewGenresMap() map[string]string {
	res := make(map[string]string, 0)

	res[ComedyFrom] = ComedyIn
	res[AnimeFrom] = AnimeIn
	res[BiographyFrom] = BiographyIn
	res[ActionFrom] = ActionIn
	res[WesternFrom] = WesternIn
	res[WarFrom] = WarIn
	res[DetectiveFrom] = DetectiveIn
	res[ChildrenFrom] = ChildrenIn
	res[DocumentaryFrom] = DocumentaryIn
	res[DramaFrom] = DramaIn
	res[TheGameFrom] = TheGameIn
	res[HistoryFrom] = HistoryIn
	res[ConcertFrom] = ConcertIn
	res[ShortFilmFrom] = ShortFilmIn
	res[CrimeFrom] = CrimeIn
	res[MelodramaFrom] = MelodramaIn
	res[MusicFrom] = MusicIn
	res[AnimationFrom] = AnimationIn
	res[MusicalFrom] = MusicalIn
	res[AdventuresFrom] = AdventuresIn
	res[FamilyFrom] = FamilyIn
	res[SportsFrom] = SportsIn
	res[TalkShowFrom] = TalkShowIn
	res[ThrillerFrom] = ThrillerIn
	res[HorrorFrom] = HorrorIn
	res[SciFiFrom] = SciFiIn
	res[NoirFrom] = NoirIn
	res[FantasyFrom] = FantasyIn
	res[CeremonyFrom] = CeremonyIn
	res[NewsFrom] = NewsIn
	res[SeriesFrom] = SeriesIn
	res[RealityTVFrom] = RealityTVIn

	return res
}

var GenresMap = NewGenresMap()

var TagsMap = NewTagsMap()

// TxDefaultOptions for Postgres
var TxDefaultOptions = &sql.TxOptions{
	Isolation: sql.LevelDefault,
	ReadOnly:  true,
}

// TxInsertOptions for Postgres
var TxInsertOptions = &sql.TxOptions{
	Isolation: sql.LevelDefault,
	ReadOnly:  false,
}
