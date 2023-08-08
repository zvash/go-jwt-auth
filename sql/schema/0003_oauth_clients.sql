-- +goose Up
CREATE TABLE oauth_clients (
    id SERIAL PRIMARY KEY,
    client_id uuid NOT NULL UNIQUE,
    secret text NOT NULL,
    domain text default 'http://localhost' NOT NULL
);

-- +goose Down
DROP TABLE oauth_clients;