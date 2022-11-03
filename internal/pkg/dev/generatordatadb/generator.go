package generatordatadb

import (
	"log"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"

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

		res[idx].Type = typesReview[pkg.RandMaxInt(len(typesReview))]
	}

	return res
}
