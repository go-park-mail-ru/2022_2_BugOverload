package repository

const (
	getUser = `SELECT nickname FROM users WHERE user_id = $1`

	getUserProfile = `
SELECT joined_date,
	   avatar,
       count_views_films,
       count_collections,
       count_reviews,
       count_ratings
FROM profiles
WHERE profile_id = $1`

	getUserProfileShort = `
SELECT joined_date,
       count_views_films,
       count_collections,
       count_reviews,
       count_ratings
FROM profiles
WHERE profile_id = $1`

	getUserCollections = `
SELECT c.name,
       EXISTS(SELECT 1 FROM collections_films cf WHERE cf.fk_collection_id IN (c.collection_id) AND cf.fk_film_id = $2)
FROM collections c
         JOIN profile_collections pc on c.collection_id = pc.fk_collection_id
WHERE pc.fk_profile_id = $1`

	getUserCountReviews = `SELECT count_reviews FROM profiles WHERE profile_id = $1`

	getUserRatingOnFilm = `SELECT score, create_date FROM profile_ratings WHERE fk_profile_id = $1 AND fk_film_id = $2`

	updateUserSettingsNickname = `UPDATE users SET nickname = $1 WHERE user_id = $2`

	setRateFilm = `
INSERT INTO profile_ratings(fk_profile_id, fk_film_id, score)
VALUES ($1, $2, $3)
ON CONFLICT (fk_profile_id, fk_film_id) DO UPDATE SET score = $3`

	dropRateFilm = `DELETE FROM profile_ratings WHERE fk_profile_id = $1 AND fk_film_id = $2`

	insertNewReview = `INSERT INTO reviews (name, type, body) VALUES ($1, $2, $3) RETURNING review_id`

	linkNewReviewAuthor = `INSERT INTO profile_reviews (fk_review_id, fk_profile_id, fk_film_id) VALUES ($1, $2, $3)`

	updateAuthorCountReviews = `
UPDATE profiles
SET count_reviews = CASE
                        WHEN count_reviews IS NULL THEN 1
                        ELSE count_reviews + 1
    END
WHERE profile_id = $1`

	updateAuthorCountRatingsUp = `
UPDATE profiles
SET count_ratings = CASE
                        WHEN count_ratings IS NULL THEN 1
                        ELSE count_ratings + 1
    END
WHERE profile_id = $1`

	updateAuthorCountRatingsDown = `
UPDATE profiles
SET count_ratings = CASE
                        WHEN count_ratings ISNULL THEN NULL
                        WHEN count_ratings = 0 THEN 0
                        WHEN count_ratings > 0 THEN count_ratings - 1
    END
WHERE profile_id = $1`

	updateFilmCountReviewPositive = `UPDATE films
SET count_positive_reviews = CASE
                        WHEN count_positive_reviews IS NULL THEN 1
                        ELSE count_positive_reviews + 1
    END
WHERE film_id = $1`

	updateFilmCountReviewNegative = `UPDATE films
SET count_negative_reviews = CASE
                        WHEN count_negative_reviews IS NULL THEN 1
                        ELSE count_negative_reviews + 1
    END
WHERE film_id = $1`

	updateFilmCountReviewNeutral = `UPDATE films
SET count_neutral_reviews = CASE
                        WHEN count_neutral_reviews IS NULL THEN 1
                        ELSE count_neutral_reviews + 1
    END
WHERE film_id = $1`
)
