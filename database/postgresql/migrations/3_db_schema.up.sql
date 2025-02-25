BEGIN;

CREATE TABLE IF NOT EXISTS users
(
    id                       BIGSERIAL PRIMARY KEY,
    name                     VARCHAR(255),
    tg_id                    BIGINT UNIQUE,
    username                 VARCHAR(255),
    balance                  BIGINT,
    updated_at               TIMESTAMP WITHOUT TIME ZONE,
    created_at               TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);

COMMIT;

