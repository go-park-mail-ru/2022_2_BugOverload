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
	Password []byte
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
		Password: []byte(user.Password),
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
		Password: string(u.Password),
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

type NodeInUserCollectionSQL struct {
	NameCollection string
	IsUsed         bool
}

type UserActivitySQL struct {
	CountReviews    sql.NullInt32
	Rating          sql.NullFloat64
	DateRating      sql.NullTime
	ListCollections []NodeInUserCollectionSQL
}

func NewUserActivitySQL() UserActivitySQL {
	return UserActivitySQL{}
}

func (u *UserActivitySQL) Convert() models.UserActivity {
	rateDate := ""
	if u.DateRating.Valid {
		rateDate = u.DateRating.Time.Format("2006.01.02")
	}

	res := models.UserActivity{
		CountReviews: int(u.CountReviews.Int32),
		Rating:       float32(u.Rating.Float64),
		DateRating:   rateDate,
		Collections:  make([]models.NodeInUserCollection, len(u.ListCollections)),
	}

	for idx, value := range u.ListCollections {
		res.Collections[idx].NameCollection = value.NameCollection
		res.Collections[idx].IsUsed = value.IsUsed
	}

	return res
}
