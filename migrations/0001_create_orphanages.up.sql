CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS orphanages (
    id               uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name             VARCHAR(50) NOT NULL,
    latitude         FLOAT NOT NULL,
    longitude        FLOAT NOT NULL,
    about            TEXT NOT NULL,
    instructions     TEXT NOT NULL,
    opened_hours     VARCHAR(10) NOT NULL,
    open_on_weekends BOOLEAN NOT NULL,
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP NULL,
    deleted_at       TIMESTAMP NULL
);
