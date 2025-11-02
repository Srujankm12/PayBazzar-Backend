package electricitybillfetch_domain

type ElectricityBillFetchRequest struct {
	ConsumerID   string `json:"consumer_id"`
	OperatorCode int    `json:"operator_code"`
}

type ElectricityBillDetails struct {
	ConsumerID   string  `json:"consumer_id"`
	ConsumerName string  `json:"consumer_name"`
	BillAmount   float64 `json:"bill_Amount"`
	BillDueDate  string  `json:"bill_due_date"`
	BillDate     string  `json:"bill_date"`
	BillNo       string  `json:"bill_no,omitempty"`
}

type ElectricityBillFetchResponse struct {
	Error      int                     `json:"error"`
	Status     int                     `json:"status"`
	Msg        string                  `json:"msg"`
	BillAmount ElectricityBillDetails  `json:"billAmount"`
}
