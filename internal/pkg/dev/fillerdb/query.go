package fillerdb

func getBatchInsertFilms(countInserts int) (string, int) {
	queryBegin := `INSERT INTO 
    	films(name, prod_year, poster_ver, poster_hor, description, 
    	      short_description, original_name, slogan, age_limit, 
    	      box_office, budget, duration, currency_budget, type, count_seasons, end_year) 
		VALUES `

	insertStatement, countAttributes := createStatement(queryBegin, countInserts)

	queryEnd := " RETURNING film_id;"

	return insertStatement + queryEnd, countAttributes
}

func getBatchInsertPersons(countInserts int) (string, int) {
	queryBegin := `INSERT INTO persons(name, original_name, birthday, growth, avatar,  gender, death) VALUES`

	insertStatement, countAttributes := createStatement(queryBegin, countInserts)

	queryEnd := " RETURNING person_id;"

	return insertStatement + queryEnd, countAttributes
}

func getBatchInsertUsers(countInserts int) (string, int) {
	queryBegin := `INSERT INTO users(nickname, email, password) VALUES`

	insertStatement, countAttributes := createStatement(queryBegin, countInserts)

	queryEnd := " RETURNING user_id;"

	return insertStatement + queryEnd, countAttributes
}

func getBatchInsertReviews(countInserts int) (string, int) {
	queryBegin := `INSERT INTO reviews(name, type, create_time, body) VALUES`

	insertStatement, countAttributes := createStatement(queryBegin, countInserts)

	queryEnd := "RETURNING review_id;"

	return insertStatement + queryEnd, countAttributes
}

func getBatchInsertProfiles(countInserts int) (string, int) {
	queryBegin := `INSERT INTO profiles(profile_id) VALUES`

	return createStatement(queryBegin, countInserts)
}

func getBatchInsertReviewsLikes(countInserts int) (string, int) {
	queryBegin := `INSERT INTO reviews_likes(fk_review_id, fk_profile_id) VALUES`

	return createStatement(queryBegin, countInserts)
}

func getBatchInsertPersonProfessions(countInserts int) (string, int) {
	queryBegin := `INSERT INTO person_professions(fk_person_id, fk_profession_id, weight) VALUES`

	return createStatement(queryBegin, countInserts)
}

func getBatchInsertFilmReviews(countInserts int) (string, int) {
	queryBegin := `INSERT INTO profile_reviews(fk_review_id, fk_profile_id, fk_film_id) VALUES`

	return createStatement(queryBegin, countInserts)
}

func getBatchInsertFilmGenres(countInserts int) (string, int) {
	queryBegin := `INSERT INTO film_genres(fk_film_id, fk_genre_id, weight) VALUES`

	return createStatement(queryBegin, countInserts)
}

func getBatchInsertFilmCountries(countInserts int) (string, int) {
	queryBegin := `INSERT INTO film_countries(fk_film_id, fk_country_id, weight) VALUES`

	return createStatement(queryBegin, countInserts)
}

func getBatchInsertFilmCompanies(countInserts int) (string, int) {
	queryBegin := `INSERT INTO film_companies(fk_film_id, fk_company_id, weight) VALUES`

	return createStatement(queryBegin, countInserts)
}

func getBatchInsertPersonGenres(countInserts int) (string, int) {
	queryBegin := `INSERT INTO person_genres(fk_person_id, fk_genre_id, weight) VALUES`

	return createStatement(queryBegin, countInserts)
}

func getBatchInsertFilmPersons(countInserts int) (string, int) {
	queryBegin := `INSERT INTO film_persons(fk_person_id, fk_film_id, fk_profession_id, character, weight) VALUES`

	return createStatement(queryBegin, countInserts)
}

func getBatchInsertFilmTags(countInserts int) (string, int) {
	queryBegin := `INSERT INTO film_tags(fk_film_id, fk_tag_id) VALUES`

	return createStatement(queryBegin, countInserts)
}

func getBatchInsertProfileViews(countInserts int) (string, int) {
	queryBegin := `INSERT INTO profile_views_films(fk_profile_id, fk_film_id) VALUES`

	return createStatement(queryBegin, countInserts)
}

func getBatchInsertProfileRatings(countInserts int) (string, int) {
	queryBegin := `INSERT INTO profile_ratings(fk_profile_id, fk_film_id, score) VALUES`

	return createStatement(queryBegin, countInserts)
}
