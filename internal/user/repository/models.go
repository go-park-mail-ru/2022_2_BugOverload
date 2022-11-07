package repository

import (
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"time"
)

type UserSQL struct {
	ID       int
	Nickname sql.NullString
	Email    sql.NullString
	Password sql.NullString
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

func NewUserSQLOnUser(user *models.User) UserSQL {
	joinedDate, _ := time.Parse("2006.01.02", user.Profile.JoinedDate)

	return UserSQL{
		ID:       user.ID,
		Nickname: sqltools.NewSQLNullString(user.Nickname),
		Email:    sqltools.NewSQLNullString(user.Email),
		Password: sqltools.NewSQLNullString(user.Password),
		Profile: ProfileSQL{
			Avatar:           sqltools.NewSQLNullString(user.Profile.Avatar),
			JoinedDate:       joinedDate,
			CountViewsFilms:  sqltools.NewSQLNullInt32(user.Profile.CountViewsFilms),
			CountCollections: sqltools.NewSQLNullInt32(user.Profile.CountCollections),
			CountReviews:     sqltools.NewSQLNullInt32(user.Profile.CountReviews),
			CountRatings:     sqltools.NewSQLNullInt32(user.Profile.CountRatings),
		},
	}
}

func (u *UserSQL) Convert() models.User {
	if !u.Profile.Avatar.Valid {
		u.Profile.Avatar.String = "avatar"
	}

	return models.User{
		ID:       u.ID,
		Nickname: u.Nickname.String,
		Email:    u.Email.String,
		Password: u.Password.String,
		Profile: models.Profile{
			Avatar:           u.Profile.Avatar.String,
			JoinedDate:       u.Profile.JoinedDate.Format("2006.01.02"),
			CountViewsFilms:  int(u.Profile.CountViewsFilms.Int32),
			CountCollections: int(u.Profile.CountCollections.Int32),
			CountReviews:     int(u.Profile.CountReviews.Int32),
			CountRatings:     int(u.Profile.CountRatings.Int32),
		},
	}
}
