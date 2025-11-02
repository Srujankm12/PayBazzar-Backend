package ott_domain

type OTTPlanRequest struct {
	OperatorCode int `json:"operator_code"`
}

type OTTPlan struct {
	Amount      string `json:"amount"`
	Duration    int    `json:"duration"`
	Description string `json:"description"`
	PlanID      string `json:"plan_id"`
}

type OTTPlanResponse struct {
	Error        int        `json:"error"`
	Msg          string     `json:"msg"`
	RechargePlan []OTTPlan  `json:"recharge_plan"`
}
