package main

import (
	"github.com/labstack/echo/v4"
	electricitybillfetch_handlers "github.com/paybazar-backend/internals/handlers/BillPayments/ElectricityBill"
	electricitybillpayment_handlers "github.com/paybazar-backend/internals/handlers/BillPayments/ElectricityRecharge"
	postpaidbill_handlers "github.com/paybazar-backend/internals/handlers/BillPayments/postpaidBillFetch"
	postpaidmobilerecharge_handlers "github.com/paybazar-backend/internals/handlers/BillPayments/postpaidmobilerecharge"
	dth_handlers "github.com/paybazar-backend/internals/handlers/Dth"
	moneytransfer_handlers "github.com/paybazar-backend/internals/handlers/MoneyTransfer"
	ott_handlers "github.com/paybazar-backend/internals/handlers/Ott"
	prepaidmobilerecharge_handlers "github.com/paybazar-backend/internals/handlers/PrepaidMobileRecharge"
	prepaidplanfetch_handlers "github.com/paybazar-backend/internals/handlers/PrepaidPlanFetch"
	payouthandlers "github.com/paybazar-backend/internals/handlers/payout"
)

func SetupRouter(
	e *echo.Echo,
	MobilerechargeHandler *prepaidmobilerecharge_handlers.RechargeHandler,
	prepaidPlanHandler *prepaidplanfetch_handlers.PrepaidPlanHandler,
	dthHandler *dth_handlers.DTHRechargeHandler,
	ottPlanHandler *ott_handlers.OTTPlanHandler,
	ottSubscriptionHandler *ott_handlers.OTTSubscriptionHandler,
	PostpaidBillFetchHandler *postpaidbill_handlers.PostpaidBillFetchHandler,
	PostpaidMobileRechargeHandler *postpaidmobilerecharge_handlers.PostpaidRechargeHandler,
	ElectricityBillFetchHandler *electricitybillfetch_handlers.ElectricityBillFetchHandler,
	ElectricityBillPaymentHandler *electricitybillpayment_handlers.ElectricityBillPaymentHandler,
	WalletCreateHandler *moneytransfer_handlers.WalletHandler,
	BeneficiaryHandler *moneytransfer_handlers.BeneficiaryHandler,
	MoneyTransferHandler *moneytransfer_handlers.MoneyTransferHandler,
	PayoutHandler *payouthandlers.PayoutHandler,
	
	
) {
	prepaidmobileRecharge := e.Group("/recharge")
	prepaidmobileRecharge.POST("", MobilerechargeHandler.Recharge)
	prepaidmobileRecharge.GET("/:partner_request_id", MobilerechargeHandler.GetRechargeByPartnerReqID)
	prepaidmobileRecharge.GET("/get/all", MobilerechargeHandler.GetAllRecharges)

	prepaidplanfetch := e.Group("/prepaid")
	prepaidplanfetch.POST("/store/plan", prepaidPlanHandler.CreatePrepaidPlan)
	prepaidplanfetch.GET("/plans", prepaidPlanHandler.GetPrepaidPlan)

	dth := e.Group("/dth")
	dth.POST("/recharge", dthHandler.RechargeDTH)
	dth.GET("/get/recharge", dthHandler.GetDTHRecharge)

	ott := e.Group("/ott")
	ott.GET("/plans", ottPlanHandler.FetchOTTPlans)

	ottSubscription := e.Group("/ott/subscription")
	ottSubscription.POST("/create", ottSubscriptionHandler.CreateSubscription)
	ottSubscription.GET("/getall", ottSubscriptionHandler.CreateSubscriptionGET)

	postpaidBillFetch := e.Group("/billpayment/postpaid")
	postpaidBillFetch.POST("/save/bill", PostpaidBillFetchHandler.SavePostPaidBill)
	postpaidBillFetch.GET("/fetch/billpayments", PostpaidBillFetchHandler.FetchBill)

	postpaidMobileRecharge := e.Group("/billpayment/postpaid/recharge")
	postpaidMobileRecharge.POST("/recharge", PostpaidMobileRechargeHandler.RechargePostpaid)
	postpaidMobileRecharge.GET("/fetch/payments", PostpaidMobileRechargeHandler.FetchPostpaidRecharge)

	electricityBillFetch := e.Group("/billpayment/electricitybill")
	electricityBillFetch.POST("/save", ElectricityBillFetchHandler.SaveElectricityBill)
	electricityBillFetch.GET("/fetch", ElectricityBillFetchHandler.FetchBill)

	electricityBillPayment := e.Group("/billpayment/electricityrecharge")
	electricityBillPayment.POST("/save/payment", ElectricityBillPaymentHandler.MakePaymentPOST)
	electricityBillPayment.GET("/fetch/bills", ElectricityBillPaymentHandler.MakePaymentGET)

	Wallet := e.Group("/moneytransfer/wallet")
	Wallet.POST("/create", WalletCreateHandler.CreateWalletPOST)
	Wallet.GET("/check", WalletCreateHandler.CheckWalletExists)
	Wallet.POST("/verifyotp", WalletCreateHandler.VerifyOtpPOST)
	Wallet.GET("/verifyotp", WalletCreateHandler.VerifyOtpGET)

	Beneficiary := e.Group("/moneytransfer/beneficiary")
	Beneficiary.POST("/create", BeneficiaryHandler.AddBeneficiaryHandler)
	Beneficiary.GET("/get/details", BeneficiaryHandler.GetBeneficiariesHandler)
	Beneficiary.DELETE("/delete", BeneficiaryHandler.DeleteBeneficiaryHandler)
	Beneficiary.GET("",BeneficiaryHandler.VerifyDeleteBeneficiaryHandler)

	
	payout := e.Group("/payout")
	payout.POST("/initiate", PayoutHandler.InitiatePayout)

}
	