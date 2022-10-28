CREATE TABLE IF NOT EXISTS users
(
    "user_id"      serial      NOT NULL PRIMARY KEY,
    "nickname"     varchar(64) NOT NULL,
    "email"        varchar(64) NOT NULL,
    "password"     text        NOT NULL,
    "joined_date"  date        NOT NULL DEFAULT NOW(),
    "is_superuser" boolean     NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS profiles
(
    "profile_id"        serial      NOT NULL PRIMARY KEY REFERENCES users (user_id) ON DELETE CASCADE,
    "avatar"            varchar(80) NOT NULL DEFAULT 'avatar',
    "count_views_films" integer     NOT NULL,
    "count_collections" integer     NOT NULL,
    "count_reviews"     integer     NOT NULL,
    "count_ratings"     integer     NOT NULL
);

CREATE TABLE IF NOT EXISTS films
(
    "film_id"                serial        NOT NULL PRIMARY KEY,
    "name"                   varchar(128)  NOT NULL,
    "prod_year"              integer       NOT NULL,
    "type"                   varchar(64)   NOT NULL DEFAULT 'film',
    "description"            TEXT          NOT NULL,
    "short_description"      varchar(180)  NOT NULL,
    "age_limit"              integer       NOT NULL DEFAULT 13,
    "box_office"             integer                DEFAULT 0,
    "duration"               integer       NOT NULL DEFAULT 90,
    "poster_hor"             varchar(80)   NOT NULL DEFAULT 'poster_hor',
    "poster_ver"             varchar(80)   NOT NULL DEFAULT 'poster_ver',
    "end_year"               integer                DEFAULT NULL,
    "count_seasons"          integer                DEFAULT NULL,
    "rating"                 numeric(3, 1) NOT NULL DEFAULT 0,
    "count_scores"           integer       NOT NULL DEFAULT 0,
    "count_negative_reviews" integer       NOT NULL DEFAULT 0,
    "count_neutral_reviews"  integer       NOT NULL DEFAULT 0,
    "count_positive_reviews" integer       NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS genres
(
    "genre_id" serial      NOT NULL PRIMARY KEY,
    "name"     varchar(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS countries
(
    "country_id" serial      NOT NULL PRIMARY KEY,
    "name"       varchar(80) NOT NULL
);

CREATE TABLE IF NOT EXISTS companies
(
    "company_id" serial      NOT NULL PRIMARY KEY,
    "name"       varchar(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS persons
(
    "person_id"   serial       NOT NULL PRIMARY KEY,
    "name"        varchar(128) NOT NULL,
    "birthday"    date         NOT NULL,
    "death"       date         NOT NULL,
    "gender"      varchar(64)  NOT NULL DEFAULT 'male',
    "count_films" integer      NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS professions
(
    "profession_id" serial      NOT NULL PRIMARY KEY,
    "name"          varchar(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS collections
(
    "collection_id" serial       NOT NULL PRIMARY KEY,
    "name"          varchar(128) NOT NULL,
    "description"   TEXT         NOT NULL,
    "poster"        varchar(80)  NOT NULL DEFAULT 'poster',
    "is_public"     boolean      NOT NULL,
    "create_time"   timestamp    NOT NULL DEFAULT NOW(),
    "count_likes"   integer      NOT NULL DEFAULT 0,
    "count_films"   integer      NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS reviews
(
    "review_id"   serial        NOT NULL PRIMARY KEY,
    "name"        VARCHAR(80)   NOT NULL,
    "score"       NUMERIC(3, 1) NOT NULL,
    "description" TEXT          NOT NULL,
    "count_likes" integer       NOT NULL DEFAULT 0,
    "create_time" TIMESTAMP     NOT NULL DEFAULT NOW()
);


-- Зависимые таблицы 1:M
CREATE TABLE IF NOT EXISTS film_images
(
    "id"          serial  NOT NULL PRIMARY KEY,
    "fk_film_id"  integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "images_list" serial  NOT NULL
);

CREATE TABLE IF NOT EXISTS person_images
(
    "id"           serial  NOT NULL PRIMARY KEY,
    "fk_person_id" integer NOT NULL REFERENCES persons (person_id) ON DELETE CASCADE,
    "images_list"  serial  NOT NULL
);

-- Зависимые таблицы N:M
CREATE TABLE IF NOT EXISTS film_genres
(
    "fk_film_id"  integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "fk_genre_id" integer NOT NULL REFERENCES genres (genre_id) ON DELETE CASCADE,
    "weight"      integer NOT NULL,
    PRIMARY KEY (fk_film_id, fk_genre_id)
);

CREATE TABLE IF NOT EXISTS film_countries
(
    "fk_film_id"    integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "fk_country_id" integer NOT NULL REFERENCES countries (country_id) ON DELETE CASCADE,
    "weight"        integer NOT NULL,
    PRIMARY KEY (fk_film_id, fk_country_id)
);

CREATE TABLE IF NOT EXISTS film_companies
(
    "fk_film_id"    integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "fk_company_id" integer NOT NULL REFERENCES companies (company_id) ON DELETE CASCADE,
    "weight"        integer NOT NULL,
    PRIMARY KEY (fk_film_id, fk_company_id)
);

CREATE TABLE IF NOT EXISTS film_persons
(
    "fk_person_id"     integer     NOT NULL REFERENCES persons (person_id) ON DELETE CASCADE,
    "fk_film_id"       integer     NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "fk_profession_id" integer     NOT NULL REFERENCES professions (profession_id) ON DELETE CASCADE,
    "character"        varchar(80) NOT NULL,
    "weight"           integer     NOT NULL,
    PRIMARY KEY (fk_person_id, fk_film_id, fk_profession_id)
);

CREATE TABLE IF NOT EXISTS profile_ratings
(
    "fk_profile_id" integer       NOT NULL REFERENCES profiles (profile_id) ON DELETE CASCADE,
    "fk_film_id"    integer       NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "score"         NUMERIC(3, 1) NOT NULL,
    "create_date"   date          NOT NULL DEFAULT NOW(),
    PRIMARY KEY (fk_profile_id, fk_film_id)
);

CREATE TABLE IF NOT EXISTS profile_views_films
(
    "fk_profile_id" integer NOT NULL REFERENCES profiles (profile_id) ON DELETE CASCADE,
    "fk_film_id"    integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "create_date"   date    NOT NULL DEFAULT NOW(),
    PRIMARY KEY (fk_profile_id, fk_film_id)
);

CREATE TABLE IF NOT EXISTS profile_reviews
(
    "fk_review_id"  integer NOT NULL REFERENCES reviews (review_id) ON DELETE CASCADE,
    "fk_profile_id" integer NOT NULL REFERENCES profiles (profile_id) ON DELETE CASCADE,
    "fk_film_id"    integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    PRIMARY KEY (fk_review_id, fk_profile_id, fk_film_id)
);

CREATE TABLE IF NOT EXISTS profile_collections
(
    "fk_profile_id"    integer NOT NULL REFERENCES profiles (profile_id) ON DELETE CASCADE,
    "fk_collection_id" integer NOT NULL REFERENCES collections (collection_id) ON DELETE CASCADE,
    PRIMARY KEY (fk_profile_id, fk_collection_id)
);

CREATE TABLE IF NOT EXISTS person_professions
(
    "fk_person_id"     integer NOT NULL REFERENCES persons (person_id) ON DELETE CASCADE,
    "fk_profession_id" integer NOT NULL REFERENCES professions (profession_id) ON DELETE CASCADE,
    "weight"           integer NOT NULL,
    PRIMARY KEY (fk_person_id, fk_profession_id)
);

CREATE TABLE IF NOT EXISTS person_genres
(
    "fk_person_id" integer NOT NULL REFERENCES persons (person_id) ON DELETE CASCADE,
    "fk_genre_id"  integer NOT NULL REFERENCES genres (genre_id) ON DELETE CASCADE,
    "weight"       integer NOT NULL,
    PRIMARY KEY (fk_person_id, fk_genre_id)
);

CREATE TABLE IF NOT EXISTS collections_genres
(
    "fk_collection_id" integer NOT NULL REFERENCES collections (collection_id) ON DELETE CASCADE,
    "fk_genre_id"      integer NOT NULL REFERENCES genres (genre_id) ON DELETE CASCADE,
    "weight"           integer NOT NULL,
    PRIMARY KEY (fk_collection_id, fk_genre_id)
);

CREATE TABLE IF NOT EXISTS collection_likes
(
    "fk_profile_id"    integer NOT NULL REFERENCES profiles (profile_id) ON DELETE CASCADE,
    "fk_collection_id" integer NOT NULL REFERENCES collections (collection_id) ON DELETE CASCADE,
    "create_date"      date    NOT NULL DEFAULT NOW(),
    PRIMARY KEY (fk_profile_id, fk_collection_id)
);

CREATE TABLE IF NOT EXISTS collections_films
(
    "fk_film_id"       integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "fk_collection_id" integer NOT NULL REFERENCES collections (collection_id) ON DELETE CASCADE,
    PRIMARY KEY (fk_film_id, fk_collection_id)
);

CREATE TABLE IF NOT EXISTS reviews_likes
(
    "fk_review_id"  integer NOT NULL REFERENCES reviews (review_id) ON DELETE CASCADE,
    "fk_profile_id" integer NOT NULL REFERENCES profiles (profile_id) ON DELETE CASCADE,
    "create_date"   date    NOT NULL DEFAULT NOW(),
    PRIMARY KEY (fk_review_id, fk_profile_id)
);
