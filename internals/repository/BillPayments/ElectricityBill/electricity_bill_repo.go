package electricitybillfetch_repository

import (
	"database/sql"
	"encoding/json"

	electricitybillfetch_domain "github.com/paybazar-backend/internals/domain/BillPayments/ElectricityBill"
)

type ElectricityBillFetchRepo struct {
	DB *sql.DB
}

func NewElectricityBillFetchRepo(db *sql.DB) *ElectricityBillFetchRepo {
	return &ElectricityBillFetchRepo{DB: db}
}

func (r *ElectricityBillFetchRepo) SaveElectricityBillFetch(req *electricitybillfetch_domain.ElectricityBillFetchRequest, res *electricitybillfetch_domain.ElectricityBillFetchResponse) error {
	data, _ := json.Marshal(res)
	_, err := r.DB.Exec(`
		INSERT INTO electricity_bill_fetch_logs (
			operator_code, consumer_id, api_error, api_status, api_msg, 
			consumer_name, bill_amount, bill_due_date, bill_date, api_response
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	`,
		req.OperatorCode, req.ConsumerID, res.Error, res.Status, res.Msg,
		res.BillAmount.ConsumerName, res.BillAmount.BillAmount, res.BillAmount.BillDueDate, res.BillAmount.BillDate, data,
	)
	return err
}
