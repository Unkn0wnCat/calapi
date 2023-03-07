package auth

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/Unkn0wnCat/calapi/internal/logger"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

func ChallengeQuery(ctx context.Context) error {
	if viper.GetBool("auth.anonymous_read") == true || viper.GetString("auth.type") == AuthTypeNone {
		return nil // Anonymous querying is allowed. Anyone is allowed.
	}

	user := ForContext(ctx)
	if user == nil {
		logger.Logger.Warn("unauthorized query attempt",
			zap.String("requestId", middleware.GetReqID(ctx)),
			zap.String("gqlPath", graphql.GetPath(ctx).String()),
		)

		graphql.AddError(ctx, &gqlerror.Error{
			Message: "A login token is required, but was not provided.",
			Path:    graphql.GetPath(ctx),
		})

		return errors.New("no user found")
	}

	if user.ID == "ANON-NO-AUTH" {
		// This login was done when auth was turned off.
		logger.Logger.Warn("anonymous query attempt",
			zap.String("requestId", middleware.GetReqID(ctx)),
			zap.String("gqlPath", graphql.GetPath(ctx).String()),
		)

		graphql.AddError(ctx, &gqlerror.Error{
			Message: "The provided login token was anonymous, but this was since disabled. Please reauthenticate.",
			Path:    graphql.GetPath(ctx),
		})

		return errors.New("anonymous auth disabled")
	}

	return nil // User is set.
}

func ChallengeMutation(ctx context.Context) error {
	if viper.GetString("auth.type") == AuthTypeNone {
		return nil // Anonymous mutations are allowed. Anyone is allowed.
	}

	user := ForContext(ctx)
	if user == nil {
		logger.Logger.Warn("unauthorized mutation attempt",
			zap.String("requestId", middleware.GetReqID(ctx)),
			zap.String("gqlPath", graphql.GetPath(ctx).String()),
		)

		graphql.AddError(ctx, &gqlerror.Error{
			Message: "A login token is required, but was not provided.",
			Path:    graphql.GetPath(ctx),
		})

		return errors.New("no user found")
	}

	if user.ID == "ANON-NO-AUTH" {
		// This login was done when auth was turned off.
		logger.Logger.Warn("anonymous mutation attempt",
			zap.String("requestId", middleware.GetReqID(ctx)),
			zap.String("gqlPath", graphql.GetPath(ctx).String()),
		)

		graphql.AddError(ctx, &gqlerror.Error{
			Message: "The provided login token was anonymous, but this was since disabled. Please reauthenticate.",
			Path:    graphql.GetPath(ctx),
		})

		return errors.New("anonymous auth disabled")
	}

	return nil // User is set.
}
