package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/zhenisduissekov/dummy_service/internal/abc/client"
	"github.com/zhenisduissekov/dummy_service/internal/abc/repository"
)

type Service struct {
	Client     *client.Client
	Repository *repository.Repository
}

type RequestRegisterItems struct {
	Id             string `json:"id" validate:"required" db:"id" example:"2022-01-01""`
	ReturnUrl      string `json:"returnUrl" validate:"required" db:"returnUrl"  example:"web.com/success"`
	FailUrl        string `json:"failUrl" validate:"required" db:"failUrl" example:"web.com/fail"`
	BillNumber     string `json:"billNumber" validate:"required" db:"billNumber" example:"12345678"`
	ExpirationDate string `json:"expirationDate"`
}

func New(db *sqlx.DB, url, user, pass, proxy string) *Service {
	return &Service{
		Client:     client.New(url, user, pass, proxy),
		Repository: repository.New(db),
	}
}

func (s *Service) Register(items RequestRegisterItems) (map[string]int, error) {
	counter := 0
	data, err := s.Client.SendRequest()
	if err != nil {
		log.Err(err).Msg("error at MakeRequest")
		return map[string]int{}, err
	}

	for {
		_, err := s.Repository.Save(items.Id, data[counter].NodeId)
		if err != nil {
			log.Err(err).Msg("error at Repository.Save")
			return map[string]int{"received items": len(data), "saved item": counter}, err
		}
		counter += 1

		if counter >= len(data) {
			return map[string]int{"received items": len(data), "saved item": counter}, nil
		}
	}
}
