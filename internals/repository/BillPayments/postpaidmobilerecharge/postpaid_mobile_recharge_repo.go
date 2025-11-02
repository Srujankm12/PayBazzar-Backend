package postpaidmobilerecharge_repository

import (
	"database/sql"
	"encoding/json"

	postpaidmobilerecharge_domain "github.com/paybazar-backend/internals/domain/BillPayments/postpaidmobilerecharge"
)

type PostpaidMobileRechargeRepo struct {
	DB *sql.DB
}

func NewPostpaidMobileRechargeRepo(db *sql.DB) *PostpaidMobileRechargeRepo {
	return &PostpaidMobileRechargeRepo{DB: db}
}

func (r *PostpaidMobileRechargeRepo) SavePostpaidRecharge(req *postpaidmobilerecharge_domain.PostpaidMobileRechargeRequest, res *postpaidmobilerecharge_domain.PostpaidMobileRechargeResponse) error {
	data, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		INSERT INTO postpaid_mobile_recharge_transactions (
			mobile_no, operator_code, amount, partner_request_id, circle, recharge_type,
			user_var1, user_var2, user_var3,
			api_error, api_msg, api_status, order_id, optransid, commission, api_response
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)
	`,
		req.MobileNo, req.OperatorCode, req.Amount, req.PartnerRequestID, req.Circle, req.RechargeType,
		req.UserVar1, req.UserVar2, req.UserVar3,
		res.Error, res.Msg, res.Status, res.OrderID, res.OpTransID, res.Commission, data,
	)
	return err
}
