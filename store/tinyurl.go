package store

import (
	"context"
	"time"

	"github.com/abs2free/tinyurl/store/model"
	"github.com/pkg/errors"
)

type TinyURLRepository interface {
	AddTinyURL(ctx context.Context, longURL, shortURL string, expireAT time.Time) (bool, error)
	BatchAddTinyURL(ctx context.Context, urls []model.TinyURL) (bool, error)
}

func (s *Mysql) AddTinyURL(ctx context.Context, longURL, shortURL string, expireAT time.Time) (bool, error) {
	var url = model.TinyURL{
		LongURL:   longURL,
		ShortURL:  shortURL,
		ExpiredAt: expireAT,
		CreatedAt: time.Now(),
	}
	result := s.db.Create(&url)
	if result.Error != nil {
		return false, errors.WithStack(result.Error)
	}
	return true, nil
}

func (s *Mysql) BatchAddTinyURL(ctx context.Context, urls []model.TinyURL) (bool, error) {
	result := s.db.Create(&urls)
	if result.Error != nil {
		return false, errors.WithStack(result.Error)
	}
	return true, nil
}
