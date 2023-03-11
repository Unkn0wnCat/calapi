package server

import (
	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Unkn0wnCat/calapi/graph"
	"github.com/Unkn0wnCat/calapi/internal/auth"
	"github.com/Unkn0wnCat/calapi/internal/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

const defaultPort = "8080"

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedResponse := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		latencyStart := time.Now()

		defer func() {
			latency := time.Since(latencyStart)

			logger.Logger.Info("HTTP request finished",
				zap.String("proto", r.Proto),
				zap.String("uri", r.RequestURI),
				zap.String("path", r.URL.Path),
				zap.String("method", r.Method),
				zap.String("remote", r.RemoteAddr),
				zap.Int("status", wrappedResponse.Status()),
				zap.Int("size", wrappedResponse.BytesWritten()),
				zap.Duration("latency", latency),
				zap.String("requestId", middleware.GetReqID(r.Context())),
			)
		}()

		next.ServeHTTP(wrappedResponse, r)
	})
}

func Serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	m := chiprometheus.NewMiddleware("calapi")
	router.Use(m)
	router.Use(middleware.RequestID)
	router.Use(logMiddleware)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		//ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.Handle("/metrics", promhttp.Handler())

	logger.Sugar.Infof("Now serving at http://localhost:%s/", port)
	logger.Sugar.Fatal(http.ListenAndServe(":"+port, router))

}
