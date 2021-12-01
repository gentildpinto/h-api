CREATE TABLE IF NOT EXISTS images (
    id           uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    path         TEXT NOT NULL,
    orphanage_id uuid NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP NULL,
    deleted_at   TIMESTAMP NULL
);