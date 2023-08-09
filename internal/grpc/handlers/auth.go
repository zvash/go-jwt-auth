package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zvash/go-jwt-auth/internal/config"
	"github.com/zvash/go-jwt-auth/internal/grpc/auth"
	"golang.org/x/net/context"
	"log"
	"strconv"
)

var grpcConfig *config.GRPCConfig

func SetGRPCConfig(c *config.GRPCConfig) {
	grpcConfig = c
}

type Server struct {
}

func (s *Server) Authenticate(ctx context.Context, message *auth.AuthRequestWithAccessToken) (*auth.User, error) {
	log.Printf("Received authenticate message from client: %v", message.Token)
	data, err := grpcConfig.DB.GetDataByValidAccessToken(ctx, message.Token)
	if err != nil {
		customError := errors.New(fmt.Sprintf("fetch token error: %v", err))
		return nil, customError
	}
	var dataMap map[string]string
	json.Unmarshal(data, &dataMap)
	userId, err := strconv.Atoi(dataMap["UserID"])
	if err != nil {
		return nil, err
	}
	dbUser, err := grpcConfig.DB.GetUserById(ctx, int32(userId))
	if err != nil {
		customError := errors.New(fmt.Sprintf("fetch user error: %v", err))
		return nil, customError
	}
	user := auth.User{
		Id:    dataMap["UserID"],
		Name:  dbUser.Name,
		Email: dbUser.Email,
	}
	return &user, nil
}
