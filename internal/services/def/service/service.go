package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/zhenisduissekov/dummy_service/internal/def/client"
	"github.com/zhenisduissekov/dummy_service/internal/def/repository"
)

type Service struct {
	Client     *client.Client
	Repository *repository.Repository
}

func New(db *sqlx.DB, url, user, pass string) *Service {
	return &Service{
		Client:     client.New(url, user, pass),
		Repository: repository.New(db),
	}
}
