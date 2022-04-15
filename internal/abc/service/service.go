package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/zhenisduissekov/dummy_service/internal/abc/client"
	"github.com/zhenisduissekov/dummy_service/internal/abc/repository"
)

type Service struct {
	Client     *client.Client
	Repository *repository.Repository
}

type RequestRegisterItems struct {
	Id             string `json:"id" validate:"required" db:"id"`
	ReturnUrl      string `json:"returnUrl" validate:"required" db:"returnUrl"`
	FailUrl        string `json:"failUrl" validate:"required" db:"failUrl"`
	BillNumber     string `json:"billNumber" validate:"required" db:"billNumber"`
	ExpirationDate string `json:"expirationDate"`
}

func New(db *sqlx.DB, addr, user, pass, proxy string) *Service {
	return &Service{
		Client:     client.New(addr, user, pass, proxy),
		Repository: repository.New(db),
	}
}

func (s *Service) Register(items RequestRegisterItems) (string, error) {
	//TODO: fill it up
	return "ok", nil
}
