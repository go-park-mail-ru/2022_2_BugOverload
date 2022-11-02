package fillerdb

const (
	insertFilms = `INSERT INTO 
    	films(name, prod_year, poster_ver, poster_hor, description, 
    	      short_description, original_name, slogan, age_limit, 
    	      box_office, budget, duration, currency_budget, type, count_seasons, end_year) 
		VALUES `
	insertFilmsEnd = " RETURNING film_id;"

	insertFilmsGenres    = `INSERT INTO film_genres(fk_film_id, fk_genre_id, weight) VALUES`
	insertFilmsCountries = `INSERT INTO film_countries(fk_film_id, fk_country_id, weight) VALUES`
	insertFilmsCompanies = `INSERT INTO film_companies(fk_film_id, fk_company_id, weight) VALUES`
	insertFilmsTags      = `INSERT INTO film_tags(fk_film_id, fk_tag_id) VALUES`

	insertPersons    = `INSERT INTO persons(name, original_name, birthday, growth, avatar,  gender, death) VALUES`
	insertPersonsEnd = " RETURNING person_id;"

	insertPersonsProfessions = `INSERT INTO person_professions(fk_person_id, fk_profession_id, weight) VALUES`
	insertPersonsGenres      = `INSERT INTO person_genres(fk_person_id, fk_genre_id, weight) VALUES`

	insertUsers    = `INSERT INTO users(nickname, email, password) VALUES`
	insertUsersEnd = " RETURNING user_id;"

	insertUsersProfiles  = `INSERT INTO profiles(profile_id) VALUES`
	insertProfileViews   = `INSERT INTO profile_views_films(fk_profile_id, fk_film_id) VALUES`
	insertProfileRatings = `INSERT INTO profile_ratings(fk_profile_id, fk_film_id, score) VALUES`

	insertReviews    = `INSERT INTO reviews(name, type, create_time, body) VALUES`
	insertReviewsEnd = " RETURNING review_id;"

	insertReviewsLikes = `INSERT INTO reviews_likes(fk_review_id, fk_profile_id) VALUES`

	insertFilmsReviews = `INSERT INTO profile_reviews(fk_review_id, fk_profile_id, fk_film_id) VALUES`

	insertFilmsPersons = `INSERT INTO film_persons(fk_person_id, fk_film_id, fk_profession_id, character, weight) VALUES`
)
