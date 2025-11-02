package database

import (
	"database/sql"
	"log"
)

func RunMigrations(db *sql.DB) {
	query := `
	-- ===========================================
	-- Mobile Recharge Tables
	-- ===========================================
	CREATE TABLE IF NOT EXISTS mobile_recharge_transactions (
		id SERIAL PRIMARY KEY,
		mobile_no VARCHAR(15),
		operator_code INT,
		amount NUMERIC,
		partner_request_id VARCHAR(100),
		circle INT,
		recharge_type INT,
		dest VARCHAR(50),
		user_var1 VARCHAR(100),
		user_var2 VARCHAR(100),
		user_var3 VARCHAR(100),
		api_error INT,
		api_msg TEXT,
		api_status INT,
		order_id VARCHAR(100),
		optransid VARCHAR(100),
		commission NUMERIC,
		api_response JSONB,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS prepaid_plans (
		id SERIAL PRIMARY KEY,
		operator_code INT NOT NULL,
		circle INT NOT NULL,
		response_data JSONB,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	-- ===========================================
	-- DTH Tables
	-- ===========================================
	CREATE TABLE IF NOT EXISTS dth_recharge_transactions (
		id SERIAL PRIMARY KEY,
		customer_id VARCHAR(50),
		operator_code INT,
		amount NUMERIC,
		partner_request_id VARCHAR(100),
		dest VARCHAR(50),
		user_var1 VARCHAR(100),
		user_var2 VARCHAR(100),
		user_var3 VARCHAR(100),
		api_error INT,
		api_msg TEXT,
		api_status INT,
		order_id VARCHAR(100),
		optransid VARCHAR(100),
		commission NUMERIC,
		api_response JSONB,
		created_at TIMESTAMP DEFAULT NOW()
	);

	-- ===========================================
	-- OTT Tables
	-- ===========================================
	CREATE TABLE IF NOT EXISTS ott_plans (
		id SERIAL PRIMARY KEY,
		operator_code INT,
		response_data JSONB,
		created_at TIMESTAMP DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS ott_subscriptions (
		id SERIAL PRIMARY KEY,
		mobile_no VARCHAR(15),
		operator_code INT,
		amount NUMERIC,
		plan_id VARCHAR(50),
		customer_email VARCHAR(100),
		partner_request_id VARCHAR(100),
		user_var1 VARCHAR(100),
		user_var2 VARCHAR(100),
		user_var3 VARCHAR(100),
		api_error INT,
		api_msg TEXT,
		api_status INT,
		order_id VARCHAR(100),
		optransid VARCHAR(100),
		commission NUMERIC,
		api_response JSONB,
		created_at TIMESTAMP DEFAULT NOW()
	);

	-- ===========================================
	-- Postpaid Tables
	-- ===========================================
	CREATE TABLE IF NOT EXISTS postpaid_bill_fetch_logs (
		id SERIAL PRIMARY KEY,
		operator_code INT,
		mobile_no VARCHAR(20),
		api_error INT,
		api_status INT,
		api_msg TEXT,
		bill_amount TEXT,
		bill_due_date TEXT,
		bill_date TEXT,
		customer_name TEXT,
		api_response JSONB,
		created_at TIMESTAMP DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS postpaid_mobile_recharge_transactions (
		id SERIAL PRIMARY KEY,
		mobile_no VARCHAR(15),
		operator_code INT,
		amount NUMERIC,
		partner_request_id VARCHAR(100),
		circle INT,
		recharge_type INT,
		user_var1 VARCHAR(100),
		user_var2 VARCHAR(100),
		user_var3 VARCHAR(100),
		api_error INT,
		api_msg TEXT,
		api_status INT,
		order_id VARCHAR(100),
		optransid VARCHAR(100),
		commission NUMERIC,
		api_response JSONB,
		created_at TIMESTAMP DEFAULT NOW()
	);

	-- ===========================================
	-- Electricity Tables
	-- ===========================================
	CREATE TABLE IF NOT EXISTS electricity_bill_fetch_logs (
		id SERIAL PRIMARY KEY,
		operator_code INT,
		consumer_id VARCHAR(50),
		api_error INT,
		api_status INT,
		api_msg TEXT,
		consumer_name TEXT,
		bill_amount TEXT,
		bill_due_date TEXT,
		bill_date TEXT,
		api_response JSONB,
		created_at TIMESTAMP DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS electricity_bill_payment_transactions (
		id SERIAL PRIMARY KEY,
		p1 VARCHAR(50),
		p2 VARCHAR(50),
		p3 VARCHAR(50),
		customer_email VARCHAR(100),
		operator_code INT,
		amount NUMERIC,
		partner_request_id VARCHAR(100),
		api_error INT,
		api_msg TEXT,
		api_status INT,
		order_id VARCHAR(100),
		optransid VARCHAR(100),
		partnerreqid VARCHAR(100),
		commission NUMERIC,
		user_var1 VARCHAR(100),
		user_var2 VARCHAR(100),
		user_var3 VARCHAR(100),
		api_response JSONB,
		created_at TIMESTAMP DEFAULT NOW()
	);

	-- ===========================================
	-- Wallet Tables
	-- ===========================================
	CREATE TABLE IF NOT EXISTS wallet_create_requests (
		id SERIAL PRIMARY KEY,
		mobile_no VARCHAR(15),
		api_error INT,
		api_msg TEXT,
		request_no VARCHAR(100),
		description TEXT,
		api_response JSONB,
		created_at TIMESTAMP DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS wallet_verify_logs (
		id SERIAL PRIMARY KEY,
		mobile_no VARCHAR(15),
		request_no VARCHAR(100),
		otp VARCHAR(10),
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		address_line1 TEXT,
		address_line2 TEXT,
		city VARCHAR(100),
		state VARCHAR(50),
		pin_code VARCHAR(10),
		api_error INT,
		api_msg TEXT,
		description TEXT,
		api_response JSONB,
		created_at TIMESTAMP DEFAULT NOW()
	);

	-- ===========================================
	-- Beneficiary Tables
	-- ===========================================
	CREATE TABLE IF NOT EXISTS beneficiary_add_requests (
		id SERIAL PRIMARY KEY,
		mobile_no VARCHAR(15),
		beneficiary_name VARCHAR(100),
		bank_name VARCHAR(100),
		account_no VARCHAR(50),
		ifsc VARCHAR(20),
		api_error INT,
		api_msg TEXT,
		beneficiary_id VARCHAR(100),
		description TEXT,
		api_response JSONB,
		created_at TIMESTAMP DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS beneficiary_list_logs (
		id SERIAL PRIMARY KEY,
		mobile_no VARCHAR(15),
		api_error INT,
		api_msg TEXT,
		description TEXT,
		beneficiary_id VARCHAR(100),
		uuid VARCHAR(100),
		account_no VARCHAR(50),
		ifsc_code VARCHAR(20),
		bank_name VARCHAR(100),
		account_holder_name VARCHAR(100),
		api_response JSONB,
		created_at TIMESTAMP DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS beneficiary_delete_requests (
		id SERIAL PRIMARY KEY,
		mobile_no VARCHAR(15),
		api_error INT,
		api_msg TEXT,
		request_no VARCHAR(100),
		description TEXT,
		api_response JSONB,
		created_at TIMESTAMP DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS beneficiary_delete_verifications (
		id SERIAL PRIMARY KEY,
		mobile_no VARCHAR(15),
		request_no VARCHAR(100),
		otp VARCHAR(10),
		beneficiary_id VARCHAR(100),
		api_error INT,
		api_msg TEXT,
		description TEXT,
		api_response JSONB,
		created_at TIMESTAMP DEFAULT NOW()
	);
		-- Money Transfer transactions
	CREATE TABLE IF NOT EXISTS money_transfer_transactions (
		id SERIAL PRIMARY KEY,
		mobile_no VARCHAR(15) NOT NULL,
		beneficiary_name VARCHAR(100) NOT NULL,
		beneficiary_code VARCHAR(100),
		partner_request_id VARCHAR(100),
		amount NUMERIC,
		account_no VARCHAR(50),
		bank_name VARCHAR(150),
		ifsc VARCHAR(20),
		transfer_type VARCHAR(10),
		api_error INT,
		api_msg TEXT,
		api_status INT,
		order_id VARCHAR(100),
		optransid VARCHAR(100),
		partnerreqid VARCHAR(100),
		user_var1 VARCHAR(100),
		user_var2 VARCHAR(100),
		user_var3 VARCHAR(100),
		api_response JSONB,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	);

	-- Money Transfer status logs
	CREATE TABLE IF NOT EXISTS money_transfer_status_logs (
		id SERIAL PRIMARY KEY,
		transaction_id VARCHAR(100) NOT NULL,
		api_error INT,
		api_msg TEXT,
		api_status TEXT,
		transaction_status TEXT,
		api_response JSONB,
		created_at TIMESTAMP DEFAULT NOW()
	);

	`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("❌ Database migration failed: %v", err)
	}

	log.Println("✅ Database migrations executed successfully!")
}
