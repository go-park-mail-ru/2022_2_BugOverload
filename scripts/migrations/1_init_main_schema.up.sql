-- Let's set it in migrations for now. It's better to install when deploying
set timezone = 'UTC-3';

-- Generator completed
CREATE TABLE IF NOT EXISTS users
(
    "user_id"           serial  NOT NULL PRIMARY KEY,
    "nickname"          text    NOT NULL,
    CONSTRAINT "nickname_length" CHECK (LENGTH("nickname") <= 64),
    "email"             text    NOT NULL UNIQUE,
    CONSTRAINT "email_length" CHECK (LENGTH("email") <= 64),
    "password"          bytea   NOT NULL,
    "is_admin"          boolean NOT NULL DEFAULT false,
    "updated_at"        date    NOT NULL DEFAULT NOW(),
    -- profile
    "avatar"            text             DEFAULT NULL,
    CONSTRAINT "avatar_length" CHECK (LENGTH("avatar") <= 64),
    "joined_date"       date             DEFAULT NOW(),
    -- Denormalize fields
    "count_views_films" integer NOT NULL DEFAULT 0,
    "count_collections" integer NOT NULL DEFAULT 0,
    "count_reviews"     integer NOT NULL DEFAULT 0,
    "count_ratings"     integer NOT NULL DEFAULT 0
);

CREATE TYPE film_type_enum AS ENUM ('serial');

CREATE TYPE currency_enum AS ENUM ('EURO');

CREATE TYPE age_limit_enum AS ENUM ('6+','12+','16+','18+','21+');

-- Generator completed
CREATE TABLE IF NOT EXISTS films
(
    "film_id"                serial    NOT NULL PRIMARY KEY,
    "name"                   text      NOT NULL,
    CONSTRAINT "name_length" CHECK (LENGTH("name") <= 80),
    "prod_year"              date      NOT NULL,
    "description"            text      NOT NULL,
    "short_description"      text               DEFAULT NULL,
    CONSTRAINT "short_description_length" CHECK (LENGTH("short_description") <= 180),
    "duration_minutes"       integer            DEFAULT NULL,
    "type"                   film_type_enum     DEFAULT NULL,
    "original_name"          text               DEFAULT NULL,
    CONSTRAINT "original_name_length" CHECK (LENGTH("original_name") <= 80),
    "slogan"                 text               DEFAULT NULL,
    CONSTRAINT "slogan_length" CHECK (LENGTH("slogan") <= 128),
    "age_limit"              age_limit_enum     DEFAULT NULL,
    "budget"                 integer            DEFAULT NULL,
    "box_office"             integer            DEFAULT NULL,
    "currency_budget"        currency_enum      DEFAULT NULL,
    "poster_hor"             text               DEFAULT NULL,
    CONSTRAINT "poster_hor_length" CHECK (LENGTH("poster_hor") <= 32),
    "poster_ver"             text               DEFAULT NULL,
    CONSTRAINT "poster_ver_length" CHECK (LENGTH("poster_ver") <= 32),
    "end_year"               date               DEFAULT NULL,
    "count_seasons"          integer            DEFAULT NULL,
    -- Denormalize fields
    "rating"                 real               DEFAULT NULL,
    "count_actors"           integer   NOT NULL DEFAULT 0,
    "count_scores"           integer   NOT NULL DEFAULT 0,
    "count_negative_reviews" integer   NOT NULL DEFAULT 0,
    "count_neutral_reviews"  integer   NOT NULL DEFAULT 0,
    "count_positive_reviews" integer   NOT NULL DEFAULT 0,
    "updated_at"             timestamp NOT NULL DEFAULT NOW()
);

