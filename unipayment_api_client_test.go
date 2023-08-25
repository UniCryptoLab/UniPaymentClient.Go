package unipayment

import (
	. "UniPaymentClient.Go/models"
	"testing"
)

var (
	authParams = AuthParams{ClientID: "74feb539-ba5a-4ae9-b901-4da4fb539574",
		ClientSecret: "BsoRhgqzhR1TYMtwTRYdPxBTvR5rxkW9K",
		AppID:        "2a9bd90b-fe95-4659-83cb-04de662fbbac"}
	apiClient = NewAPIClient(NewConfiguration(authParams)).UnipaymentApiClient
)

func TestApiClient_GetCurrencies(t *testing.T) {
	response, _, err := apiClient.GetCurrencies()
	if err != nil {
		t.Error(err)
		return
	}
	want := "OK"
	if response.Code != want {
		t.Errorf("got %q, wanted %q", response.Code, want)
	}
}

func TestApiClient_GetExchangeRateByFiatCurrency(t *testing.T) {
	response, _, err := apiClient.GetExchangeRateByFiatCurrency("USD")
	if err != nil {
		t.Error(err)
		return
	}
	want := "OK"
	if response.Code != want {
		t.Errorf("got %q, wanted %q", response.Code, want)
	}
}

func TestApiClient_QueryIps(t *testing.T) {
	response, _, err := apiClient.QueryIps()
	if err != nil {
		t.Error(err)
		return
	}
	want := "OK"
	if response.Code != want {
		t.Errorf("got %q, wanted %q", response.Code, want)
	}
}

func TestApiClient_CheckIPN(t *testing.T) {
	notify := "{\"ipn_type\":\"invoice\",\"event\":\"invoice_created\",\"app_id\":\"cee1b9e2-d90c-4b63-9824-d621edb38012\",\"invoice_id\":\"12wQquUmeCPUx3qmp3aHnd\",\"order_id\":\"ORDER_123456\",\"price_amount\":2.0,\"price_currency\":\"USD\",\"network\":null,\"address\":null,\"pay_currency\":null,\"pay_amount\":0.0,\"exchange_rate\":0.0,\"paid_amount\":0.0,\"confirmed_amount\":0.0,\"refunded_price_amount\":0.0,\"create_time\":\"2022-09-14T04:57:54.5599307Z\",\"expiration_time\":\"2022-09-14T05:02:54.559933Z\",\"status\":\"New\",\"error_status\":\"None\",\"ext_args\":\"Merchant Pass Through Data\",\"transactions\":null,\"notify_id\":\"fd58cedd-67c6-4053-ae65-2f6fb09a7d2c\",\"notify_time\":\"0001-01-01T00:00:00\"}"
	response, _, err := apiClient.CheckIPN(notify)
	if err != nil {
		t.Error(err)
		return
	}
	want := "OK"
	if response.Code != want {
		t.Errorf("got %q, wanted %q", response.Code, want)
	}
}

func TestApiClient_QueryInvoices(t *testing.T) {
	queryInvoiceRequest := QueryInvoicesRequest{OrderId: "ORDER_123456", PageNo: 1, PageSize: 10}
	response, _, err := apiClient.QueryInvoices(queryInvoiceRequest)
	if err != nil {
		t.Error(err)
		return
	}
	want := "OK"
	if response.Code != want {
		t.Errorf("got %q, wanted %q", response.Code, want)
	}
}

func TestApiClient_CreateInvoice(t *testing.T) {
	createInvoiceRequest := CreateInvoiceRequest{
		Title:         "MacBook Pro",
		Description:   "MacBook Pro(256G)",
		Lang:          "en",
		AppId:         authParams.AppID,
		PriceAmount:   2.0,
		PriceCurrency: "USD",
		PayCurrency:   "USDT",
		NotifyUrl:     "https://demo-payment.requestcatcher.com/test",
		RedirectUrl:   "https://www.example.com",
		OrderId:       "ORDER_123456",
		ExtArgs:       "Merchant Pass Through Data",
		ConfirmSpeed:  "Medium",
	}
	response, _, err := apiClient.CreateInvoice(createInvoiceRequest)
	if err != nil {
		t.Error(err)
		return
	}
	want := "OK"
	if response.Code != want {
		t.Errorf("got %q, wanted %q", response.Code, want)
	}
}

func TestApiClient_QueryInvoiceById(t *testing.T) {
	response, _, err := apiClient.QueryInvoiceById("SrAARgNrPgvveiBQtNc4gk")
	if err != nil {
		t.Error(err)
		return
	}
	want := "OK"
	if response.Code != want {
		t.Errorf("got %q, wanted %q", response.Code, want)
	}
}

func TestApiClient_GetWalletBalances(t *testing.T) {
	response, _, err := apiClient.GetWalletBalances()
	if err != nil {
		t.Error(err)
		return
	}
	want := "OK"
	if response.Code != want {
		t.Errorf("got %q, wanted %q", response.Code, want)
	}
}

func TestApiClient_CreateWithdrawal(t *testing.T) {
	createWithdrawalRequest := CreateWithdrawalRequest{
		Network:     "NETWORK_BSC",
		AssetType:   "USDT",
		Amount:      1.01,
		AutoConfirm: false,
		IncludeFee:  false,
	}
	response, _, err := apiClient.CreateWithdrawal(createWithdrawalRequest)
	if err != nil {
		t.Error(err)
		return
	}
	want := "OK"
	if response.Code != want {
		t.Errorf("got %q, wanted %q", response.Code, want)
	}
}

func TestApiClient_GetWithdrawalById(t *testing.T) {
	response, _, err := apiClient.GetWithdrawalById("a6389658-ac47-42f7-b71e-4bd1dc51ee2d")
	if err != nil {
		t.Error(err)
		return
	}
	want := "OK"
	if response.Code != want {
		t.Errorf("got %q, wanted %q", response.Code, want)
	}
}

func TestApiClient_CancelWithdrawal(t *testing.T) {
	cancelWithdrawalRequest := CancelWithdrawalRequest{Id: "a6389658-ac47-42f7-b71e-4bd1dc51ee2d"}
	response, _, err := apiClient.CancelWithdrawal(cancelWithdrawalRequest)
	if err != nil {
		t.Error(err)
		return
	}
	want := "OK"
	if response.Code != want {
		t.Errorf("got %q, wanted %q", response.Code, want)
	}
}

func TestApiClient_CreatePayout(t *testing.T) {
	createPayoutRequest := CreatePayoutRequest{
		Items:     []PayoutRequestItem{{Address: "USDC", Amount: 1.01}},
		Network:   "NETWORK_BSC",
		AssetType: "USDT",
	}
	response, _, err := apiClient.CreatePayout(createPayoutRequest)
	if err != nil {
		t.Error(err)
		return
	}
	want := "OK"
	if response.Code != want {
		t.Errorf("got %q, wanted %q", response.Code, want)
	}
}

func TestApiClient_GetPayoutById(t *testing.T) {
	response, _, err := apiClient.GetPayoutById("a6389658-ac47-42f7-b71e-4bd1dc51ee2d")
	if err != nil {
		t.Error(err)
		return
	}
	want := "OK"
	if response.Code != want {
		t.Errorf("got %q, wanted %q", response.Code, want)
	}
}

func TestApiClient_QueryPayouts(t *testing.T) {
	response, _, err := apiClient.QueryPayouts()
	if err != nil {
		t.Error(err)
		return
	}
	want := "OK"
	if response.Code != want {
		t.Errorf("got %q, wanted %q", response.Code, want)
	}
}
