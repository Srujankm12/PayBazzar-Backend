package moneytransfer_repository

import (
	"database/sql"
	"encoding/json"

	moneytransfer_domain "github.com/paybazar-backend/internals/domain/MoneyTransfer"
)


type WalletCreateRepo struct {
	DB *sql.DB
}

func NewWalletCreateRepo(db *sql.DB) *WalletCreateRepo {
	return &WalletCreateRepo{DB: db}
}

func (r *WalletCreateRepo) SaveWalletRequest(req *moneytransfer_domain.CreateWalletRequest, res *moneytransfer_domain.CreateWalletResponse) error {
	data, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		INSERT INTO wallet_create_requests (
			mobile_no, api_error, api_msg, request_no, description, api_response
		) VALUES ($1, $2, $3, $4, $5, $6)
	`, req.MobileNo, res.Error, res.Msg, res.RequestNo, res.Description, data)
	return err
}

type WalletVerifyRepo struct {
	DB *sql.DB
}

func NewWalletVerifyRepo(db *sql.DB) *WalletVerifyRepo {
	return &WalletVerifyRepo{DB: db}
}

func (r *WalletVerifyRepo) SaveWalletVerification(req *moneytransfer_domain.VerifyOtpRequest, res *moneytransfer_domain.VerifyOtpResponse) error {
	data, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		INSERT INTO wallet_verify_logs (
			mobile_no, request_no, otp, first_name, last_name, address_line1, address_line2,
			city, state, pin_code, api_error, api_msg, description, api_response
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)
	`, req.MobileNo, req.RequestNo, req.Otp, req.FirstName, req.LastName, req.AddressLine1,
		req.AddressLine2, req.City, req.State, req.PinCode,
		res.Error, res.Msg, res.Description, data)
	return err
}
