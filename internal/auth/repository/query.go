package repository

const (
	createUser = `
		INSERT INTO users (
			email,
			nickname,
			password
		)
		VALUES (
			$1,
			$2,
			$3
		)`

	getUserByEmail = `
		SELECT 
		    user_id,
			nickname,
			email,
			password
		FROM users
		WHERE email = $1`

	getProfileAvatar = `
		SELECT 
			avatar
		FROM profiles
		WHERE profile_id = $1`

	checkExist = `
		SELECT EXISTS(
		    SELECT 1
			FROM users
			WHERE email = $1
		)`
)
