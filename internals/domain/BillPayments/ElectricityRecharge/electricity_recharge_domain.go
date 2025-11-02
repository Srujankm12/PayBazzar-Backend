package electricitybillpayment_domain

type ElectricityBillPaymentRequest struct {
	P1               string  `json:"p1"`
	P2               string  `json:"p2"`
	P3               string  `json:"p3"`
	CustomerEmail    string  `json:"customer_email"`
	OperatorCode     int     `json:"operator_code"`
	Amount           float64 `json:"amount"`
	PartnerRequestID string  `json:"partner_request_id"`
	UserVar1         string  `json:"user_var1"`
	UserVar2         string  `json:"user_var2"`
	UserVar3         string  `json:"user_var3"`
}

type ElectricityBillPaymentResponse struct {
	Error         int     `json:"error"`
	Msg           string  `json:"msg"`
	Status        int     `json:"status"`
	OrderID       string  `json:"orderid"`
	OpTransID     string  `json:"optransid"`
	PartnerReqID  string  `json:"partnerreqid"`
	Commission    float64 `json:"commission"`
	UserVar1      string  `json:"user_var1"`
	UserVar2      string  `json:"user_var2"`
	UserVar3      string  `json:"user_var3"`
}
