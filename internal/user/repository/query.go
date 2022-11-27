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
SELECT c.collection_id,
       c.name,
       EXISTS(SELECT 1 FROM collections_films cf WHERE cf.fk_collection_id IN (c.collection_id) AND cf.fk_film_id = $2)
FROM collections c
         JOIN user_collections uc on c.collection_id = uc.fk_collection_id
WHERE uc.fk_user_id = $1`

	getUserCountReviews = `SELECT count_reviews FROM users WHERE user_id = $1`

	getUserRatingOnFilm = `SELECT score, create_date FROM user_ratings WHERE fk_user_id = $1 AND fk_film_id = $2`

	updateUserSettingsNickname = `UPDATE users SET nickname = $1 WHERE user_id = $2`

	checkUserRateExist = `SELECT EXISTS(SELECT 1 FROM user_ratings WHERE fk_user_id = $1 AND fk_film_id = $2)`

	updateUserRateFilm = `
UPDATE user_ratings
SET score = $3, create_date = now()
WHERE fk_user_id = $1 AND fk_film_id = $2`

	getFilmRatingsCount = `
SELECT count_ratings FROM films
WHERE film_id = $1`

	setUserRateFilm = `
INSERT INTO user_ratings(fk_user_id, fk_film_id, score)
VALUES ($1, $2, $3)`

	deleteUserRateFilm = `DELETE FROM user_ratings WHERE fk_user_id = $1 AND fk_film_id = $2`

	updateFilmRating = `
UPDATE films
SET rating = (
    SELECT cast(sum(score) as real) / cast(count(score) as real)
    FROM user_ratings ur
    WHERE ur.fk_film_id = $1
)
WHERE film_id = $1
RETURNING rating`

	insertNewReview = `INSERT INTO reviews (name, type, body) VALUES ($1, $2, $3) RETURNING review_id`

	linkNewReviewAuthor = `INSERT INTO user_reviews (fk_review_id, fk_user_id, fk_film_id) VALUES ($1, $2, $3)`

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
WHERE film_id = $1
RETURNING count_ratings`

	updateFilmCountRatingsDown = `
UPDATE films
SET count_ratings = count_ratings - 1
WHERE film_id = $1
RETURNING count_ratings`

	getUserCollectionByCreateDate = `
SELECT c.collection_id,
       c.name,
       c.poster,
       c.count_films,
       c.count_likes,
       c.create_time,
       c.updated_at
FROM collections c
         JOIN user_collections uc on c.collection_id = uc.fk_collection_id
WHERE uc.fk_user_id = $1
  AND uc.user_type_relation = 'author'
  AND c.create_time < $2
ORDER BY c.create_time DESC
LIMIT $3`

	getUserCollectionByUpdateDate = `
SELECT c.collection_id,
       c.name,
       c.poster,
       c.count_films,
       c.count_likes,
       c.create_time,
       c.updated_at
FROM collections c
         JOIN user_collections uc on c.collection_id = uc.fk_collection_id
WHERE uc.fk_user_id = $1
  AND uc.user_type_relation = 'author'
  AND c.updated_at < $2
ORDER BY c.updated_at DESC
LIMIT $3`

	checkFilmExist                    = `SELECT EXISTS (SELECT 1 FROM films WHERE film_id = $1)`
	checkCollectionExist              = `SELECT EXISTS (SELECT 1 FROM collections WHERE collection_id = $1)`
	checkUserAccessToUpdateCollection = `
SELECT EXISTS(
    SELECT 1 FROM user_collections
    WHERE fk_user_id = $1 AND fk_collection_id = $2 AND user_type_relation = 'author'
)`
	checkFilmExistInCollection = `SELECT EXISTS(SELECT 1 FROM collections_films WHERE fk_collection_id = $1 AND fk_film_id = $2)`
	addFilmToCollection        = `INSERT INTO collections_films (fk_collection_id, fk_film_id) VALUES ($1, $2)`
	dropFilmFromCollection     = `DELETE FROM collections_films WHERE fk_collection_id = $1 AND fk_film_id = $2`
)
