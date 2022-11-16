-- Зависимые таблицы N:M
DROP TABLE IF EXISTS film_genres;

DROP TABLE IF EXISTS film_countries;

DROP TABLE IF EXISTS film_companies;

DROP TABLE IF EXISTS film_tags;

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

DROP TABLE IF EXISTS reviews_likes;

-- Зависимые таблицы 1:M
DROP TABLE IF EXISTS film_images;

DROP TABLE IF EXISTS person_images;

-- Главные таблицы
DROP TABLE IF EXISTS reviews;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS serials;

DROP TABLE IF EXISTS films;

DROP TABLE IF EXISTS genres;

DROP TABLE IF EXISTS tags;

DROP TABLE IF EXISTS countries;

DROP TABLE IF EXISTS companies;

DROP TABLE IF EXISTS persons;

DROP TABLE IF EXISTS professions;

DROP TABLE IF EXISTS collections;

-- Enums
DROP TYPE IF EXISTS film_type_enum CASCADE;

DROP TYPE IF EXISTS currency_enum CASCADE;

DROP TYPE IF EXISTS age_limit_enum CASCADE;

DROP TYPE IF EXISTS gender_enum CASCADE;

DROP TYPE IF EXISTS type_review_enum CASCADE;
