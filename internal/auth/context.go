package auth

import (
	"context"
	"github.com/Unkn0wnCat/calapi/internal/logger"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
			authScheme := auth[0]

			if len(auth) != 2 || (!strings.EqualFold(authScheme, "bearer") && !strings.EqualFold(authScheme, "jwt")) {
				// No authentication provided
				next.ServeHTTP(w, r)
				return
			}
			authPayload := auth[1]

			user, err := ParseJWT(authPayload)
			if err != nil {
				logger.Logger.Warn("invalid token provided",
					zap.String("requestId", middleware.GetReqID(r.Context())),
					zap.Error(err),
				)
				http.Error(w, "Invalid Token", http.StatusUnauthorized)
				return
			}

			logger.Logger.Info("token validated",
				zap.String("requestId", middleware.GetReqID(r.Context())),
				zap.String("userId", user.ID),
				zap.String("username", user.Username),
			)

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}
