package pkg

const (
	// Images params request
	ImageObjectFilmPosterHor   = "film_poster_hor"   // key - 1 to 1 from prev request
	ImageObjectFilmPosterVer   = "film_poster_ver"   // key - 1 to 1 from prev request
	ImageObjectFilmImage       = "film_image"        // key - filmID/filmImageKey - Example 1/2
	ImageObjectDefault         = "default"           // key - login or signup
	ImageObjectUserAvatar      = "user_avatar"       // key - 1 to 1 from prev request
	ImageObjectPersonAvatar    = "person_avatar"     // key - 1 to 1 from prev request
	ImageObjectPersonImage     = "person_image"      // key - personID/personImageKey - Example 5/12
	ImageObjectCollectionImage = "collection_poster" // key - 1 to 1 from prev request

	// Image Def Images
	DefFilmPosterHor = "hor"
	DefFilmPosterVer = "ver"
	DefUserAvatar    = "avatar"
	DefPersonAvatar  = "avatar"

	ImageCountSignupLogin = 12

	// S3
	FilmsBucket       = "films/"
	DefBucket         = "default/"
	PersonsBucket     = "persons/"
	UsersBucket       = "users/"
	CollectionsBucket = "collections/"
)
