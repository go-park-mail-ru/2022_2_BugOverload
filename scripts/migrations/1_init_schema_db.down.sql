-- Зависимые таблицы N:M
DROP TABLE  IF EXISTS film_genres;

DROP TABLE IF EXISTS film_countries;

DROP TABLE IF EXISTS film_persons;

DROP TABLE IF EXISTS profile_ratings;

DROP TABLE IF EXISTS profile_views_films;

DROP TABLE IF EXISTS profile_reviews;

DROP TABLE IF EXISTS profile_collections;

DROP TABLE IF EXISTS person_professions;

DROP TABLE IF EXISTS person_genres;

DROP TABLE IF EXISTS collections_genres;

DROP TABLE IF EXISTS collection_likes;

DROP TABLE IF EXISTS collections_films;

-- Зависимые таблицы 1:M
DROP TABLE IF EXISTS film_images;

DROP TABLE IF EXISTS person_images;

-- Главные таблицы 1:M
DROP TABLE IF EXISTS profiles;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS films;

DROP TABLE IF EXISTS genres;

DROP TABLE IF EXISTS countries;

DROP TABLE IF EXISTS persons;

DROP TABLE IF EXISTS professions;

DROP TABLE IF EXISTS collections;
