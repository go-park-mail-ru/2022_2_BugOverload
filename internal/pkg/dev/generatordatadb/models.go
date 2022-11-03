package generatordatadb

type UserFace struct {
	ID       int
	Nickname string `faker:"username"`
	Email    string `faker:"email"`
	Password string `faker:"password"`
}

type ProfileFake struct {
	Avatar           string
	JoinedDate       string
	CountViewsFilms  int
	CountCollections int
	CountReviews     int
	CountRatings     int
}

type ReviewFace struct {
	ID   int
	Name string `faker:"word"`
	Type string
	Time string `faker:"timestamp"`
	Body string `faker:"lang=rus"`
}
