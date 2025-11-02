package ott_repository

import (
	"database/sql"
	"encoding/json"

	ott_domain "github.com/paybazar-backend/internals/domain/Ott"
)

type OTTSubscriptionRepo struct {
	DB *sql.DB
}

func NewOTTSubscriptionRepo(db *sql.DB) *OTTSubscriptionRepo {
	return &OTTSubscriptionRepo{DB: db}
}

func (r *OTTSubscriptionRepo) SaveOTTSubscription(req *ott_domain.OTTSubscriptionRequest, res *ott_domain.OTTSubscriptionResponse) error {
	data, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		INSERT INTO ott_subscriptions (
			mobile_no, operator_code, amount, plan_id, customer_email, 
			partner_request_id, user_var1, user_var2, user_var3, 
			api_error, api_msg, api_status, order_id, optransid, commission, api_response
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	`,
		req.MobileNo, req.OperatorCode, req.Amount, req.PlanID, req.CustomerEmail,
		req.PartnerRequestID, req.UserVar1, req.UserVar2, req.UserVar3,
		res.Error, res.Msg, res.Status, res.OrderID, res.OpTransID, res.Commission, data,
	)
	return err
}
