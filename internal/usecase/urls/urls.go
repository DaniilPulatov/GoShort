package urls

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"time"
	"url-shortener/internal/domain/entites"
	"url-shortener/internal/shortening"
)

const shortUrlLifeTime = time.Hour * 24 * 7

// Shorten method creates token from a long url and send DTO to url repo
func (s *service) Shorten(ctx context.Context, input *entites.InputUrl) (string, error) {
	id := uuid.New().ID()
	if input.Identifier == "" {
		input.Identifier = shortening.Shorten(id)
	}

	oldRecord, err := s.urlRepo.GetByUrl(ctx, input.RealUrl)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return "", err
	}
	// if the oldRecord is not nil, check is link expired or not.
	// if expired then delete to create a new one
	if oldRecord != nil {
		if !time.Now().Local().Before(oldRecord.ExpiresAt) {
			log.Printf("URL already expired: %s\n", input.RealUrl)
			if err := s.urlRepo.Delete(ctx, oldRecord.RealUrl); err != nil {
				return "", err
			}
		} else {
			log.Println("Shorten url already exists")
			return "", errors.New("url already exists")
		}
	}
	shortUrl := &entites.ShortenUrl{
		RealUrl:    input.RealUrl,
		Identifier: input.Identifier,
		ID:         id,
		ExpiresAt:  time.Now().Local().Add(shortUrlLifeTime),
	}
	if err := s.urlRepo.Create(ctx, shortUrl); err != nil {
		return "", err
	}
	newUrl, err := shortening.AddBaseUrl(os.Getenv("BASE_URL"), input.Identifier)
	if err != nil {
		log.Println("Error adding base url", err)
		return "", err
	}
	return newUrl, nil
}

// Redirect method extract send obj with required url to handler and increment number of times the link was used
func (s *service) Redirect(ctx context.Context, identifier string) (*entites.ShortenUrl, error) {
	urlObj, err := s.urlRepo.GetByIdentifier(ctx, identifier)
	if err != nil {
		log.Println("Error Redirect(): ", err)
		return nil, err
	}
	if urlObj == nil {
		log.Println("Error Redirect(): token not found")
		return nil, errors.New("token not found")
	}
	log.Println(urlObj == nil)
	urlObj.Usages++
	if err := s.urlRepo.UpdateUsage(ctx, urlObj); err != nil {
		return nil, err
	}
	return urlObj, nil
}
