package repository

const (
	updateUserAvatar       = `UPDATE users SET avatar = $1 WHERE user_id = $1`
	updateFilmPosterHor    = `UPDATE films SET poster_hor = $1`
	updateFilmPosterVer    = `UPDATE films SET poster_ver = $1`
	updateFilmImage        = `` // Dangerous, needed update DB
	updatePersonAvatar     = `UPDATE persons SET avatar = $1`
	updatePersonImage      = `` // Dangerous, needed update DB
	updateCollectionPoster = `UPDATE collections SET poster = $1`
)
