package repository

import (
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type userSQL struct {
	ID       int
	Nickname string
	Email    string
	Password string
	Profile  profileSQL
}

type profileSQL struct {
	Avatar           sql.NullString
	JoinedDate       string
	CountViewsFilms  sql.NullInt32
	CountCollections sql.NullInt32
	CountReviews     sql.NullInt32
	CountRatings     sql.NullInt32
}

func newUserSQL() userSQL {
	return userSQL{
		Profile: profileSQL{},
	}
}

//  для линтера
// func newUserSQLOnUser(user models.User) userSQL {
//	return userSQL{
//		ID:       user.ID,
//		Nickname: user.Nickname,
//		Email:    user.Email,
//		Password: user.Password,
//		Profile: profileSQL{
//			Avatar:           innerPKG.NewSQLNullString(user.Profile.Avatar),
//			JoinedDate:       user.Profile.JoinedDate,
//			CountViewsFilms:  innerPKG.NewSQLNullInt32(user.Profile.CountViewsFilms),
//			CountCollections: innerPKG.NewSQLNullInt32(user.Profile.CountCollections),
//			CountReviews:     innerPKG.NewSQLNullInt32(user.Profile.CountReviews),
//			CountRatings:     innerPKG.NewSQLNullInt32(user.Profile.CountRatings),
//		},
//	}
// }

func (u *userSQL) convert() models.User {
	return models.User{
		ID:       u.ID,
		Nickname: u.Nickname,
		Email:    u.Email,
		Password: u.Password,
		Profile: models.Profile{
			Avatar:           u.Profile.Avatar.String,
			JoinedDate:       u.Profile.JoinedDate,
			CountViewsFilms:  int(u.Profile.CountViewsFilms.Int32),
			CountCollections: int(u.Profile.CountCollections.Int32),
			CountReviews:     int(u.Profile.CountReviews.Int32),
			CountRatings:     int(u.Profile.CountRatings.Int32),
		},
	}
}
