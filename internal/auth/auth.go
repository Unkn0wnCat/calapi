package auth

import (
	"context"
	"errors"
	"github.com/Unkn0wnCat/calapi/internal/ghost"
	"github.com/Unkn0wnCat/calapi/internal/logger"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	AuthTypeGhost = "GHOST"
	AuthTypeNone  = "NONE"
)

type User struct {
	ID       string
	Username string
	Name     string
}

func Authenticate(ctx context.Context, username string, password string) (*User, error) {
	switch viper.GetString("auth.type") {
	case AuthTypeGhost:
		return AuthenticateGhost(ctx, username, password)
	case AuthTypeNone:
		logger.Logger.Info("anonymously authenticated",
			zap.String("requestId", middleware.GetReqID(ctx)),
		)
		return &User{
			ID:       "ANON-NO-AUTH",
			Username: "anonymous",
			Name:     "Anonymous User",
		}, nil
	default:
		return nil, errors.New("unknown authentication method")
	}
}

func AuthenticateGhost(ctx context.Context, username string, password string) (*User, error) {
	api := ghost.GhostAPI{
		BaseURL: viper.GetString("auth.ghost.base_url"),
		Jar:     nil,
	}

	_, err := api.Login(username, password)
	if err != nil {
		logger.Logger.Warn("invalid ghost credentials",
			zap.String("requestId", middleware.GetReqID(ctx)),
			zap.Error(err),
		)
		return nil, err
	}

	ghostUser, err := api.UserSelf()
	if err != nil {
		logger.Logger.Error("ghost error",
			zap.String("requestId", middleware.GetReqID(ctx)),
			zap.Error(err),
		)
		return nil, err
	}

	if len(ghostUser.Users) == 0 {
		logger.Logger.Error("unexpected empty response from Ghost API",
			zap.String("requestId", middleware.GetReqID(ctx)),
		)
		return nil, errors.New("unexpected empty response from Ghost API")
	}

	logger.Logger.Info("ghost authentication success",
		zap.String("requestId", middleware.GetReqID(ctx)),
		zap.String("userId", ghostUser.Users[0].ID),
		zap.String("username", ghostUser.Users[0].Email),
		zap.Error(err),
	)

	user := &User{
		ID:       ghostUser.Users[0].ID,
		Username: ghostUser.Users[0].Email,
		Name:     ghostUser.Users[0].Name,
	}

	return user, nil
}
