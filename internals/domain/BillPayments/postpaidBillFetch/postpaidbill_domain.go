package postpaidbill_domain

type PostpaidBillRequest struct {
	OperatorCode int    `json:"operator_code"`
	MobileNo     string `json:"mobile_no"`
}

type BillAmountDetails struct {
	BillAmount     string `json:"billAmount"`
	BillNetAmount  string `json:"billnetamount"`
	BillDate       string `json:"billdate"`
	DueDate        string `json:"dueDate"`
	AcceptPayment  string `json:"acceptPayment"`
	AcceptPartPay  string `json:"acceptPartPay"`
	CellNumber     string `json:"cellNumber"`
	UserName       string `json:"userName"`
}

type PostpaidBillResponse struct {
	Error      int               `json:"error"`
	Status     int               `json:"status"`
	Msg        string            `json:"msg"`
	BillAmount BillAmountDetails `json:"billAmount"`
}
