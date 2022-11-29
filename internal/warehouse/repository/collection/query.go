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

	checkUserIsCollectionAuthor = `
SELECT EXISTS (
    SELECT 1 FROM user_collections
    WHERE fk_collection_id = $1
      AND fk_user_id = $2
      AND user_type_relation = 'author'
)`
	checkCollectionIsPublic = `SELECT is_public FROM collections WHERE collection_id = $1`

	getCollectionFilmsByDate = `
SELECT f.film_id,
       f.name,
       f.original_name,
       extract(YEAR FROM f.prod_date),
       f.poster_ver,
       f.type,
       f.rating
FROM films f
WHERE f.film_id IN (
    SELECT cf.fk_film_id
    FROM collections_films cf
    WHERE fk_collection_id = $1
)
ORDER BY f.prod_date DESC`

	getCollectionFilmsByRating = `
SELECT f.film_id,
       f.name,
       f.original_name,
       extract(YEAR FROM f.prod_date),
       f.poster_ver,
       f.type,
       f.rating
FROM films f
WHERE f.film_id IN (
    SELECT cf.fk_film_id
    FROM collections_films cf
    WHERE fk_collection_id = $1
)
ORDER BY f.rating DESC`

	getCollectionShortInfo = `
SELECT name, description
FROM collections
WHERE collection_id = $1`

	getAuthorByID = `
SELECT nickname
FROM users
WHERE user_id = $1`
)
