-- Generator completed
CREATE TABLE IF NOT EXISTS users
(
    "user_id"     serial      NOT NULL PRIMARY KEY,
    "nickname"    varchar(64) NOT NULL UNIQUE,
    "email"       varchar(64) NOT NULL UNIQUE,
    "password"    bytea       NOT NULL,
    "is_admin"    boolean     NOT NULL DEFAULT false,
    "last_update" date        NOT NULL DEFAULT NOW()
);

-- Generator completed
CREATE TABLE IF NOT EXISTS profiles
(
    "profile_id"        serial NOT NULL PRIMARY KEY REFERENCES users (user_id) ON DELETE CASCADE,
    "avatar"            varchar(32) DEFAULT NULL,
    "joined_date"       date        DEFAULT NOW(),
    "count_views_films" integer     DEFAULT NULL,
    "count_collections" integer     DEFAULT NULL,
    "count_reviews"     integer     DEFAULT NULL,
    "count_ratings"     integer     DEFAULT NULL
);

-- Generator completed
CREATE TABLE IF NOT EXISTS films
(
    "film_id"                serial       NOT NULL PRIMARY KEY,
    "name"                   varchar(80)  NOT NULL,
    "prod_year"              integer      NOT NULL,
    "description"            TEXT         NOT NULL,
    "short_description"      varchar(180) NOT NULL,
    "duration"               integer      NOT NULL,
    "type"                   varchar(32)  DEFAULT NULL,
    "original_name"          varchar(80)  DEFAULT NULL,
    "slogan"                 varchar(128) DEFAULT NULL,
    "age_limit"              integer      DEFAULT NULL,
    "budget"                 integer      DEFAULT NULL,
    "box_office"             integer      DEFAULT NULL,
    "currency_budget"        varchar(8)   DEFAULT NULL,
    "poster_hor"             varchar(32)  DEFAULT NULL,
    "poster_ver"             varchar(32)  DEFAULT NULL,
    "end_year"               integer      DEFAULT NULL,
    "count_seasons"          integer      DEFAULT NULL,
    "rating"                 real         DEFAULT NULL,
    "count_actors"           integer      DEFAULT NULL,
    "count_scores"           integer      DEFAULT NULL,
    "count_negative_reviews" integer      DEFAULT NULL,
    "count_neutral_reviews"  integer      DEFAULT NULL,
    "count_positive_reviews" integer      DEFAULT NULL,
    "update_time"            timestamp    DEFAULT NOW()
);

-- Generator completed
CREATE TABLE IF NOT EXISTS genres
(
    "genre_id" serial      NOT NULL PRIMARY KEY,
    "name"     varchar(64) NOT NULL UNIQUE
);

-- Generator completed
CREATE TABLE IF NOT EXISTS countries
(
    "country_id" serial      NOT NULL PRIMARY KEY,
    "name"       varchar(64) NOT NULL UNIQUE
);

-- Generator completed
CREATE TABLE IF NOT EXISTS companies
(
    "company_id" serial      NOT NULL PRIMARY KEY,
    "name"       varchar(64) NOT NULL UNIQUE
);

-- Generator completed
CREATE TABLE IF NOT EXISTS tags
(
    "tag_id" serial      NOT NULL PRIMARY KEY,
    "name"   varchar(64) NOT NULL UNIQUE
);

-- Generator completed
CREATE TABLE IF NOT EXISTS persons
(
    "person_id"     serial        NOT NULL PRIMARY KEY,
    "name"          varchar(128)  NOT NULL,
    "birthday"      date          NOT NULL,
    "growth"        numeric(3, 2) NOT NULL,
    "original_name" varchar(80) DEFAULT NULL,
    "avatar"        varchar(16) DEFAULT NULL,
    "death"         date        DEFAULT NULL,
    "gender"        varchar(16) DEFAULT NULL,
    "count_films"   integer     DEFAULT NULL
);

