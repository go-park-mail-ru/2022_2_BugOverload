package generatordatadb

import (
	"go-park-mail-ru/2022_2_BugOverload/pkg"
	"log"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
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

func (g *DBGenerator) GenerateReviews(count int, maxLengthBody uint) []ReviewFace {
	res := make([]ReviewFace, count)

	typesReview := []string{"positive", "neutral", "negative"}

	for idx := range res {
		err := faker.FakeData(&res[idx],
			options.WithRandomStringLength(maxLengthBody),
			options.WithFieldsToIgnore("ID", "Type"))
		if err != nil {
			log.Println(err)
		}

		res[idx].Type = typesReview[pkg.Rand(len(typesReview))]
	}

	return res
}
