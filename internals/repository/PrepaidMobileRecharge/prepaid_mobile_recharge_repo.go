package prepaidmobilerecharge_repository

import (
	"database/sql"
	"encoding/json"

	prepaidmobilerecharge_domain "github.com/paybazar-backend/internals/domain/PrepaidMobileRecharge"
)

type RechargeRepo struct {
	DB *sql.DB
}

func NewRechargeRepo(db *sql.DB) *RechargeRepo {
	return &RechargeRepo{DB: db}
}

func (r *RechargeRepo) SaveRecharge(req *prepaidmobilerecharge_domain.RechargeRequest, res *prepaidmobilerecharge_domain.RechargeResponse) error {
	resJSON, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		INSERT INTO recharge_transactions (
			mobile_no, operator_code, amount, partner_request_id, circle,
			recharge_type, dest, user_var1, user_var2, user_var3,
			api_error, api_msg, api_status, order_id, optransid, commission, api_response
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)
	`,
		req.MobileNo, req.OperatorCode, req.Amount, req.PartnerRequestID, req.Circle,
		req.RechargeType, req.Dest, req.UserVar1, req.UserVar2, req.UserVar3,
		res.Error, res.Msg, res.Status, res.OrderID, res.OpTransID, res.Commission, resJSON,
	)
	return err
}
func (r *RechargeRepo) GetAllRecharges() ([]prepaidmobilerecharge_domain.RechargeResponse, error) {
	rows, err := r.DB.Query(`
		SELECT 
			api_error, api_msg, api_status, order_id, optransid, partner_request_id, commission,
			user_var1, user_var2, user_var3
		FROM recharge_transactions
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recharges []prepaidmobilerecharge_domain.RechargeResponse
	for rows.Next() {
		var rec prepaidmobilerecharge_domain.RechargeResponse
		if err := rows.Scan(
			&rec.Error, &rec.Msg, &rec.Status, &rec.OrderID, &rec.OpTransID,
			&rec.PartnerReq, &rec.Commission, &rec.UserVar1, &rec.UserVar2, &rec.UserVar3,
		); err != nil {
			return nil, err
		}
		recharges = append(recharges, rec)
	}
	return recharges, nil
}

func (r *RechargeRepo) GetRechargeByPartnerReqID(partnerReqID string) (*prepaidmobilerecharge_domain.RechargeResponse, error) {
	row := r.DB.QueryRow(`
		SELECT 
			api_error, api_msg, api_status, order_id, optransid, partner_request_id, commission,
			user_var1, user_var2, user_var3
		FROM recharge_transactions
		WHERE partner_request_id = $1
	`, partnerReqID)

	var rec prepaidmobilerecharge_domain.RechargeResponse
	err := row.Scan(
		&rec.Error, &rec.Msg, &rec.Status, &rec.OrderID, &rec.OpTransID,
		&rec.PartnerReq, &rec.Commission, &rec.UserVar1, &rec.UserVar2, &rec.UserVar3,
	)
	if err != nil {
		return nil, err
	}
	return &rec, nil
}
