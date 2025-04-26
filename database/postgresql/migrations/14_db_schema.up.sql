BEGIN;

CREATE TABLE IF NOT EXISTS users
(
    id                       BIGSERIAL PRIMARY KEY,
    tg_id                    BIGINT UNIQUE,
    username                 VARCHAR(255) UNIQUE,
    wallet                   VARCHAR(255) UNIQUE,
    balance                  BIGINT default 0,
    updated_at               TIMESTAMP WITHOUT TIME ZONE,
    created_at               TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);


CREATE TABLE IF NOT EXISTS cards
(
    id            BIGSERIAL PRIMARY KEY,
    name          VARCHAR(100) NOT NULL UNIQUE,
    description   TEXT,
    max_level     BIGINT,
    updated_at               TIMESTAMP WITHOUT TIME ZONE,
    created_at               TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);

CREATE TABLE IF NOT EXISTS users_cards 
(
    user_id        BIGINT REFERENCES users(id) ON DELETE CASCADE,
    card_id        BIGINT REFERENCES cards(id) ON DELETE CASCADE,
    card_name      VARCHAR(100) NOT NULL,
    current_level  BIGINT DEFAULT 0,
    auto_clicker   BOOLEAN DEFAULT FALSE,
    updated_at     TIMESTAMP WITHOUT TIME ZONE,
    created_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    PRIMARY KEY (user_id, card_id)
);

CREATE TABLE IF NOT EXISTS cost_cards
(
    card_id        BIGINT REFERENCES cards(id) ON DELETE CASCADE,
    card_name      VARCHAR(100) NOT NULL UNIQUE,
    base_price     BIGINT NOT NULL default 100,
    growth_coefficient FLOAT NOT NULL,
    updated_at     TIMESTAMP WITHOUT TIME ZONE,
    created_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    PRIMARY KEY (card_id)
);

CREATE TABLE IF NOT EXISTS background_pictures
(
    pictures_id     BIGSERIAL PRIMARY KEY,
    url             VARCHAR(1000),
    updated_at               TIMESTAMP WITHOUT TIME ZONE,
    created_at               TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);

CREATE TABLE IF NOT EXISTS users_background_pictures
(
    user_id        BIGINT REFERENCES users(id) ON DELETE CASCADE,
    pictures_id    BIGINT REFERENCES background_pictures(pictures_id) ON DELETE CASCADE,
    is_active       BOOLEAN DEFAULT FALSE,
    created_at               TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);


COMMIT;

