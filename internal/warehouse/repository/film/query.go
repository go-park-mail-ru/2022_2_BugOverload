package film

const (
	getFilmByID = `
SELECT name,
       original_name,
       extract(YEAR FROM  prod_date),
       slogan,
       description,
	   short_description,
       age_limit,
       duration_minutes,
       poster_hor,
       budget,
       box_office_dollars,
       currency_budget,
       type,
       rating,
       count_actors,
       count_ratings,
       count_negative_reviews,
       count_neutral_reviews,
       count_positive_reviews
FROM films
WHERE film_id = $1`

	getSerialByID = `
SELECT count_seasons,
       end_year
FROM serials
WHERE film_id = $1`

	GetShortSerialByID = `
SELECT end_year
FROM serials
WHERE film_id = $1`

	getReviewsByFilmID = `
SELECT r.name,
       r.type,
       r.body,
       r.count_likes,
       r.create_time,
       u.user_id,
       u.nickname,
       u.avatar,
       u.count_reviews
FROM reviews r
         JOIN user_reviews ur on r.review_id = ur.fk_review_id
         JOIN users u on ur.fk_user_id = u.user_id
WHERE ur.fk_film_id = $1
ORDER BY r.create_time DESC
LIMIT $2 OFFSET $3`

	getFilmGenres = `
SELECT g.name
FROM genres g
         JOIN film_genres fg on g.genre_id = fg.fk_genre_id
WHERE fg.fk_film_id = $1
ORDER BY weight DESC`

	getFilmCompanies = `
SELECT c.name
FROM companies c
         JOIN film_companies fc on c.company_id = fc.fk_company_id
WHERE fc.fk_film_id = $1
ORDER BY weight DESC`

	getFilmCountries = `
SELECT c.name
FROM countries c
         JOIN film_countries fc on c.country_id = fc.fk_country_id
WHERE fc.fk_film_id = $1
ORDER BY weight DESC`

	getFilmTags = `
SELECT t.name
FROM tags t
         JOIN film_tags ft on t.tag_id = ft.fk_tag_id
WHERE ft.fk_film_id = $1
ORDER BY t.name DESC`

	getFilmImages = `
SELECT image_key
FROM film_images
WHERE film_id = $1
ORDER BY weight DESC
LIMIT $2`

	getFilmActors = `
SELECT fp.fk_person_id, p.name, p.avatar, fp.character
FROM persons p
         JOIN film_persons fp on p.person_id = fp.fk_person_id
WHERE fp.fk_film_id = $1
  AND fp.character IS NOT NULL
GROUP BY fp.fk_film_id, fp.weight, fp.character, p.name, fp.fk_person_id, p.avatar
ORDER BY fp.weight DESC`

	getFilmPersons = `
SELECT fp.fk_person_id, p.name, fp.fk_profession_id
FROM persons p
         JOIN film_persons fp on p.person_id = fp.fk_person_id
WHERE fp.fk_film_id = $1
  AND fp.character IS NULL
GROUP BY fp.fk_film_id, fp.fk_profession_id, fp.weight, p.name, fp.fk_person_id
ORDER BY fp.fk_profession_id, fp.weight DESC`

	getFilmRecommendation = `
SELECT f.film_id,
       f.name,
       extract(YEAR FROM  f.prod_date),
       f.type,
       f.poster_hor,
       f.short_description,
       f.rating
FROM films f
WHERE f.poster_hor IS NOT NULL AND f.film_id BETWEEN 27 AND 29
   or film_id = 15
ORDER BY random()
LIMIT 1`

	GetNewFilms = `
SELECT f.film_id,
       f.name,
       to_char(f.prod_date, 'YYYY.MM.DD'),
       f.poster_ver,
       f.rating,
       f.duration_minutes,
       f.description
FROM films f
WHERE f.prod_date > now()
  AND f.type = 'film'
ORDER BY f.prod_date, f.film_id
LIMIT $1
OFFSET $2`

	getRelease = `
SELECT f.film_id,
       f.name,
       to_char(f.prod_date, 'YYYY.MM.DD'),
       f.poster_ver,
       f.rating
FROM films f
WHERE f.prod_date > now()
ORDER BY f.prod_date, f.film_id
LIMIT 5`

	getTrailer = `
SELECT m.trailer
FROM media m
WHERE m.film_id = $1
`
)
