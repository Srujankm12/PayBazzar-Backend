package moneytransfer_repository

import (
	"database/sql"
	"encoding/json"

	moneytransfer_domain "github.com/paybazar-backend/internals/domain/MoneyTransfer"
)

type BeneficiaryRepository struct {
	DB *sql.DB
}

func NewBeneficiaryRepository(db *sql.DB) *BeneficiaryRepository {
	return &BeneficiaryRepository{DB: db}
}

// ------------------ ADD BENEFICIARY ------------------
func (r *BeneficiaryRepository) InsertAddBeneficiary(req *moneytransfer_domain.AddBeneficiaryRequest, res *moneytransfer_domain.AddBeneficiaryResponse) error {
	data, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		INSERT INTO beneficiary_add_requests (
			mobile_no, beneficiary_name, bank_name, account_no, ifsc, 
			api_error, api_msg, beneficiary_id, description, api_response
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	`,
		req.MobileNo, req.BeneficiaryName, req.BankName, req.AccountNo, req.IFSC,
		res.Error, res.Msg, res.BeneficiaryID, res.Description, data,
	)
	return err
}

// ------------------ GET BENEFICIARY ------------------
func (r *BeneficiaryRepository) InsertBeneficiaryList(mobileNo string, res *moneytransfer_domain.GetBeneficiaryResponse) error {
	data, _ := json.Marshal(res)
	for _, b := range res.BeneficiaryList {
		_, _ = r.DB.Exec(`
			INSERT INTO beneficiary_list_logs (
				mobile_no, api_error, api_msg, description,
				beneficiary_id, uuid, account_no, ifsc_code, bank_name, account_holder_name, api_response
			) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		`,
			mobileNo, res.Error, res.Msg, res.Description,
			b.BeneficiaryID, b.UUID, b.AccountDetail.AccountNumber,
			b.AccountDetail.IFSCCode, b.AccountDetail.BankName,
			b.AccountDetail.AccountHolderName, data,
		)
	}
	return nil
}

// ------------------ DELETE BENEFICIARY REQUEST ------------------
func (r *BeneficiaryRepository) InsertDeleteBeneficiary(req *moneytransfer_domain.DeleteBeneficiaryRequest, res *moneytransfer_domain.DeleteBeneficiaryResponse) error {
	data, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		INSERT INTO beneficiary_delete_requests (
			mobile_no, beneficiary_id, api_error, api_msg, request_no, description, api_response
		) VALUES ($1,$2,$3,$4,$5,$6,$7)
	`,
		req.MobileNo, req.BeneficiaryID, res.Error, res.Msg, res.RequestNo, res.Description, data,
	)
	return err
}

// ------------------ VERIFY DELETE BENEFICIARY ------------------
func (r *BeneficiaryRepository) InsertVerifyDeleteBeneficiary(req *moneytransfer_domain.VerifyDeleteBeneficiaryRequest, res *moneytransfer_domain.VerifyDeleteBeneficiaryResponse) error {
	data, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		INSERT INTO beneficiary_delete_verifications (
			mobile_no, request_no, otp, beneficiary_id, api_error, api_msg, description, api_response
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`,
		req.MobileNo, req.RequestNo, req.OTP, req.BeneficiaryID,
		res.Error, res.Msg, res.Description, data,
	)
	return err
}
