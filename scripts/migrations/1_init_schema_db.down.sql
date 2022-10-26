-- Зависимые таблицы N:M
DROP TABLE  IF EXISTS cinema_genres;

DROP TABLE IF EXISTS cinema_countries;

DROP TABLE IF EXISTS cinema_persons;

DROP TABLE IF EXISTS profile_ratings;

DROP TABLE IF EXISTS profile_views_cinemas;

DROP TABLE IF EXISTS profile_reviews;

DROP TABLE IF EXISTS profile_collections;

DROP TABLE IF EXISTS person_professions;

DROP TABLE IF EXISTS person_genres;

DROP TABLE IF EXISTS collections_genres;

DROP TABLE IF EXISTS collection_likes;

DROP TABLE IF EXISTS collections_cinemas;

-- Зависимые таблицы 1:M
DROP TABLE IF EXISTS cinema_images;

DROP TABLE IF EXISTS person_images;

-- Главные таблицы 1:M
DROP TABLE IF EXISTS mg_user;

DROP TABLE IF EXISTS profile;

DROP TABLE IF EXISTS cinema;

DROP TABLE IF EXISTS genre;

DROP TABLE IF EXISTS country;

DROP TABLE IF EXISTS person;

DROP TABLE IF EXISTS profession;

DROP TABLE IF EXISTS collection;
