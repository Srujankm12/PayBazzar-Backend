package postpaidbill_repository

import (
	"database/sql"
	"encoding/json"

	postpaidbill_domain "github.com/paybazar-backend/internals/domain/BillPayments/postpaidBillFetch"
)

type PostpaidBillFetchRepo struct {
	DB *sql.DB
}

func NewPostpaidBillFetchRepo(db *sql.DB) *PostpaidBillFetchRepo {
	return &PostpaidBillFetchRepo{DB: db}
}

func (r *PostpaidBillFetchRepo) SavePostpaidBillFetch(req *postpaidbill_domain.PostpaidBillRequest, res *postpaidbill_domain.PostpaidBillResponse) error {
	data, _ := json.Marshal(res)

	_, err := r.DB.Exec(`
		INSERT INTO postpaid_bill_fetch_logs (
			operator_code, mobile_no, api_error, api_status, api_msg, bill_amount, bill_due_date, bill_date, customer_name, bill_no, api_response
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
	`,
		req.OperatorCode,
		req.MobileNo,
		res.Error,
		res.Status,
		res.Msg,
		res.BillAmount.BillAmount,
		res.BillAmount.DueDate,
		res.BillAmount.BillDate,
		res.BillAmount.UserName,
		res.BillAmount.CellNumber,
		data,
	)
	return err
}
