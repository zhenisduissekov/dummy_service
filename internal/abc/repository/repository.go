package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var (
	SaveQuery = `INSERT INTO orders (abc_order_id, def_id) VALUES($1, $2) RETURNING id`
)

type Repository struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db,
	}
}

func (db *Repository) Save(OrderId, defId string) (int, error) {
	var id int
	row := db.DB.QueryRow(SaveQuery, OrderId, defId)
	if err := row.Scan(&id); err != nil {
		log.Err(err).Msg("error at row.Scan")
		return 0, err
	}
	return id, nil
}
