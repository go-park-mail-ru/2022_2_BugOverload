package fillerdb

const (
	insertFilms = `
INSERT INTO films(name, prod_date, poster_ver, poster_hor, description,
                  short_description, original_name, slogan, age_limit,
                  box_office_dollars, budget, duration_minutes, currency_budget, type)
VALUES`

	insertFilmsMedia = `INSERT INTO media(film_id, ticket, trailer) VALUES`

	insertSerials = `INSERT INTO serials(film_id, count_seasons, end_year) VALUES`

	insertFilmsGenres    = `INSERT INTO film_genres(fk_film_id, fk_genre_id, weight) VALUES`
	insertFilmsCountries = `INSERT INTO film_countries(fk_film_id, fk_country_id, weight) VALUES`
	insertFilmsCompanies = `INSERT INTO film_companies(fk_film_id, fk_company_id, weight) VALUES`
	insertFilmsTags      = `INSERT INTO film_tags(fk_film_id, fk_tag_id) VALUES`
	insertFilmsImages    = `INSERT INTO film_images(film_id, image_key, weight) VALUES`

	insertPersons = `INSERT INTO persons(name, original_name, birthday, growth_meters, avatar,  gender, death) VALUES`

	insertPersonsProfessions = `INSERT INTO person_professions(fk_person_id, fk_profession_id, weight) VALUES`
	insertPersonsGenres      = `INSERT INTO person_genres(fk_person_id, fk_genre_id, weight) VALUES`
	insertPersonsImages      = `INSERT INTO person_images(person_id, image_key, weight) VALUES`

	insertUsers = `INSERT INTO users(nickname, email, password) VALUES`

	insertProfileViews   = `INSERT INTO user_views_films(fk_user_id, fk_film_id, create_date) VALUES`
	insertProfileRatings = `INSERT INTO user_ratings(fk_user_id, fk_film_id, score, create_date) VALUES`

	insertReviews = `INSERT INTO reviews(name, type, create_time, body) VALUES`

	insertReviewsLikes = `INSERT INTO reviews_likes(fk_review_id, fk_user_id, create_date) VALUES`

	insertFilmsReviews = `INSERT INTO user_reviews(fk_review_id, fk_user_id, fk_film_id) VALUES`

	insertFilmsPersons = `INSERT INTO film_persons(fk_person_id, fk_film_id, fk_profession_id, character, weight) VALUES`

	updateFilms = `
UPDATE films f
SET (rating, count_ratings) =
        (SELECT SUM(ur.score) / CAST(COUNT(*) AS float) AS rating,
                COALESCE(COUNT(*), 0)
         FROM user_ratings ur
         WHERE f.film_id = fk_film_id
         GROUP BY ur.fk_film_id),
    count_negative_reviews =
        (SELECT COUNT(*)
         FROM user_reviews
                  JOIN reviews r on user_reviews.fk_review_id = r.review_id
         WHERE f.film_id = fk_film_id
           AND r.type = 'negative'),
    count_neutral_reviews  =
        (SELECT COALESCE(COUNT(*), 0)
         FROM user_reviews
                  JOIN reviews r on user_reviews.fk_review_id = r.review_id
         WHERE f.film_id = fk_film_id
           AND r.type = 'neutral'),
    count_positive_reviews =
        (SELECT COALESCE(COUNT(*), 0)
         FROM user_reviews
                  JOIN reviews r on user_reviews.fk_review_id = r.review_id
         WHERE f.film_id = fk_film_id
           AND r.type = 'positive'),
    count_actors           =
        (SELECT COALESCE(COUNT(*), 0)
         FROM film_persons fp
         WHERE f.film_id = fk_film_id
           AND fp.fk_profession_id = (SELECT profession_id FROM professions p WHERE p.name = 'актер'))`

	updatePersons = `
UPDATE persons p
SET count_films = COALESCE((SELECT SUM(COUNT(DISTINCT (fk_film_id, fk_person_id))) OVER ()
                   FROM film_persons fp
                   WHERE fp.fk_person_id = p.person_id
                   GROUP BY fk_film_id, fk_person_id
                   LIMIT 1), 0);`

	updateProfiles = `
UPDATE users p
SET count_views_films= COALESCE((SELECT COUNT(*) count
         FROM user_views_films uvf
         WHERE p.user_id = uvf.fk_user_id
         GROUP BY uvf.fk_user_id), 0),
    count_reviews=
        COALESCE((SELECT COUNT(*)
         FROM user_reviews ur
         WHERE p.user_id = ur.fk_user_id
         GROUP BY ur.fk_user_id), 0),
    count_ratings=
        COALESCE((SELECT COUNT(*)
         FROM user_ratings ur
         WHERE p.user_id = ur.fk_user_id
         GROUP BY ur.fk_user_id), 0),
    count_collections=
        COALESCE((SELECT COUNT(*)
         FROM user_collections uc
         WHERE p.user_id = uc.fk_user_id
         GROUP BY uc.fk_user_id), 0)`

	updateReviews = `
UPDATE reviews r
SET count_likes = COALESCE((SELECT COUNT(*)
                   FROM reviews_likes rl
                   WHERE rl.fk_review_id = r.review_id
                   GROUP BY rl.fk_review_id), 0);`

	insertCollections        = `INSERT INTO collections(name, description, poster, create_time, is_public) VALUES`
	insertProfileCollections = `INSERT INTO user_collections(fk_collection_id, fk_user_id) VALUES`
)
