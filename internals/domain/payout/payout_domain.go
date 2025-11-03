package payoutdomain

type PayoutRequest struct {
	MobileNo        string `json:"mobile_no"`
	AccountNo       string `json:"account_no"`
	IFSC            string `json:"ifsc"`
	BankName        string `json:"bank_name"`
	BeneficiaryName string `json:"beneficiary_name"`
	Amount          string `json:"amount"` // keep string if API expects string; convert/validate as needed
	TransferType    string `json:"transfer_type"`
	PartnerRequestID string `json:"partner_request_id"`
}

type PayoutResponse struct {
	Error      int    `json:"error"`
	Msg        string `json:"msg"`
	Status     int    `json:"status"`
	OrderID    string `json:"orderid"`
	OpTransID  string `json:"optransid"`
	PartnerReq string `json:"partnerreqid"`
}
