package repository

const (
	getFilmsByTagRating = `
SELECT f.film_id,
       f.name,
       f.original_name,
       extract(YEAR FROM  f.prod_date),
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
       extract(YEAR FROM  f.prod_date),
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
       extract(YEAR FROM  f.prod_date),
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

	getUserCollectionByCreateDate = `
SELECT c.collection_id,
       c.name,
       c.poster,
       c.count_films,
       c.count_likes
FROM collections c
         JOIN profile_collections pc on c.collection_id = pc.fk_collection_id
WHERE pc.fk_user_id = $1
  AND pc.user_type_relation = 'author'
  AND c.create_time < $2
ORDER BY c.create_time DESC
LIMIT $3`

	getUserCollectionByUpdateDate = `
SELECT c.collection_id,
       c.name,
       c.poster,
       c.count_films,
       c.count_likes
FROM collections c
         JOIN profile_collections pc on c.collection_id = pc.fk_collection_id
WHERE pc.fk_user_id = $1
  AND pc.user_type_relation = 'author'
  AND c.updated_at < $2
ORDER BY c.updated_at DESC
LIMIT $3`
)
