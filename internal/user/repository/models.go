package repository

import (
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"time"
)

type UserSQL struct {
	ID               int
	Nickname         sql.NullString
	Email            sql.NullString
	Password         []byte
	Avatar           sql.NullString
	JoinedDate       time.Time
	CountViewsFilms  sql.NullInt32
	CountCollections sql.NullInt32
	CountReviews     sql.NullInt32
	CountRatings     sql.NullInt32
}

func NewUserSQL() UserSQL {
	return UserSQL{}
}

func NewUserSQLOnUser(user *models.User) UserSQL {
	joinedDate, _ := time.Parse(innerPKG.DateFormat, user.JoinedDate)

	return UserSQL{
		ID:               user.ID,
		Nickname:         sqltools.NewSQLNullString(user.Nickname),
		Email:            sqltools.NewSQLNullString(user.Email),
		Password:         []byte(user.Password),
		Avatar:           sqltools.NewSQLNullString(user.Avatar),
		JoinedDate:       joinedDate,
		CountViewsFilms:  sqltools.NewSQLNullInt32(user.CountViewsFilms),
		CountCollections: sqltools.NewSQLNullInt32(user.CountCollections),
		CountReviews:     sqltools.NewSQLNullInt32(user.CountReviews),
		CountRatings:     sqltools.NewSQLNullInt32(user.CountRatings),
	}
}

func (u *UserSQL) Convert() models.User {
	return models.User{
		ID:               u.ID,
		Nickname:         u.Nickname.String,
		Email:            u.Email.String,
		Password:         string(u.Password),
		Avatar:           u.Avatar.String,
		JoinedDate:       u.JoinedDate.Format(innerPKG.DateFormat),
		CountViewsFilms:  int(u.CountViewsFilms.Int32),
		CountCollections: int(u.CountCollections.Int32),
		CountReviews:     int(u.CountReviews.Int32),
		CountRatings:     int(u.CountRatings.Int32),
	}
}

type NodeInUserCollectionSQL struct {
	ID     int
	Name   string
	IsUsed bool
}

type UserActivitySQL struct {
	CountReviews    sql.NullInt32
	Rating          sql.NullInt32
	DateRating      sql.NullTime
	ListCollections []NodeInUserCollectionSQL
}

func NewUserActivitySQL() UserActivitySQL {
	return UserActivitySQL{}
}

func (u *UserActivitySQL) Convert() models.UserActivity {
	rateDate := ""
	if u.DateRating.Valid {
		rateDate = u.DateRating.Time.Format(innerPKG.DateFormat)
	}

	res := models.UserActivity{
		CountReviews: int(u.CountReviews.Int32),
		Rating:       int(u.Rating.Int32),
		DateRating:   rateDate,
		Collections:  make([]models.NodeInUserCollection, len(u.ListCollections)),
	}

	for idx, value := range u.ListCollections {
		res.Collections[idx].ID = value.ID
		res.Collections[idx].Name = value.Name
		res.Collections[idx].IsUsed = value.IsUsed
	}

	return res
}
