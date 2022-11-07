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

	getUserShort = `
SELECT joined_date,
       count_views_films,
       count_collections,
       count_reviews,
       count_ratings
FROM profiles
WHERE profile_id = $1`
)
