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

	updateUserSettingsNickname = `UPDATE users SET nickname = $1 WHERE user_id = $2`

	updateUserSettingsPassword = `UPDATE users SET password = $1 WHERE user_id = $2`

	getPass = `
SELECT password FROM users WHERE user_id = $1`

	setRateFilm = `
INSERT INTO profile_ratings(fk_profile_id, fk_film_id, score)
VALUES ($1, $2, $3)
ON CONFLICT (fk_profile_id, fk_film_id) DO UPDATE SET score = $3`

	dropRateFilm = `DELETE FROM profile_ratings WHERE fk_profile_id = $1 AND fk_film_id = $2`

	insertNewReview = `INSERT INTO reviews (name, type, body) VALUES ($1, $2, $3) RETURNING review_id`

	linkNewReviewAuthor = `INSERT INTO profile_reviews (fk_review_id, fk_profile_id, fk_film_id) VALUES ($1, $2, $3)`
)
