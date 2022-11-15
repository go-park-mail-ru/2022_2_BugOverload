package repository

const (
	getUserProfile = `
SELECT nickname,
       joined_date,
       avatar,
       count_views_films,
       count_collections,
       count_reviews,
       count_ratings
FROM users
WHERE user_id = $1`

	getUserProfileShort = `
SELECT joined_date,
       count_views_films,
       count_collections,
       count_reviews,
       count_ratings
FROM users
WHERE user_id = $1`

	getUserCollections = `
SELECT c.name,
       EXISTS(SELECT 1 FROM collections_films cf WHERE cf.fk_collection_id IN (c.collection_id) AND cf.fk_film_id = $2)
FROM collections c
         JOIN profile_collections pc on c.collection_id = pc.fk_collection_id
WHERE pc.fk_user_id = $1`

	getUserCountReviews = `SELECT count_reviews FROM users WHERE user_id = $1`

	getUserRatingOnFilm = `SELECT score, create_date FROM profile_ratings WHERE fk_user_id = $1 AND fk_film_id = $2`

	updateUserSettingsNickname = `UPDATE users SET nickname = $1 WHERE user_id = $2`

	updateUserSettingsPassword = `UPDATE users SET password = $1 WHERE user_id = $2`

	getPass = `SELECT password FROM users WHERE user_id = $1`

	setRateFilm = `
INSERT INTO profile_ratings(fk_user_id, fk_film_id, score)
VALUES ($1, $2, $3)
ON CONFLICT (fk_user_id, fk_film_id) DO UPDATE SET score = $3`

	dropRateFilm = `DELETE FROM profile_ratings WHERE fk_user_id = $1 AND fk_film_id = $2`

	insertNewReview = `INSERT INTO reviews (name, type, body) VALUES ($1, $2, $3) RETURNING review_id`

	linkNewReviewAuthor = `INSERT INTO profile_reviews (fk_review_id, fk_user_id, fk_film_id) VALUES ($1, $2, $3)`

	updateAuthorCountReviews = `
UPDATE users
SET count_reviews = count_reviews + 1
WHERE user_id = $1`

	updateAuthorCountRatingsUp = `
UPDATE users
SET count_ratings = count_ratings + 1
WHERE user_id = $1`

	updateAuthorCountRatingsDown = `
UPDATE users
SET count_ratings = count_ratings - 1
WHERE user_id = $1`

	updateFilmCountReviewPositive = `
UPDATE films
SET count_positive_reviews = count_positive_reviews + 1
WHERE film_id = $1`

	updateFilmCountReviewNegative = `
UPDATE films
SET count_negative_reviews = count_negative_reviews + 1
WHERE film_id = $1`

	updateFilmCountReviewNeutral = `
UPDATE films
SET count_neutral_reviews = count_neutral_reviews + 1
WHERE film_id = $1`

	updateFilmCountRatingsUp = `
UPDATE films
SET count_ratings = count_ratings + 1
WHERE film_id = $1`

	updateFilmCountRatingsDown = `
UPDATE films
SET count_ratings = count_ratings - 1
WHERE film_id = $1`
)
