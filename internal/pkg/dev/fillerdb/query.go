package fillerdb

import (
	"fmt"
	"strings"
)

func GetBatchInsertFilms(countInserts int) (string, int) {
	queryBegin := `INSERT INTO 
    	films(name, prod_year, poster_ver, poster_hor, description, 
    	      short_description, original_name, slogan, age_limit, 
    	      box_office, budget, duration, currency_budget, type, count_seasons, end_year) 
		VALUES `

	countAttributes := strings.Count(queryBegin, ",") + 1

	placeholders := CreatePlaceholders(countAttributes, countInserts)

	queryEnd := "RETURNING film_id;"

	insertStatement := fmt.Sprintf("%s %s %s", queryBegin, placeholders, queryEnd)

	return insertStatement, countAttributes
}

func GetBatchInsertPersons(countInserts int) (string, int) {
	queryBegin := `INSERT INTO persons(name, original_name, birthday, growth, avatar,  gender, death) VALUES`

	countAttributes := strings.Count(queryBegin, ",") + 1

	placeholders := CreatePlaceholders(countAttributes, countInserts)

	queryEnd := "RETURNING person_id;"

	insertStatement := fmt.Sprintf("%s %s %s", queryBegin, placeholders, queryEnd)

	return insertStatement, countAttributes
}

func GetBatchInsertUsers(countInserts int) (string, int) {
	queryBegin := `INSERT INTO users(nickname, email, password) VALUES`

	countAttributes := strings.Count(queryBegin, ",") + 1

	placeholders := CreatePlaceholders(countAttributes, countInserts)

	queryEnd := "RETURNING user_id;"

	insertStatement := fmt.Sprintf("%s %s %s", queryBegin, placeholders, queryEnd)

	return insertStatement, countAttributes
}

func GetBatchInsertReviews(countInserts int) (string, int) {
	queryBegin := `INSERT INTO reviews(name, type, create_time, body) VALUES`

	countAttributes := strings.Count(queryBegin, ",") + 1

	placeholders := CreatePlaceholders(countAttributes, countInserts)

	queryEnd := "RETURNING review_id;"

	insertStatement := fmt.Sprintf("%s %s %s", queryBegin, placeholders, queryEnd)

	return insertStatement, countAttributes
}

func GetBatchInsertProfiles(countInserts int) (string, int) {
	queryBegin := `INSERT INTO profiles(profile_id) VALUES`

	countAttributes := strings.Count(queryBegin, ",") + 1

	placeholders := CreatePlaceholders(countAttributes, countInserts)

	insertStatement := fmt.Sprintf("%s %s", queryBegin, placeholders)

	return insertStatement, countAttributes
}

func GetBatchInsertReviewsLikes(countInserts int) (string, int) {
	queryBegin := `INSERT INTO reviews_likes(fk_review_id, fk_profile_id) VALUES`

	countAttributes := strings.Count(queryBegin, ",") + 1

	placeholders := CreatePlaceholders(countAttributes, countInserts)

	insertStatement := fmt.Sprintf("%s %s", queryBegin, placeholders)

	return insertStatement, countAttributes
}

func GetBatchInsertPersonProfessions(countInserts int) (string, int) {
	queryBegin := `INSERT INTO person_professions(fk_person_id, fk_profession_id, weight) VALUES`

	countAttributes := strings.Count(queryBegin, ",") + 1

	placeholders := CreatePlaceholders(countAttributes, countInserts)

	insertStatement := fmt.Sprintf("%s %s", queryBegin, placeholders)

	return insertStatement, countAttributes
}
