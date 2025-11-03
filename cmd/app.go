package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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
	electricitybillfetch_repository "github.com/paybazar-backend/internals/repository/BillPayments/ElectricityBill"
	electricitybillpayment_repository "github.com/paybazar-backend/internals/repository/BillPayments/ElectricityRecharge"
	postpaidbill_repository "github.com/paybazar-backend/internals/repository/BillPayments/postpaidBillFetch"
	postpaidmobilerecharge_repository "github.com/paybazar-backend/internals/repository/BillPayments/postpaidmobilerecharge"
	dth_repository "github.com/paybazar-backend/internals/repository/Dth"
	moneytransfer_repository "github.com/paybazar-backend/internals/repository/MoneyTransfer"
	ott_repository "github.com/paybazar-backend/internals/repository/Ott"
	prepaidmobilerecharge_repository "github.com/paybazar-backend/internals/repository/PrepaidMobileRecharge"
	prepaidplanfetch_repository "github.com/paybazar-backend/internals/repository/PrepaidPlanFetch"
	payoutrepo "github.com/paybazar-backend/internals/repository/payout"
	electricitybillfetch_service "github.com/paybazar-backend/internals/service/BillPayments/ElectricityBill"
	electricitybillpayment_service "github.com/paybazar-backend/internals/service/BillPayments/ElectricityRecharge"
	postpaidbill_service "github.com/paybazar-backend/internals/service/BillPayments/postpaidBillFetch"
	postpaidmobilerecharge_service "github.com/paybazar-backend/internals/service/BillPayments/postpaidmobilerecharge"
	dth_service "github.com/paybazar-backend/internals/service/Dth"
	moneytranfer_service "github.com/paybazar-backend/internals/service/MoneyTranfer"
	ott_service "github.com/paybazar-backend/internals/service/Ott"
	prepaidmobilerecharge_service "github.com/paybazar-backend/internals/service/PrepaidMobileRecharge"
	prepaidplanfetch_service "github.com/paybazar-backend/internals/service/PrepaidPlanFetch"
	payoutservice "github.com/paybazar-backend/internals/service/payout"
	"github.com/paybazar-backend/pkg/database"
)

func InitializeApp() *echo.Echo {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}


	db := database.ConnectPostgres()
	database.RunMigrations(db)

	apiBaseURL := "https://v2bapi.rechargkit.biz"
	apiToken := os.Getenv("API_TOKEN")
	    fmt.Println("Loaded Token:", apiToken)

	MobileRecharge := prepaidmobilerecharge_repository.NewRechargeRepo(db)
	rechargeService := prepaidmobilerecharge_service.NewRechargeService(MobileRecharge, apiBaseURL, apiToken)
	rechargeHandler := prepaidmobilerecharge_handlers.NewRechargeHandler(rechargeService)

	prepaidPlanRepo := prepaidplanfetch_repository.NewPrepaidPlanRepository(db)
	prepaidPlanService := prepaidplanfetch_service.NewPrepaidPlanService(prepaidPlanRepo, apiBaseURL, apiToken)
	prepaidPlanHandler := prepaidplanfetch_handlers.NewPrepaidPlanHandler(prepaidPlanService)

	dthRechargeRepo := dth_repository.NewDTHRechargeRepo(db)
	dthService := dth_service.NewDTHRechargeService(dthRechargeRepo, apiBaseURL, apiToken)
	dthHandler := dth_handlers.NewDTHRechargeHandler(dthService)

	ottPlanRepo := ott_repository.NewOTTPlanRepo(db)
	ottPlanService := ott_service.NewOTTPlanService(ottPlanRepo, apiBaseURL, apiToken)
	ottPlanHandler := ott_handlers.NewOTTPlanHandler(ottPlanService)

	ottSubscriptionRepo := ott_repository.NewOTTSubscriptionRepo(db)
	ottSubscriptionService := ott_service.NewOTTSubscriptionService(ottSubscriptionRepo, apiBaseURL, apiToken)
	ottSubscriptionHandler := ott_handlers.NewOTTSubscriptionHandler(ottSubscriptionService)

	PostpaidBillFetchRepo := postpaidbill_repository.NewPostpaidBillFetchRepo(db)
	PostpaidBillFetchService := postpaidbill_service.NewPostpaidBillFetchService(PostpaidBillFetchRepo, apiBaseURL, apiToken)
	PostpaidBillFetchHandler := postpaidbill_handlers.NewPostpaidBillFetchHandler(PostpaidBillFetchService)

	PostPaidMobileRechargeRepo := postpaidmobilerecharge_repository.NewPostpaidMobileRechargeRepo(db)
	PostPaidMobileRechargeService := postpaidmobilerecharge_service.NewPostpaidRechargeService(PostPaidMobileRechargeRepo, apiBaseURL, apiToken)
	PostpaidMobileRechargeHandler := postpaidmobilerecharge_handlers.NewPostpaidRechargeHandler(PostPaidMobileRechargeService)

	ElectricityBillFetchRepo := electricitybillfetch_repository.NewElectricityBillFetchRepo(db)
	ElectricityBillFetchService := electricitybillfetch_service.NewElectricityBillFetchService(ElectricityBillFetchRepo, apiBaseURL, apiToken)
	ElectricityBillFetchHandler := electricitybillfetch_handlers.NewElectricityBillFetchHandler(ElectricityBillFetchService)

	ElectricityBillPaymentRepo := electricitybillpayment_repository.NewElectricityBillPaymentRepo(db)
	ElectricityBillPaymentService := electricitybillpayment_service.NewElectricityBillPaymentService(ElectricityBillPaymentRepo, apiBaseURL, apiToken)
	ElectricityBillPaymentHandler := electricitybillpayment_handlers.NewElectricityBillPaymentHandler(ElectricityBillPaymentService)

	WalletCreateRepo := moneytransfer_repository.NewWalletCreateRepo(db)
	WalletCreateService := moneytranfer_service.NewWalletService(WalletCreateRepo, nil, apiBaseURL, apiToken)
	WalletCreateHandler := moneytransfer_handlers.NewWalletHandler(WalletCreateService)

	BeneficiaryRepo := moneytransfer_repository.NewBeneficiaryRepository(db)
	BeneficiaryService := moneytranfer_service.NewBeneficiaryService(BeneficiaryRepo, apiBaseURL, apiToken)
	BeneficiaryHandler := moneytransfer_handlers.NewBeneficiaryHandler(BeneficiaryService)

	MoneyTransfer := moneytransfer_repository.NewMoneyTransferRepository(db)
	MoneyTransferService := moneytranfer_service.NewMoneyTransferService(MoneyTransfer, apiBaseURL, apiToken)
	MoneyTransferHandler := moneytransfer_handlers.NewMoneyTransferHandler(MoneyTransferService)	

	PayoutRepo := payoutrepo.NewPayoutRepo(db)
	PayoutService := payoutservice.NewPayoutService(PayoutRepo, apiBaseURL, apiToken)
	PayoutHandler := payouthandlers.NewPayoutHandler(PayoutService)

	e := echo.New()
	SetupRouter(e, rechargeHandler, prepaidPlanHandler, dthHandler, ottPlanHandler, ottSubscriptionHandler, PostpaidBillFetchHandler, PostpaidMobileRechargeHandler, ElectricityBillFetchHandler, ElectricityBillPaymentHandler, WalletCreateHandler, BeneficiaryHandler, MoneyTransferHandler, PayoutHandler)

	return e
}
