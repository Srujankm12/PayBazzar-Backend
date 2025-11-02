package ott_repository

import (
	"database/sql"
	"encoding/json"

	ott_domain "github.com/paybazar-backend/internals/domain/Ott"
)

type OTTPlanRepo struct {
	DB *sql.DB
}

func NewOTTPlanRepo(db *sql.DB) *OTTPlanRepo {
	return &OTTPlanRepo{DB: db}
}

func (r *OTTPlanRepo) SaveOTTPlan(req *ott_domain.OTTPlanRequest, res *ott_domain.OTTPlanResponse) error {
	data, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		INSERT INTO ott_plans (operator_code, response_data)
		VALUES ($1, $2)
	`, req.OperatorCode, data)
	return err
}
