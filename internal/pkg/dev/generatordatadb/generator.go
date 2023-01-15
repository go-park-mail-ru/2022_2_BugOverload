package generatordatadb

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"log"
	"strings"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/dev/devpkg"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

type DBGenerator struct{}

func NewDBGenerator() *DBGenerator {
	return &DBGenerator{}
}

func (g *DBGenerator) GenerateUsers(count int) []UserFace {
	res := make([]UserFace, count)

	for idx := range res {
		err := faker.FakeData(&res[idx], options.WithFieldsToIgnore("ID"))
		if err != nil {
			log.Println(err)
		}
	}

	return res
}

func strLength(number int) uint {
	return uint(pkg.RandIntInInterval(number+1, number+1/del))
}

const del = 3

func (g *DBGenerator) GenerateReviews(volume *devpkg.Volume) []ReviewFace {
	res := make([]ReviewFace, volume.CountReviews)

	typesReview := []string{"positive", "neutral", "negative"}

	for idx := range res {
		err := faker.FakeData(&res[idx],
			options.WithRandomStringLength(strLength(volume.MaxLengthReviewsBody)),
			options.WithFieldsToIgnore("ID", "Type"))
		if err != nil {
			log.Println(err)
		}

		res[idx].Type = typesReview[pkg.RandMaxInt(len(typesReview))]
	}

	return res
}

func (g *DBGenerator) GenerateFilms(volume *devpkg.Volume) []devpkg.FilmFiller {
	res := make([]devpkg.FilmFiller, volume.CountFilms)

	age := []string{"6+", "12+", "16+", "18+", "21+"}

	currency := []string{"USD", "EURO"}

	types := []string{devpkg.TypeSerial, devpkg.TypeFilm}

	for idx := range res {
		// Film
		// Date data
		res[idx].ProdDate = strings.ReplaceAll(faker.Date(), "-", ".")

		// Str data
		res[idx].Name = faker.Name(options.WithRandomStringLength(strLength(volume.MaxFilmsNameLength)))
		res[idx].PosterVer = faker.Word(options.WithRandomStringLength(strLength(volume.MaxFilmsPosterVerLength)))
		res[idx].PosterHor = faker.Word(options.WithRandomStringLength(strLength(volume.MaxFilmsPosterHorLength)))
		res[idx].Description = faker.Word(options.WithRandomStringLength(strLength(volume.MaxFilmDescriptionLength)))
		res[idx].ShortDescription = faker.Word(options.WithRandomStringLength(strLength(volume.MaxFilmShortDescriptionLength)))
		res[idx].OriginalName = faker.Word(options.WithRandomStringLength(strLength(volume.MaxFilmsOriginalNameLength)))
		res[idx].Slogan = faker.Word(options.WithRandomStringLength(strLength(volume.MaxFilmsSloganLength)))

		// Int Data
		res[idx].BoxOfficeDollars = pkg.RandMaxInt(int(strLength(volume.MaxFilmsBoxOfficeDollars)))
		res[idx].Budget = pkg.RandMaxInt(int(strLength(volume.MaxFilmsBudget)))
		res[idx].DurationMinutes = pkg.RandMaxInt(int(strLength(volume.MaxFilmsDurationMinutes)))

		// Enum data
		res[idx].CurrencyBudget = currency[pkg.RandMaxInt(len(currency))]
		res[idx].AgeLimit = age[pkg.RandMaxInt(len(age))]

		res[idx].ProdCountries = make([]string, volume.MaxFilmsProdCountriesCount)
		res[idx].ProdCompanies = make([]string, volume.MaxFilmsProdCompaniesCount)
		res[idx].Genres = make([]string, volume.MaxFilmsPGenresCount)

		// Media
		res[idx].Ticket = faker.Word(options.WithRandomStringLength(strLength(volume.MaxFilmsTicketLength)))
		res[idx].Trailer = faker.Word(options.WithRandomStringLength(strLength(volume.MaxFilmsTrailerLength)))

		// Serial
		// Enum data
		typeFilm := types[pkg.RandMaxInt(len(types))]
		res[idx].Type = typeFilm

		if typeFilm == devpkg.TypeSerial {
			// Int data
			res[idx].CountSeasons = pkg.RandMaxInt(int(strLength(volume.MaxFilmsCountSeasons))+1) + 1

			// Date data
			res[idx].EndYear = strings.ReplaceAll(faker.Date(), "-", ".")[:len(constparams.OnlyDate)]
		}
	}

	return res
}
