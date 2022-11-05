package repository

import (
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"time"
)

type UserSQL struct {
	ID       int
	Nickname string
	Email    string
	Password string
	Profile  ProfileSQL
}

type ProfileSQL struct {
	Avatar           sql.NullString
	JoinedDate       time.Time
	CountViewsFilms  sql.NullInt32
	CountCollections sql.NullInt32
	CountReviews     sql.NullInt32
	CountRatings     sql.NullInt32
}

func NewUserSQL() UserSQL {
	return UserSQL{
		Profile: ProfileSQL{},
	}
}

//  для линтера
// func newUserSQLOnUser(user models.User) UserSQL {
//	return UserSQL{
//		ID:       user.ID,
//		Nickname: user.Nickname,
//		Email:    user.Email,
//		Password: user.Password,
//		Profile: ProfileSQL{
//			Avatar:           innerPKG.NewSQLNullString(user.Profile.Avatar),
//			JoinedDate:       user.Profile.JoinedDate,
//			CountViewsFilms:  innerPKG.NewSQLNullInt32(user.Profile.CountViewsFilms),
//			CountCollections: innerPKG.NewSQLNullInt32(user.Profile.CountCollections),
//			CountReviews:     innerPKG.NewSQLNullInt32(user.Profile.CountReviews),
//			CountRatings:     innerPKG.NewSQLNullInt32(user.Profile.CountRatings),
//		},
//	}
// }

func (u *UserSQL) Convert() models.User {
	if !u.Profile.Avatar.Valid {
		u.Profile.Avatar.String = "avatar"
	}

	return models.User{
		ID:       u.ID,
		Nickname: u.Nickname,
		Email:    u.Email,
		Password: u.Password,
		Profile: models.Profile{
			Avatar:           u.Profile.Avatar.String,
			JoinedDate:       u.Profile.JoinedDate.Format("2006-01-02"),
			CountViewsFilms:  int(u.Profile.CountViewsFilms.Int32),
			CountCollections: int(u.Profile.CountCollections.Int32),
			CountReviews:     int(u.Profile.CountReviews.Int32),
			CountRatings:     int(u.Profile.CountRatings.Int32),
		},
	}
}
