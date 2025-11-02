package moneytransfer_domain

type CreateWalletRequest struct {
	MobileNo string `json:"mobile_no"`
}

type CreateWalletResponse struct {
	Error       int    `json:"error"`
	Msg         string `json:"msg"`
	MobileNo    string `json:"MobileNo"`
	RequestNo   string `json:"RequestNo"`
	Description string `json:"description"`
	UserVar1    string `json:"user_var1"`
	UserVar2    string `json:"user_var2"`
	UserVar3    string `json:"user_var3"`
}


type VerifyOtpRequest struct {
	MobileNo     string `json:"mobile_no"`
	RequestNo    string `json:"request_no"`
	Otp          string `json:"otp"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	City         string `json:"city"`
	State        string `json:"state"`
	PinCode      string `json:"pinCode"`
}

type VerifyOtpResponse struct {
	Error       int    `json:"error"`
	Msg         string `json:"msg"`
	MobileNo    string `json:"MobileNo"`
	Description string `json:"description"`
	UserVar1    string `json:"user_var1"`
	UserVar2    string `json:"user_var2"`
	UserVar3    string `json:"user_var3"`
}
