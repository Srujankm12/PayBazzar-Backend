package moneytransfer_repository

import (
	"database/sql"
	"encoding/json"

	moneytransfer_domain "github.com/paybazar-backend/internals/domain/MoneyTransfer"
)

type MoneyTransferRepository struct {
	DB *sql.DB
}

func NewMoneyTransferRepository(db *sql.DB) *MoneyTransferRepository {
	return &MoneyTransferRepository{DB: db}
}

// Insert a transfer attempt/transaction
func (r *MoneyTransferRepository) InsertMoneyTransfer(req *moneytransfer_domain.MRTransferRequest, res *moneytransfer_domain.MRTransferResponse) error {
	apiResp, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		INSERT INTO money_transfer_transactions (
			mobile_no, beneficiary_name, beneficiary_code, partner_request_id, amount,
			account_no, bank_name, ifsc, transfer_type,
			api_error, api_msg, api_status, order_id, optransid, partnerreqid,
			user_var1, user_var2, user_var3, api_response, created_at, updated_at
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,
			$10,$11,$12,$13,$14,$15,
			$16,$17,$18,$19,NOW(),NOW()
		)
	`,
		req.MobileNo, req.BeneficiaryName, req.BeneficiaryCode, req.PartnerRequestID, req.Amount,
		req.AccountNo, req.BankName, req.IFSC, req.TransferType,
		res.Error, res.Msg, res.Status, res.OrderID, res.OpTransID, res.PartnerReq,
		req.UserVar1, req.UserVar2, req.UserVar3, apiResp,
	)
	return err
}

// Insert transfer status check response
func (r *MoneyTransferRepository) InsertTransferStatusLog(req *moneytransfer_domain.TransferStatusRequest, res *moneytransfer_domain.TransferStatusResponse) error {
	apiResp, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		INSERT INTO money_transfer_status_logs (
			transaction_id, api_error, api_msg, api_status, transaction_status, api_response, created_at
		) VALUES ($1,$2,$3,$4,$5,$6,NOW())
	`,
		req.TransactionID, res.Error, res.Msg, res.Status, res.Status, apiResp,
	)
	return err
}
