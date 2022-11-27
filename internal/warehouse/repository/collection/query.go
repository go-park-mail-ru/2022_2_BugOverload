package collection

const (
	getFilmsByTagRating = `
SELECT f.film_id,
       f.name,
       f.original_name,
       extract(YEAR FROM f.prod_date),
       f.poster_ver,
       f.type,
       f.rating
FROM films f
         JOIN film_tags ft on f.film_id = ft.fk_film_id
         JOIN tags t on ft.fk_tag_id = t.tag_id
WHERE t.name = $1 AND f.rating < $2
ORDER BY f.rating DESC
LIMIT $3`

	getFilmsByTagDate = `
SELECT f.film_id,
       f.name,
       f.original_name,
       extract(YEAR FROM f.prod_date),
       f.poster_ver,
       f.type,
       f.rating
FROM films f
         JOIN film_tags ft on f.film_id = ft.fk_film_id
         JOIN tags t on ft.fk_tag_id = t.tag_id
WHERE t.name = $1 AND f.prod_date < NOW()
ORDER BY f.prod_date DESC
LIMIT $2
OFFSET $3`

	getFilmsByGenreRating = `
SELECT f.film_id,
       f.name,
       f.original_name,
       extract(YEAR FROM f.prod_date),
       f.poster_ver,
       f.type,
       f.rating
FROM films f
        JOIN film_genres fg on f.film_id = fg.fk_film_id
        JOIN genres g on g.genre_id = fg.fk_genre_id
WHERE g.name = $1 AND f.rating < $2
ORDER BY f.rating DESC
LIMIT $3`

	getFilmsByGenreDate = `
SELECT f.film_id,
       f.name,
       f.original_name,
       extract(YEAR FROM f.prod_date),
       f.poster_ver,
       f.type,
       f.rating
FROM films f
        JOIN film_genres fg on f.film_id = fg.fk_film_id
        JOIN genres g on g.genre_id = fg.fk_genre_id
WHERE g.name = $1 AND f.prod_date < NOW()
ORDER BY f.prod_date DESC
LIMIT $2
OFFSET $3`

	getTagDescription = `SELECT description FROM tags WHERE name = $1`
)
