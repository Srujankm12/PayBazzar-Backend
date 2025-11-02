package moneytransfer_domain

type MRTransferRequest struct {
	MobileNo        string  `json:"mobile_no"`
	BeneficiaryName string  `json:"beneficiary_name"`
	BeneficiaryCode string  `json:"beneficiary_code"`
	PartnerRequestID string `json:"partner_request_id"`
	Amount          float64 `json:"amount"`
	AccountNo       string  `json:"account_no"`
	BankName        string  `json:"bank_name"`
	IFSC            string  `json:"ifsc"`
	TransferType    string  `json:"transfer_type"` // '5' -> IMPS, '6' -> NEFT
	UserVar1        string  `json:"user_var1,omitempty"`
	UserVar2        string  `json:"user_var2,omitempty"`
	UserVar3        string  `json:"user_var3,omitempty"`
}

type MRTransferResponse struct {
	Error      int     `json:"error"`
	Msg        string  `json:"msg"`
	Status     int     `json:"status"`
	OrderID    string  `json:"orderid"`
	OpTransID  string  `json:"optransid"`
	PartnerReq string  `json:"partnerreqid"`
	Commission float64 `json:"commission"`
	UserVar1   string  `json:"user_var1"`
	UserVar2   string  `json:"user_var2"`
	UserVar3   string  `json:"user_var3"`
}

type TransferStatusRequest struct {
	TransactionID string `json:"transaction_id"`
}

type TransferStatusResponse struct {
	Error         int    `json:"error"`
	Msg           string `json:"msg"`
	Status        string `json:"status"`
	TransactionId string `json:"TransactionId"`
	UserVar1      string `json:"user_var1"`
	UserVar2      string `json:"user_var2"`
	UserVar3      string `json:"user_var3"`
}
