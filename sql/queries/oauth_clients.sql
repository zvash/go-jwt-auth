-- name: GetClientWithClientId :one
SELECT * FROM oauth_clients WHERE client_id = $1 LIMIT 1;

-- name: GetClientWithIdAndSecret :one
SELECT * FROM oauth_clients WHERE client_id = $1 AND secret = $2 LIMIT 1;

-- name: GetAllClients :many
SELECT * FROM oauth_clients;