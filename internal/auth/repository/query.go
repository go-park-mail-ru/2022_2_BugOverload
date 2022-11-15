package repository

const (
	createUser = `
INSERT INTO users (email,
                   nickname,
                   password,
                   count_collections,
                   avatar)
VALUES ($1, $2, $3, 2, $4)
RETURNING user_id, email, nickname, avatar`

	createDefCollections = `
INSERT INTO collections(name, description)
VALUES ('Избранное', 'Ваши сохранные фильмы'), ('Буду смотреть', 'Фильмы, которые вы отметили для просмотра')
RETURNING collection_id;`

	linkUserDefCollections = `
INSERT INTO profile_collections (fk_user_id, fk_collection_id)
VALUES ($1, $2), ($1, $3)`

	getUserByEmail = `
SELECT 
	user_id,
	email,
	nickname,
	password,
	avatar
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

	checkExist = `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	updateUserPassword = `UPDATE users SET password = $1 WHERE user_id = $2`
)
