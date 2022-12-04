package auth

const (
	CreateUser = `
INSERT INTO users (email,
                   nickname,
                   password,
                   count_collections,
                   avatar)
VALUES ($1, $2, $3, 2, $4)
RETURNING user_id`

	CreateDefCollections = `
INSERT INTO collections(name, description, poster)
VALUES ('Избранное', 'Ваши сохранные фильмы', '1'), ('Буду смотреть', 'Фильмы, которые вы отметили для просмотра', '2')
RETURNING collection_id;`

	LinkUserDefCollections = `
INSERT INTO user_collections (fk_user_id, fk_collection_id)
VALUES ($1, $2), ($1, $3)`

	GetUserByEmail = `
SELECT 
	user_id,
	email,
	nickname,
	password,
	avatar
FROM users
WHERE email = $1`

	GetUserByID = `
SELECT 
	email,
	nickname,
	password,
	avatar
FROM users
WHERE user_id = $1`

	CheckExist = `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	UpdateUserPassword = `UPDATE users SET password = $1 WHERE user_id = $2`
)
