package prepaidmobilerecharge_domain

type RechargeRequest struct {
	MobileNo         string  `json:"mobile_no"`
	OperatorCode     int     `json:"operator_code"`
	Amount           float64 `json:"amount"`
	PartnerRequestID string  `json:"partner_request_id"`
	Circle           int     `json:"circle"`
	RechargeType     int     `json:"recharge_type"`
	Dest             string  `json:"dest,omitempty"`
	UserVar1         string  `json:"user_var1,omitempty"`
	UserVar2         string  `json:"user_var2,omitempty"`
	UserVar3         string  `json:"user_var3,omitempty"`
}

type RechargeResponse struct {
	Error       int     `json:"error"`
	Msg         string  `json:"msg"`
	Status      int     `json:"status"`
	OrderID     string  `json:"orderid"`
	OpTransID   string  `json:"optransid"`
	PartnerReq  string  `json:"partnerreqid"`
	Commission  float64 `json:"commission"`
	UserVar1    string  `json:"user_var1"`
	UserVar2    string  `json:"user_var2"`
	UserVar3    string  `json:"user_var3"`
}
