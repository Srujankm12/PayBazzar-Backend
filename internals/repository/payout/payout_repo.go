package payoutrepo

import (
	"database/sql"
	"encoding/json"
	"fmt"

	payoutdomain "github.com/paybazar-backend/internals/domain/payout"
)

type PayoutRepo struct {
	DB *sql.DB
}

func NewPayoutRepo(db *sql.DB) *PayoutRepo {
	return &PayoutRepo{DB: db}
}

// SavePayout saves the request and response (response can be nil on error)
func (r *PayoutRepo) SavePayout(req *payoutdomain.PayoutRequest, res *payoutdomain.PayoutResponse) error {
	resJSON := []byte("null")
	if res != nil {
		b, _ := json.Marshal(res)
		resJSON = b
	}

	_, err := r.DB.Exec(`
		INSERT INTO payout_transactions (
			mobile_no, account_no, ifsc, bank_name, beneficiary_name,
			amount, transfer_type, partner_request_id,
			api_error, api_msg, api_status, order_id, optransid, partnerreqid, api_response
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)
	`,
		req.MobileNo, req.AccountNo, req.IFSC, req.BankName, req.BeneficiaryName,
		req.Amount, req.TransferType, req.PartnerRequestID,
		nil, nil, nil, nil, nil, nil, resJSON,
	)
	if err != nil {
		return fmt.Errorf("save payout exec: %w", err)
	}
	return nil
}

// UpdatePayoutWithResponse updates the last inserted row matching partner_request_id (or other logic)
func (r *PayoutRepo) UpdatePayoutWithResponse(partnerReq string, res *payoutdomain.PayoutResponse) error {
	b, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		UPDATE payout_transactions
		SET api_error=$1, api_msg=$2, api_status=$3, order_id=$4, optransid=$5, partnerreqid=$6, api_response=$7
		WHERE partner_request_id=$8
	`,
		res.Error, res.Msg, res.Status, res.OrderID, res.OpTransID, res.PartnerReq, b, partnerReq,
	)
	if err != nil {
		return err
	}
	return nil
}
