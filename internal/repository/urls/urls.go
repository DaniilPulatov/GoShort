package urls

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
	"url-shortener/internal/domain/entites"
)

func (r *repo) Create(ctx context.Context, url *entites.ShortenUrl) error {
	var (
		insertQuery = `insert into urls (id, real_url, identifier, expires_at) values ($1, $2, $3, $4) returning id;`
		id          int
	)
	if err := r.pool.QueryRow(ctx, insertQuery, url.ID, url.RealUrl, url.Identifier, url.ExpiresAt).Scan(&id); err != nil {
		log.Println("Error Create() inserting new url:", err)
		return err
	}
	log.Printf("Created new short url with id: %d\n", id)
	return nil
}

func (r *repo) UpdateUsage(ctx context.Context, url *entites.ShortenUrl) error {
	var (
		incrementQuery = `update urls set usages=$1 where identifier=$2;`
		usages         int
	)
	if _, err := r.pool.Exec(ctx, incrementQuery, url.Usages, url.Identifier); err != nil {
		log.Println("Error UpdateUsage() updating urls:", err)
		return err
	}
	log.Printf("Usages for short url %v: %d\n", url.Identifier, usages)
	return nil
}

func (r *repo) GetByUrl(ctx context.Context, url string) (*entites.ShortenUrl, error) {
	var (
		res         entites.ShortenUrl
		selectQuery = `select id, real_url, identifier, usages, created_at, expires_at from urls where real_url = $1;`
	)
	if err := r.pool.QueryRow(ctx, selectQuery, url).Scan(&res.ID, &res.RealUrl,
		&res.Identifier, &res.Usages,
		&res.CreatedAt, &res.ExpiresAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("Error GetByUrl():", err)
			return nil, err
		}
		log.Println("Error GetByUrl():", err)
		return nil, err
	}
	return &res, nil
}

func (r *repo) GetByIdentifier(ctx context.Context, identifier string) (*entites.ShortenUrl, error) {
	var (
		res         entites.ShortenUrl
		selectQuery = `select id, real_url, identifier, usages, created_at, expires_at from urls where identifier = $1;`
	)
	if err := r.pool.QueryRow(ctx, selectQuery, identifier).Scan(&res.ID, &res.RealUrl,
		&res.Identifier, &res.Usages,
		&res.CreatedAt, &res.ExpiresAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("Error GetByIdentifier():", err)
			return nil, err
		}
		log.Println("Error GetByIdentifier():", err)
		return nil, err
	}
	return &res, nil
}

func (r *repo) Delete(ctx context.Context, realUrl string) error {
	deleteQuery := `delete from urls where real_url = $1;`
	if _, err := r.pool.Exec(ctx, deleteQuery, realUrl); err != nil {
		log.Println("Error Delete():", err)
		return err
	}
	log.Printf("Deleted url with token: %s\n", realUrl)
	return nil
}
