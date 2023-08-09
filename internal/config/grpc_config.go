package config

import (
	"database/sql"
	"github.com/zvash/go-jwt-auth/internal/database"
	"google.golang.org/grpc"
)

type GRPCConfig struct {
	DB     *database.Queries
	Server *grpc.Server
}

func (app *GRPCConfig) SetDB(connection *sql.DB) {
	app.DB = database.New(connection)
}

func (app *GRPCConfig) SetServer(server *grpc.Server) {
	app.Server = server
}