-- Generator completed
CREATE TABLE IF NOT EXISTS professions
(
    "profession_id" serial      NOT NULL PRIMARY KEY,
    "name"          varchar(64) NOT NULL
);

-- Generator completed
CREATE TABLE IF NOT EXISTS collections
(
    "collection_id" serial      NOT NULL PRIMARY KEY,
    "name"          varchar(64) NOT NULL,
    "description"   TEXT                 DEFAULT NULL,
    "poster"        varchar(32)          DEFAULT NULL,
    "is_public"     boolean     NOT NULL DEFAULT false,
    "create_time"   timestamp   NOT NULL DEFAULT NOW(),
    "count_likes"   integer              DEFAULT NULL,
    "count_films"   integer              DEFAULT NULL
);

-- Generator completed
CREATE TABLE IF NOT EXISTS reviews
(
    "review_id"   serial      NOT NULL PRIMARY KEY,
    "name"        VARCHAR(64) NOT NULL,
    "type"        VARCHAR(16) NOT NULL,
    "body"        TEXT        NOT NULL,
    "count_likes" integer              DEFAULT NULL,
    "create_time" TIMESTAMP   NOT NULL DEFAULT NOW()
);


-- 1:M
-- Generator completed
CREATE TABLE IF NOT EXISTS film_images
(
    "film_id"     serial NOT NULL PRIMARY KEY REFERENCES films (film_id) ON DELETE CASCADE,
    "images_list" varchar(100) DEFAULT NULL
);

-- Generator completed
CREATE TABLE IF NOT EXISTS person_images
(
    "person_id"   serial       NOT NULL PRIMARY KEY REFERENCES persons (person_id) ON DELETE CASCADE,
    "images_list" varchar(100) NOT NULL
);

--     1 1
--     1 2
--     1 3
--     1 4

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
    "character"        varchar(64) DEFAULT NULL,
    "weight"           integer NOT NULL,
    PRIMARY KEY (fk_person_id, fk_film_id, fk_profession_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS profile_ratings
(
    "fk_profile_id" integer       NOT NULL REFERENCES profiles (profile_id) ON DELETE CASCADE,
    "fk_film_id"    integer       NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "score"         NUMERIC(3, 1) NOT NULL,
    "create_date"   date          NOT NULL DEFAULT NOW(),
    PRIMARY KEY (fk_profile_id, fk_film_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS profile_views_films
(
    "fk_profile_id" integer NOT NULL REFERENCES profiles (profile_id) ON DELETE CASCADE,
    "fk_film_id"    integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    "create_date"   date    NOT NULL DEFAULT NOW(),
    PRIMARY KEY (fk_profile_id, fk_film_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS profile_reviews
(
    "fk_review_id"  integer NOT NULL REFERENCES reviews (review_id) ON DELETE CASCADE,
    "fk_profile_id" integer NOT NULL REFERENCES profiles (profile_id) ON DELETE CASCADE,
    "fk_film_id"    integer NOT NULL REFERENCES films (film_id) ON DELETE CASCADE,
    PRIMARY KEY (fk_review_id, fk_profile_id, fk_film_id)
);

-- Generator completed
CREATE TABLE IF NOT EXISTS profile_collections
(
    "fk_profile_id"    integer NOT NULL REFERENCES profiles (profile_id) ON DELETE CASCADE,
    "fk_collection_id" integer NOT NULL REFERENCES collections (collection_id) ON DELETE CASCADE,
    PRIMARY KEY (fk_profile_id, fk_collection_id)
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

-- Generator completed
CREATE TABLE IF NOT EXISTS reviews_likes
(
    "fk_review_id"  integer NOT NULL REFERENCES reviews (review_id) ON DELETE CASCADE,
    "fk_profile_id" integer NOT NULL REFERENCES profiles (profile_id) ON DELETE CASCADE,
    "create_date"   date    NOT NULL DEFAULT NOW(),
    PRIMARY KEY (fk_review_id, fk_profile_id)
);