-- Generator completed
CREATE TABLE IF NOT EXISTS genres
(
    "genre_id" serial NOT NULL PRIMARY KEY,
    "name"     text   NOT NULL UNIQUE,
    CONSTRAINT "name_length" CHECK (LENGTH("name") <= 64)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS countries
(
    "country_id" serial NOT NULL PRIMARY KEY,
    "name"       text   NOT NULL UNIQUE,
    CONSTRAINT "name_length" CHECK (LENGTH("name") <= 64)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS companies
(
    "company_id" serial NOT NULL PRIMARY KEY,
    "name"       text   NOT NULL UNIQUE,
    CONSTRAINT "name_length" CHECK (LENGTH("name") <= 64)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS tags
(
    "tag_id" serial NOT NULL PRIMARY KEY,
    "name"   text   NOT NULL UNIQUE,
    CONSTRAINT "name_length" CHECK (LENGTH("name") <= 64)
);

CREATE TYPE gender_enum AS ENUM ('female');

-- Generator completed
CREATE TABLE IF NOT EXISTS persons
(
    "person_id"     serial  NOT NULL PRIMARY KEY,
    "name"          text    NOT NULL,
    CONSTRAINT "name_length" CHECK (LENGTH("name") <= 128),
    "birthday"      date    NOT NULL,
    "growth_meters" numeric(3, 2)    DEFAULT NULL,
    "original_name" text             DEFAULT NULL,
    CONSTRAINT "original_name_length" CHECK (LENGTH("original_name") <= 80),
    "avatar"        text             DEFAULT NULL,
    CONSTRAINT "avatar_length" CHECK (LENGTH("avatar") <= 32),
    "death"         date             DEFAULT NULL,
    "gender"        gender_enum      DEFAULT NULL,
    -- Denormalize fields
    "count_films"   integer NOT NULL DEFAULT 0
);

-- Generator completed
CREATE TABLE IF NOT EXISTS professions
(
    "profession_id" serial NOT NULL PRIMARY KEY,
    "name"          text   NOT NULL,
    CONSTRAINT "name_length" CHECK (LENGTH("name") <= 64)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS collections
(
    "collection_id" serial    NOT NULL PRIMARY KEY,
    "name"          text      NOT NULL,
    CONSTRAINT "name_length" CHECK (LENGTH("name") <= 64),
    "description"   text               DEFAULT NULL,
    "poster"        text               DEFAULT NULL,
    CONSTRAINT "poster_length" CHECK (LENGTH("poster") <= 32),
    "is_public"     boolean   NOT NULL DEFAULT false,
    "create_time"   timestamp NOT NULL DEFAULT NOW(),
    -- Denormalize fields
    "count_likes"   integer   NOT NULL DEFAULT 0,
    "count_films"   integer   NOT NULL DEFAULT 0
);

CREATE TYPE type_review_enum AS ENUM ('positive', 'negative','neutral');

-- Generator completed
CREATE TABLE IF NOT EXISTS reviews
(
    "review_id"   serial           NOT NULL PRIMARY KEY,
    "name"        text             NOT NULL,
    CONSTRAINT "name_length" CHECK (LENGTH("name") <= 64),
    "type"        type_review_enum NOT NULL,
    "body"        text             NOT NULL,
    "count_likes" integer          NOT NULL DEFAULT 0,
    "create_time" timestamp        NOT NULL DEFAULT NOW()
);


-- 1:M
-- Generator completed
CREATE TABLE IF NOT EXISTS film_images
(
    "film_id"   serial  NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "image_key" text DEFAULT NULL,
    "weight"    integer NOT NULL,
    CONSTRAINT "image_length" CHECK (LENGTH(image_key) <= 32),
    PRIMARY KEY (image_key, film_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS person_images
(
    "person_id" serial  NOT NULL REFERENCES persons (person_id) ON DELETE CASCADE,
    "image_key" text DEFAULT NULL,
    "weight"    integer NOT NULL,
    CONSTRAINT "image_length" CHECK (LENGTH(image_key) <= 32),
    PRIMARY KEY (image_key, person_id)
);

-- N:M
-- Generator completed
CREATE TABLE IF NOT EXISTS film_genres
(
    "fk_film_id"  integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "fk_genre_id" integer NOT NULL REFERENCES genres (genre_id) ON DELETE CASCADE,
    "weight"      integer NOT NULL,
    PRIMARY KEY (fk_film_id, fk_genre_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS film_tags
(
    "fk_film_id" integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "fk_tag_id"  integer NOT NULL REFERENCES tags (tag_id) ON DELETE CASCADE,
    PRIMARY KEY (fk_film_id, fk_tag_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS film_countries
(
    "fk_film_id"    integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "fk_country_id" integer NOT NULL REFERENCES countries (country_id) ON DELETE CASCADE,
    "weight"        integer NOT NULL,
    PRIMARY KEY (fk_film_id, fk_country_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS film_companies
(
    "fk_film_id"    integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "fk_company_id" integer NOT NULL REFERENCES companies (company_id) ON DELETE CASCADE,
    "weight"        integer NOT NULL,
    PRIMARY KEY (fk_film_id, fk_company_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS film_persons
(
    "fk_person_id"     integer NOT NULL REFERENCES persons (person_id) ON DELETE CASCADE,
    "fk_film_id"       integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "fk_profession_id" integer NOT NULL REFERENCES professions (profession_id) ON DELETE CASCADE,
    "character"        text DEFAULT NULL,
    CONSTRAINT "character_length" CHECK (LENGTH("character") <= 64),
    "weight"           integer NOT NULL,
    PRIMARY KEY (fk_person_id, fk_film_id, fk_profession_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS profile_ratings
(
    "fk_user_id"  integer  NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    "fk_film_id"  integer  NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "score"       smallint NOT NULL,
    "create_date" date     NOT NULL DEFAULT NOW(),
    PRIMARY KEY (fk_user_id, fk_film_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS profile_views_films
(
    "fk_user_id"  integer NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    "fk_film_id"  integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "create_date" date    NOT NULL DEFAULT NOW(),
    PRIMARY KEY (fk_user_id, fk_film_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS profile_reviews
(
    "fk_review_id" integer NOT NULL REFERENCES reviews (review_id) ON DELETE CASCADE,
    "fk_user_id"   integer NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    "fk_film_id"   integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    PRIMARY KEY (fk_review_id, fk_user_id, fk_film_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS profile_collections
(
    "fk_user_id"       integer NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    "fk_collection_id" integer NOT NULL REFERENCES collections (collection_id) ON DELETE CASCADE,
    PRIMARY KEY (fk_user_id, fk_collection_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS person_professions
(
    "fk_person_id"     integer NOT NULL REFERENCES persons (person_id) ON DELETE CASCADE,
    "fk_profession_id" integer NOT NULL REFERENCES professions (profession_id) ON DELETE CASCADE,
    "weight"           integer NOT NULL,
    PRIMARY KEY (fk_person_id, fk_profession_id)
);

-- Generator completed
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
    "fk_user_id"       integer NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    "fk_collection_id" integer NOT NULL REFERENCES collections (collection_id) ON DELETE CASCADE,
    "create_date"      date    NOT NULL DEFAULT NOW(),
    PRIMARY KEY (fk_user_id, fk_collection_id)
);

CREATE TABLE IF NOT EXISTS collections_films
(
    "fk_film_id"       integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "fk_collection_id" integer NOT NULL REFERENCES collections (collection_id) ON DELETE CASCADE,
    PRIMARY KEY (fk_film_id, fk_collection_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS reviews_likes
(
    "fk_review_id" integer NOT NULL REFERENCES reviews (review_id) ON DELETE CASCADE,
    "fk_user_id"   integer NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    "create_date"  date    NOT NULL DEFAULT NOW(),
    PRIMARY KEY (fk_review_id, fk_user_id)
);
