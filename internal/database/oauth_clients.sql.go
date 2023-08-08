// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: oauth_clients.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const getAllClients = `-- name: GetAllClients :many
SELECT id, client_id, secret, domain FROM oauth_clients
`

func (q *Queries) GetAllClients(ctx context.Context) ([]OauthClient, error) {
	rows, err := q.db.QueryContext(ctx, getAllClients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OauthClient
	for rows.Next() {
		var i OauthClient
		if err := rows.Scan(
			&i.ID,
			&i.ClientID,
			&i.Secret,
			&i.Domain,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getClientWithClientId = `-- name: GetClientWithClientId :one
SELECT id, client_id, secret, domain FROM oauth_clients WHERE client_id = $1 LIMIT 1
`

func (q *Queries) GetClientWithClientId(ctx context.Context, clientID uuid.UUID) (OauthClient, error) {
	row := q.db.QueryRowContext(ctx, getClientWithClientId, clientID)
	var i OauthClient
	err := row.Scan(
		&i.ID,
		&i.ClientID,
		&i.Secret,
		&i.Domain,
	)
	return i, err
}

const getClientWithIdAndSecret = `-- name: GetClientWithIdAndSecret :one
SELECT id, client_id, secret, domain FROM oauth_clients WHERE client_id = $1 AND secret = $2 LIMIT 1
`

type GetClientWithIdAndSecretParams struct {
	ClientID uuid.UUID
	Secret   string
}

func (q *Queries) GetClientWithIdAndSecret(ctx context.Context, arg GetClientWithIdAndSecretParams) (OauthClient, error) {
	row := q.db.QueryRowContext(ctx, getClientWithIdAndSecret, arg.ClientID, arg.Secret)
	var i OauthClient
	err := row.Scan(
		&i.ID,
		&i.ClientID,
		&i.Secret,
		&i.Domain,
	)
	return i, err
}