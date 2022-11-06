package fillerdb

const (
	insertFilms = `INSERT INTO 
    	films(name, prod_year, poster_ver, poster_hor, description, 
    	      short_description, original_name, slogan, age_limit, 
    	      box_office, budget, duration, currency_budget, type, count_seasons, end_year) 
		VALUES `

	insertFilmsGenres    = `INSERT INTO film_genres(fk_film_id, fk_genre_id, weight) VALUES`
	insertFilmsCountries = `INSERT INTO film_countries(fk_film_id, fk_country_id, weight) VALUES`
	insertFilmsCompanies = `INSERT INTO film_companies(fk_film_id, fk_company_id, weight) VALUES`
	insertFilmsTags      = `INSERT INTO film_tags(fk_film_id, fk_tag_id) VALUES`
	insertFilmsImages    = `INSERT INTO film_images(film_id, images_list) VALUES`

	insertPersons = `INSERT INTO persons(name, original_name, birthday, growth, avatar,  gender, death) VALUES`

	insertPersonsProfessions = `INSERT INTO person_professions(fk_person_id, fk_profession_id, weight) VALUES`
	insertPersonsGenres      = `INSERT INTO person_genres(fk_person_id, fk_genre_id, weight) VALUES`
	insertPersonsImages      = `INSERT INTO person_images(person_id, images_list) VALUES`

	insertUsers = `INSERT INTO users(nickname, email, password) VALUES`

	insertUsersProfiles  = `INSERT INTO profiles(profile_id, joined_date) VALUES`
	insertProfileViews   = `INSERT INTO profile_views_films(fk_profile_id, fk_film_id, create_date) VALUES`
	insertProfileRatings = `INSERT INTO profile_ratings(fk_profile_id, fk_film_id, score, create_date) VALUES`

	insertReviews = `INSERT INTO reviews(name, type, create_time, body) VALUES`

	insertReviewsLikes = `INSERT INTO reviews_likes(fk_review_id, fk_profile_id, create_date) VALUES`

	insertFilmsReviews = `INSERT INTO profile_reviews(fk_review_id, fk_profile_id, fk_film_id) VALUES`

	insertFilmsPersons = `INSERT INTO film_persons(fk_person_id, fk_film_id, fk_profession_id, character, weight) VALUES`

	updateFilms = `UPDATE films f
SET (rating, count_scores) =
        (SELECT SUM(pr.score) / COUNT(*) AS rating,
                COUNT(*)
         FROM profile_ratings pr
         WHERE f.film_id = fk_film_id
         GROUP BY pr.fk_film_id
         ORDER BY rating DESC),
    count_negative_reviews =
        (SELECT COUNT(*)
         FROM profile_reviews
                  JOIN reviews r on profile_reviews.fk_review_id = r.review_id
         WHERE f.film_id = fk_film_id
           AND r.type = 'negative'
         HAVING COUNT(profile_reviews.fk_film_id) > 0),
    count_neutral_reviews  =
        (SELECT COUNT(*)
         FROM profile_reviews
                  JOIN reviews r on profile_reviews.fk_review_id = r.review_id
         WHERE f.film_id = fk_film_id
           AND r.type = 'neutral'
         HAVING COUNT(*) > 0),
    count_positive_reviews =
        (SELECT COUNT(*) as count
         FROM profile_reviews
                  JOIN reviews r on profile_reviews.fk_review_id = r.review_id
         WHERE f.film_id = fk_film_id
           AND r.type = 'positive'
         HAVING COUNT(*) > 0),
    count_actors           =
        (SELECT COUNT(*) as count
         FROM film_persons fp
         WHERE f.film_id = fk_film_id
           AND fp.fk_profession_id = (SELECT profession_id FROM professions p WHERE p.name = 'актер')
         HAVING COUNT(fp.fk_film_id) > 0)`

	updatePersons = `UPDATE persons p
SET count_films = (SELECT COUNT(*) as count
                   FROM film_persons fp
                   WHERE fp.fk_person_id = p.person_id
                   GROUP BY fp.fk_person_id, fp.fk_profession_id
                   HAVING fp.fk_profession_id = (SELECT profession_id FROM professions WHERE name = 'актер'));`

	updateProfiles = `UPDATE profiles p
SET count_views_films=
        (SELECT COUNT(*)
         FROM profile_views_films pvf
         WHERE p.profile_id = pvf.fk_profile_id
         GROUP BY pvf.fk_profile_id
         HAVING COUNT(pvf.fk_profile_id) > 0),
    count_reviews=
        (SELECT COUNT(*)
         FROM profile_reviews pr
         WHERE p.profile_id = pr.fk_profile_id
         GROUP BY pr.fk_profile_id
         HAVING COUNT(pr.fk_profile_id) > 0),
    count_ratings=
        (SELECT COUNT(*)
         FROM profile_ratings pr
         WHERE p.profile_id = pr.fk_profile_id
         GROUP BY pr.fk_profile_id
         HAVING COUNT(pr.fk_profile_id) > 0),
    count_collections=
        (SELECT COUNT(*)
         FROM profile_collections pc
         WHERE p.profile_id = pc.fk_profile_id
         GROUP BY pc.fk_profile_id
         HAVING COUNT(pc.fk_profile_id) > 0)`

	updateReviews = `
UPDATE reviews r
SET count_likes = (SELECT COUNT(*) as count
                   FROM reviews_likes rl
                   WHERE rl.fk_review_id = r.review_id
                   GROUP BY rl.fk_review_id);`

	insertCollections        = `INSERT INTO collections(name, description, poster, create_time) VALUES`
	insertProfileCollections = `INSERT INTO profile_collections(fk_collection_id, fk_profile_id) VALUES`
)
