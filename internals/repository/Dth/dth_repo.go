package dth_repository

import (
	"database/sql"
	"encoding/json"

	dth_domain "github.com/paybazar-backend/internals/domain/Dth"
)

type DTHRechargeRepo struct {
	DB *sql.DB
}

func NewDTHRechargeRepo(db *sql.DB) *DTHRechargeRepo {
	return &DTHRechargeRepo{DB: db}
}

func (r *DTHRechargeRepo) SaveDTHRecharge(req *dth_domain.DTHRechargeRequest, res *dth_domain.DTHRechargeResponse) error {
	data, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		INSERT INTO dth_recharge_transactions 
			(customer_id, operator_code, amount, partner_request_id, dest, user_var1, user_var2, user_var3, api_error, api_msg, api_status, order_id, optransid, commission, api_response)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	`,
		req.CustomerID, req.OperatorCode, req.Amount, req.PartnerRequestID, req.Dest,
		req.UserVar1, req.UserVar2, req.UserVar3,
		res.Error, res.Msg, res.Status, res.OrderID, res.OpTransID, res.Commission, data,
	)
	return err
}
