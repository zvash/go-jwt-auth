-- name: GetDataByValidAccessToken :one
SELECT "data" FROM oauth2_tokens WHERE access = $1 AND expires_at > NOW() LIMIT 1;