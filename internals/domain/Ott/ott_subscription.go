package ott_domain

type OTTSubscriptionRequest struct {
	MobileNo         string  `json:"mobile_no"`
	OperatorCode     int     `json:"operator_code"`
	Amount           float64 `json:"amount"`
	PlanID           string  `json:"plan_id,omitempty"`
	CustomerEmail    string  `json:"customer_email"`
	PartnerRequestID string  `json:"partner_request_id"`
	UserVar1         string  `json:"user_var1,omitempty"`
	UserVar2         string  `json:"user_var2,omitempty"`
	UserVar3         string  `json:"user_var3,omitempty"`
}

type OTTSubscriptionResponse struct {
	Error          int     `json:"error"`
	Msg            string  `json:"msg"`
	Status         int     `json:"status"`
	OrderID        string  `json:"orderid"`
	OpTransID      string  `json:"optransid"`
	PartnerReqID   string  `json:"partner_request_id"`
	Commission     float64 `json:"commission"`
	UserVar1       string  `json:"user_var1"`
	UserVar2       string  `json:"user_var2"`
	UserVar3       string  `json:"user_var3"`
}
