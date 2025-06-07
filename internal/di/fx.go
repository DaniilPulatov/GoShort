package di

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"net"
	"net/http"
	"os"
	"time"
	"url-shortener/internal/migrations"
	urlRepo "url-shortener/internal/repository/urls"
	"url-shortener/internal/rest"
	urlHandler "url-shortener/internal/rest/handler/urls"
	urlSerivse "url-shortener/internal/usecase/urls"
	"url-shortener/pkg/postgresDB"
)

func NewMux() *gin.Engine {
	return gin.New()
}

func NewHTTPServer(lc fx.Lifecycle, server *rest.Server) *http.Server {
	srv := &http.Server{
		Addr:              net.JoinHostPort(os.Getenv("HOST"), os.Getenv("PORT")),
		Handler:           server,
		ReadHeaderTimeout: 10 * time.Second,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Printf("Error starting server: %s\n", err)
				}
			}()
			server.Init()
			log.Println("Server started")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down server")
			return srv.Close()
		},
	})
	return srv
}

func PostgresProvider(lc fx.Lifecycle) (postgresDB.Pool, error) {
	pool, err := postgresDB.NewPostgresDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("Error connecting to postgres", err)
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			pool.Close()
			return nil
		},
	})
	return pool, nil
}

func NewModule() fx.Option {
	return fx.Module(
		"url-shortener",
		fx.Provide(
			NewMux,
			PostgresProvider,

			fx.Annotate(
				urlRepo.NewUrlRepo,
				fx.As(new(urlRepo.UrlRepository)),
			),
			fx.Annotate(
				urlSerivse.NewUrlService,
				fx.As(new(urlSerivse.UrlService)),
			),
			fx.Annotate(
				urlHandler.NewUrlHandler,
				fx.As(new(urlHandler.UrlHandler)),
			),

			rest.NewServer,
			http.NewServeMux,
			NewHTTPServer,
		),
		fx.Invoke(
			func(lc fx.Lifecycle) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return migrations.ApplyMigrations(os.Getenv("MIGRATIONS_DIR"), os.Getenv("DATABASE_URL"))
					},
				})
			},
			func(*http.Server) {},
		),
	)
}
