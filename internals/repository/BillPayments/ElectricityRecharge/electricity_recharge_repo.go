package electricitybillpayment_repository

import (
	"database/sql"
	"encoding/json"

	electricitybillpayment_domain "github.com/paybazar-backend/internals/domain/BillPayments/ElectricityRecharge"
)

type ElectricityBillPaymentRepo struct {
	DB *sql.DB
}

func NewElectricityBillPaymentRepo(db *sql.DB) *ElectricityBillPaymentRepo {
	return &ElectricityBillPaymentRepo{DB: db}
}

func (r *ElectricityBillPaymentRepo) SavePayment(req *electricitybillpayment_domain.ElectricityBillPaymentRequest, res *electricitybillpayment_domain.ElectricityBillPaymentResponse) error {
	data, _ := json.Marshal(res)

	_, err := r.DB.Exec(`
		INSERT INTO electricity_bill_payment_transactions (
			p1, p2, p3, customer_email, operator_code, amount, partner_request_id,
			api_error, api_msg, api_status, order_id, optransid, partnerreqid, commission,
			user_var1, user_var2, user_var3, api_response
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)
	`,
		req.P1, req.P2, req.P3, req.CustomerEmail, req.OperatorCode, req.Amount, req.PartnerRequestID,
		res.Error, res.Msg, res.Status, res.OrderID, res.OpTransID, res.PartnerReqID, res.Commission,
		req.UserVar1, req.UserVar2, req.UserVar3, data,
	)
	return err
}
