CREATE TABLE IF NOT EXISTS user
(
    "user_id"      serial      NOT NULL PRIMARY KEY,
    "nickname"     varchar(64) NOT NULL,
    "email"        varchar(64) NOT NULL,
    "password"     text        NOT NULL,
    "joined_date"  DATE        NOT NULL,
    "is_superuser" BOOLEAN     NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS profile
(
    "profile_id" serial      NOT NULL PRIMARY KEY,
    "fk_user_id" integer     NOT NULL,
    "avatar"     varchar(80) NOT NULL DEFAULT 'default'
);

CREATE TABLE IF NOT EXISTS cinema
(
    "cinema_id"              serial       NOT NULL PRIMARY KEY,
    "name"                   varchar(128) NOT NULL,
    "description"            TEXT         NOT NULL,
    "short_description"      TEXT         NOT NULL,
    "type"                   varchar(64)  NOT NULL DEFAULT 'film',
    "prod_date"              DATE         NOT NULL,
    "end_date"               DATE         NOT NULL,
    "prod_company"           varchar(64)  NOT NULL,
    "prod_country"           varchar(64)  NOT NULL,
    "age_limit"              integer      NOT NULL,
    "duration"               TIME         NOT NULL,
    "poster_hor"             varchar(80)  NOT NULL DEFAULT 'default',
    "poster_ver"             varchar(80)  NOT NULL DEFAULT 'default',
    "rating"                 FLOAT        NOT NULL,
    "count_scores"           integer      NOT NULL,
    "count_reviews"          integer      NOT NULL,
    "count_negative_reviews" integer      NOT NULL,
    "count_neutral_reviews"  integer      NOT NULL,
    "count_positive_reviews" integer      NOT NULL
);

CREATE TABLE IF NOT EXISTS genre
(
    "genre_id" serial      NOT NULL PRIMARY KEY,
    "name"     varchar(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS person
(
    "person_id"     serial       NOT NULL PRIMARY KEY,
    "name"          varchar(128) NOT NULL,
    "birth_date"    DATE         NOT NULL,
    "count_cinemas" integer      NOT NULL
);

CREATE TABLE IF NOT EXISTS profession
(
    "profession_id" serial      NOT NULL PRIMARY KEY,
    "name"          varchar(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS collection
(
    "collection_id"     serial       NOT NULL PRIMARY KEY,
    "name"              varchar(128) NOT NULL,
    "description"       TEXT         NOT NULL,
    "short_description" TEXT         NOT NULL,
    "type"              varchar(64)  NOT NULL,
    "date_interval"     TEXT         NOT NULL,
    "count_cinemas"     integer      NOT NULL,
    "sum_duration"      TIME         NOT NULL,
    "age_limit"         integer      NOT NULL,
    "poster_ver"        varchar(80)  NOT NULL DEFAULT 'default',
    "poster_hor"        varchar(80)  NOT NULL DEFAULT 'default',
    "is_public"         BOOLEAN      NOT NULL
);


-- Зависимые таблицы
CREATE TABLE IF NOT EXISTS cinema_genres
(
    "id"           serial  NOT NULL PRIMARY KEY,
    "fk_cinema_id" integer NOT NULL REFERENCES cinema (cinema_id) ON DELETE CASCADE,
    "fk_genre_id"  integer NOT NULL REFERENCES genre (genre_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS cinema_images
(
    "id"           serial  NOT NULL PRIMARY KEY,
    "fk_cinema_id" integer NOT NULL REFERENCES cinema (cinema_id) ON DELETE CASCADE,
    "images_list"  serial  NOT NULL
);

CREATE TABLE IF NOT EXISTS profile_ratings
(
    "id"            serial  NOT NULL PRIMARY KEY,
    "fk_profile_id" integer NOT NULL REFERENCES profile (profile_id) ON DELETE CASCADE,
    "fk_cinema_id"  integer NOT NULL REFERENCES cinema (cinema_id) ON DELETE CASCADE,
    "score"         FLOAT   NOT NULL
);

CREATE TABLE IF NOT EXISTS profile_reviews
(
    "id"            serial  NOT NULL PRIMARY KEY,
    "fk_profile_id" integer NOT NULL REFERENCES profile (profile_id) ON DELETE CASCADE,
    "fk_cinema_id"  integer NOT NULL REFERENCES cinema (cinema_id) ON DELETE CASCADE,
    "score"         FLOAT   NOT NULL,
    "description"   TEXT    NOT NULL
);

CREATE TABLE IF NOT EXISTS profile_collections
(
    "id"               serial  NOT NULL PRIMARY KEY,
    "fk_profile_id"    integer NOT NULL REFERENCES profile (profile_id) ON DELETE CASCADE,
    "fk_collection_id" integer NOT NULL REFERENCES collection (collection_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS person_images
(
    "id"           serial  NOT NULL PRIMARY KEY,
    "fk_person_id" integer NOT NULL REFERENCES person (person_id) ON DELETE CASCADE,
    "images_list"  serial  NOT NULL
);

CREATE TABLE IF NOT EXISTS person_professions
(
    "id"               serial  NOT NULL PRIMARY KEY,
    "fk_person_id"     integer NOT NULL REFERENCES person (person_id) ON DELETE CASCADE,
    "fk_profession_id" integer NOT NULL REFERENCES profession (profession_id) ON DELETE CASCADE,
);

CREATE TABLE IF NOT EXISTS person_genres
(
    "id"           serial  NOT NULL PRIMARY KEY,
    "fk_person_id" integer NOT NULL REFERENCES person (person_id) ON DELETE CASCADE,
    "fk_genre_id"  integer NOT NULL REFERENCES genre (genre_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS collections_genres
(
    "id"               serial  NOT NULL PRIMARY KEY,
    "fk_collection_id" integer NOT NULL REFERENCES collection (collection_id) ON DELETE CASCADE,
    "fk_genre_id"      integer NOT NULL REFERENCES genre (genre_id) ON DELETE CASCADE
);
