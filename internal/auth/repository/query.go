package repository

const (
	createUser = `
INSERT INTO users (email,
                   nickname,
                   password)
VALUES ($1, $2, $3)
RETURNING user_id`

	createUserProfile = `INSERT INTO profiles (profile_id) VALUES ($1)`

	createDefCollections = `
INSERT INTO collections(name, description)
VALUES ('Избранное', 'Ваши сохранные фильмы'), ('Буду смотреть', 'Фильмы, которые вы отметили для просмотра')
RETURNING collection_id;`

	linkUserProfileDefCollections = `
INSERT INTO profile_collections (fk_profile_id, fk_collection_id)
VALUES ($1, $2), ($1, $3)`

	getUserByEmail = `
SELECT 
	user_id,
	email,
	nickname,
	password
FROM users
WHERE email = $1`

	getUserByID = `
SELECT 
	user_id,
	email,
	nickname,
	password
FROM users
WHERE user_id = $1`

	getProfileAvatar = `SELECT avatar FROM profiles WHERE profile_id = $1`

	checkExist = `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
)
