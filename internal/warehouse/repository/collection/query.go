package collection

const (
	GetFilmsByTagRating = `
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
WHERE t.name = $1 AND (f.rating <= $2 OR f.rating IS NULL)
ORDER BY f.rating DESC NULLS LAST
LIMIT $3`

	GetFilmsByTagDate = `
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

	GetFilmsByGenreRating = `
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
WHERE g.name = $1 AND (f.rating <= $2 OR f.rating IS NULL)
ORDER BY f.rating DESC NULLS LAST
LIMIT $3`

	GetFilmsByGenreDate = `
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

	GetTagDescription = `SELECT description FROM tags WHERE name = $1`

	CheckUserIsCollectionAuthor = `
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
ORDER BY f.rating DESC NULLS LAST`

	getCollectionShortInfo = `
SELECT name, description
FROM collections
WHERE collection_id = $1`

	getAuthorID = `
SELECT fk_user_id
FROM user_collections
WHERE fk_collection_id = $1`

	getAuthorCollectionByID = `
SELECT nickname, count_collections, avatar
FROM users
WHERE user_id = $1`

	getSimilarFilms = `
SELECT f.film_id,
       f.name,
       f.original_name,
       extract(YEAR FROM f.prod_date),
       f.poster_ver,
       f.type,
       f.rating
FROM films f
WHERE f.film_id IN (
    SELECT fk_film_id
    FROM film_genres
    WHERE fk_film_id <> $1 AND fk_genre_id IN (
        SELECT fk_genre_id
        FROM film_genres
        WHERE fk_film_id = $1
        ORDER BY weight DESC
        LIMIT 2
    )
    ORDER BY weight DESC)
    LIMIT 10`
)
