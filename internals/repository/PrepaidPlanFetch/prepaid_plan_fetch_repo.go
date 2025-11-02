package prepaidplanfetch_repository

import (
	"database/sql"
	"encoding/json"

	prepaidplanfetch_domain "github.com/paybazar-backend/internals/domain/PrepaidPlanFetch"
)

type PrepaidPlanRepository struct {
	DB *sql.DB
}

func NewPrepaidPlanRepository(db *sql.DB) *PrepaidPlanRepository {
	return &PrepaidPlanRepository{DB: db}
}

func (r *PrepaidPlanRepository) SavePlan(req *prepaidplanfetch_domain.PrepaidPlanRequest, res *prepaidplanfetch_domain.PrepaidPlanResponse) error {
	data, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		INSERT INTO prepaid_plans (operator_code, circle, response_data)
		VALUES ($1, $2, $3)
	`, req.OperatorCode, req.Circle, data)
	return err
}
