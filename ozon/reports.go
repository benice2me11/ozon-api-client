package ozon

import (
	"context"
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Reports struct {
	client *core.Client
}

type GetReportsListParams struct {
	// Page number
	Page int32 `json:"page"`

	// The number of values on the page:
	//   - default value is 100,
	//   - maximum value is 1000
	PageSize int32 `json:"page_size"`

	// Default: "ALL"
	// Report type:
	//   - ALL — all reports,
	//   - SELLER_PRODUCTS — products report,,
	//   - SELLER_TRANSACTIONS — transactions report,
	//   - SELLER_PRODUCT_PRICES — product prices report,
	//   - SELLER_STOCK — stocks report,
	//   - SELLER_PRODUCT_MOVEMENT — products movement report,
	//   - SELLER_RETURNS — returns report,
	//   - SELLER_POSTINGS — shipments report,
	//   - SELLER_FINANCE — financial report
	ReportType string `json:"report_type" default:"ALL"`
}

type GetReportsListResponse struct {
	core.CommonResponse

	// Method result
	Result GetReportsListResult `json:"result"`
}

type GetReportsListResult struct {
	// Array with generated reports
	Reports []GetReportsListResultReport `json:"reports"`

	// Total number of reports
	Total int32 `json:"total"`
}

type GetReportsListResultReport struct {
	// Unique report identifier
	Code string `json:"code"`

	// Report creation date
	CreatedAt time.Time `json:"created_at"`

	// Error code when generating the report
	Error string `json:"error"`

	// Link to CSV file
	//
	// For a report with the SELLER_RETURNS type,
	// the link is available within 5 minutes after making a request.
	File string `json:"file"`

	// Array with the filters specified when the seller created the report
	Params map[string]string `json:"params"`

	// Report type:
	//   - SELLER_PRODUCTS — products report,
	//   - SELLER_TRANSACTIONS — transactions report,
	//   - SELLER_PRODUCT_PRICES — product prices report,
	//   - SELLER_STOCK — stocks report,
	//   - SELLER_PRODUCT_MOVEMENT — products movement report,
	//   - SELLER_RETURNS — returns report,
	//   - SELLER_POSTINGS — shipments report,
	//   - SELLER_FINANCE — financial report
	ReportType string `json:"report_type"`

	// Report generation status
	//   - `success`
	//   - `failed`
	Status string `json:"status"`
}

// Returns the list of reports that have been generated before
func (c Reports) GetList(ctx context.Context, params *GetReportsListParams) (*GetReportsListResponse, error) {
	url := "/v1/report/list"

	resp := &GetReportsListResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetReportDetailsParams struct {
	// Unique report identifier
	Code string `json:"code"`
}

type GetReportDetailsResponse struct {
	core.CommonResponse

	// Report details
	Result GetReportDetailResult `json:"result"`
}

type GetReportDetailResult struct {
	// Unique report identifier
	Code string `json:"code"`

	// Report creation date
	CreatedAt time.Time `json:"created_at"`

	// Error code when generating the report
	Error string `json:"error"`

	// Link to CSV file
	File string `json:"file"`

	// Array with the filters specified when the seller created the report
	Params map[string]string `json:"params"`

	// Report type
	ReportType ReportType `json:"report_type"`

	// Report generation status
	Status ReportInfoStatus `json:"status"`
}

// Returns information about a created report by its identifier
func (c Reports) GetReportDetails(ctx context.Context, params *GetReportDetailsParams) (*GetReportDetailsResponse, error) {
	url := "/v1/report/info"

	resp := &GetReportDetailsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetFinancialReportParams struct {
	// Report generation period
	Date GetFinancialReportDatePeriod `json:"date"`

	// Number of the page returned in the request
	Page int64 `json:"page"`

	// true, если нужно добавить дополнительные параметры в ответ
	WithDetails bool `json:"with_details"`

	// Number of items on the page
	PageSize int64 `json:"page_size"`
}

type GetFinancialReportDatePeriod struct {
	// Date from which the report is calculated
	From time.Time `json:"from"`

	// Date up to which the report is calculated
	To time.Time `json:"to"`
}

type GetFinancialReportResponse struct {
	core.CommonResponse

	// Method result
	Result GetFinancialResultResult `json:"result"`
}

type GetFinancialResultResult struct {
	// Reports list
	CashFlows []GetFinancialResultResultCashflow `json:"cash_flows"`

	// Detailed info
	Details GetFinancialResultResultDetail `json:"details"`

	// Number of pages with reports
	PageCount int64 `json:"page_count"`
}

type GetFinancialResultResultCashflow struct {
	// Period data
	Period GetFinancialResultResultCashflowPeriod `json:"period"`

	// Sum of sold products prices
	OrdersAmount float64 `json:"orders_amount"`

	// Sum of returned products prices
	ReturnsAmount float64 `json:"returns_amount"`

	// Ozon sales commission
	CommissionAmount float64 `json:"commission_amount"`

	// Additional services cost
	ServicesAmount float64 `json:"services_amount"`

	// Logistic services cost
	ItemDeliveryAndReturnAmount float64 `json:"item_delivery_and_return_amount"`

	// Code of the currency used to calculate the commissions
	CurrencyCode string `json:"currency_code"`
}

type GetFinancialResultResultCashflowPeriod struct {
	// Period identifier
	Id int64 `json:"id"`

	// Period start
	Begin time.Time `json:"begin"`

	// Period end
	End time.Time `json:"end"`
}

type GetFinancialResultResultDetail struct {
	// Balance on the beginning of period
	BeginBalanceAmount float64 `json:"begin_balance_amount"`

	// Orders
	Delivery GetFinancialResultResultDetailDelivery `json:"delivery"`

	// Amount to be paid for the period
	InvoiceTransfer float64 `json:"invoice_transfer"`

	// Transfer under loan agreements
	Loan float64 `json:"loan"`

	// Paid for the period
	Payments []GetFinancialResultResultDetailPayment `json:"payments"`

	// Period data
	Period GetFinancialResultResultDetailPeriod `json:"period"`

	// Returns and cancellations
	Return GetFinancialResultResultDetailReturn `json:"return"`

	// rFBS transfers
	RFBS GetFinancialResultResultDetailRFBS `json:"rfbs"`

	// Services
	Services GetFinancialResultResultDetailService `json:"services"`

	// Compensation and other accruals
	Others GetFinancialResultResultDetailOthers `json:"others"`

	// Balance at the end of the period
	EndBalanceAmount float64 `json:"end_balance_amount"`
}

type GetFinancialResultResultDetailDelivery struct {
	// Total amount
	Total float64 `json:"total"`

	// Amount for which products were purchased, including commission fees
	Amount float64 `json:"amount"`

	// Processing and delivery fees
	DeliveryServices GetFinancialResultResultDetailDeliveryServices `json:"delivery_services"`
}

type GetFinancialResultResultDetailDeliveryServices struct {
	// Total amount
	Total float64 `json:"total"`

	// Details
	Items []GetFinancialResultResultDetailDeliveryServicesItem `json:"items"`
}

type GetFinancialResultResultDetailDeliveryServicesItem struct {
	// Operation name
	Name DetailsDeliveryItemName `json:"name"`

	// Amount by operation
	Price float64 `json:"price"`
}

type GetFinancialResultResultDetailPayment struct {
	// Currency
	CurrencyCode string `json:"currency_code"`

	// Payment amount
	Payment float64 `json:"payment"`
}

type GetFinancialResultResultDetailPeriod struct {
	// Period start
	Begin time.Time `json:"begin"`

	// Period end
	End time.Time `json:"end"`

	// Period identifier
	Id int64 `json:"id"`
}

type GetFinancialResultResultDetailReturn struct {
	// Total amount
	Total float64 `json:"total"`

	// Amount of returns received, including commission fees
	Amount float64 `json:"amount"`

	// Returns and cancellation fees
	ReturnServices GetFinancialResultResultDetailReturnServices `json:"return_services"`
}

type GetFinancialResultResultDetailReturnServices struct {
	// Total amount
	Total float64 `json:"total"`

	// Details
	Items []GetFinancialResultResultDetailReturnServicesItem `json:"items"`
}

type GetFinancialResultResultDetailReturnServicesItem struct {
	// Operation name
	Name DetailsReturnServiceName `json:"name"`

	// Amount by operation
	Price float64 `json:"price"`
}

type GetFinancialResultResultDetailRFBS struct {
	// Total amount
	Total float64 `json:"total"`

	// Transfers from customers
	TransferDelivery float64 `json:"transfer_delivery"`

	// Return of transfers to customers
	TransferDeliveryReturn float64 `json:"transfer_delivery_return"`

	// Compensation of delivery fees
	CompensationDeliveryReturn float64 `json:"compensation_delivery_return"`

	// Transfers of partial refunds to customers
	PartialCompensation float64 `json:"partial_compensation"`

	// Compensation of partial refunds
	PartialCompensationReturn float64 `json:"partial_compensation_return"`
}

type GetFinancialResultResultDetailService struct {
	// Total amount
	Total float64 `json:"total"`

	// Details
	Items []GetFinancialResultResultDetailServiceItem `json:"items"`
}

type GetFinancialResultResultDetailServiceItem struct {
	// Operation name
	Name DetailsServiceItemName `json:"name"`

	// Amount by operation
	Price float64 `json:"price"`
}

type GetFinancialResultResultDetailOthers struct {
	// Total amount
	Total float64 `json:"total"`

	// Details
	Items []GetFinancialResultResultDetailOthersItem `json:"items"`
}

type GetFinancialResultResultDetailOthersItem struct {
	// Operation name
	Name DetailsOtherItemName `json:"name"`

	// Amount by operation
	Price float64 `json:"price"`
}

// Returns information about a created report by its identifier
func (c Reports) GetFinancial(ctx context.Context, params *GetFinancialReportParams) (*GetFinancialReportResponse, error) {
	url := "/v1/finance/cash-flow-statement/list"

	resp := &GetFinancialReportResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProductsReportParams struct {
	// Default: "DEFAULT"
	// Response language:
	//   - RU — Russian
	//   - EN — English
	Language string `json:"language" default:"DEFAULT"`

	// Product identifier in the seller's system
	OfferId []string `json:"offer_id"`

	// Search by record content, checks for availability
	Search string `json:"search"`

	// Product identifier in the Ozon system, SKU
	SKU []int64 `json:"sku"`

	// Default: "ALL"
	// Filter by product visibility
	Visibility string `json:"visibility" default:"ALL"`
}

type GetProductsReportResponse struct {
	core.CommonResponse

	// Method result
	Result GetProductsReportResult `json:"result"`
}

type GetProductsReportResult struct {
	// Unique report identifier
	Code string `json:"code"`
}

// Method for getting a report with products data. For example, Ozon ID, number of products, prices, status
func (c Reports) GetProducts(ctx context.Context, params *GetProductsReportParams) (*GetProductsReportResponse, error) {
	url := "/v1/report/products/create"

	resp := &GetProductsReportResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetReturnsReportParams struct {
	// Filter
	Filter *GetReturnsReportsFilter `json:"filter,omitempty"`

	// Default: "DEFAULT"
	// Response language:
	//   - RU — Russian
	//   - EN — English
	Language string `json:"language" default:"DEFAULT"`
}

type GetReturnsReportsFilter struct {
	// Order delivery scheme: fbs — delivery from seller's warehouse
	DeliverySchema string `json:"delivery_schema"`

	// Date from which the data is displayed in the report.
	//
	// Available for the last three months only
	DateFrom time.Time `json:"date_from"`

	// Date up to which the data is displayed in the report.
	//
	// Available for the last three months only
	DateTo time.Time `json:"date_to"`

	// Order status
	Status string `json:"status"`
}

type GetReturnsReportResponse struct {
	core.CommonResponse

	Result GetReturnsReportResult `json:"result"`
}

type GetReturnsReportResult struct {
	// Unique report identifier. The report is available for downloading within 3 days after making a request.
	Code string `json:"code"`
}

// Method for getting a report on FBO and FBS returns
func (c Reports) GetReturns(ctx context.Context, params *GetReturnsReportParams) (*GetReturnsReportResponse, error) {
	url := "/v2/report/returns/create"

	resp := &GetReturnsReportResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetShipmentReportParams struct {
	// Filter
	Filter *GetShipmentReportFilter `json:"filter,omitempty"`

	// Default: "DEFAULT"
	// Response language:
	//   - RU — Russian
	//   - EN — English
	Language string `json:"language" default:"DEFAULT"`
}

type GetShipmentReportFilter struct {
	// Cancellation reason identifier
	CancelReasonId []int64 `json:"cancel_reason_id"`

	// The scheme of operation is FBO or FBS.
	//
	// Only one of the parameters can be passed to the array per query:
	//
	// fbo - to get a report by FBO scheme,
	// fbs - to get a report by FBS scheme
	DeliverySchema []string `json:"delivery_schema"`

	// Product identifier
	OfferId string `json:"offer_id"`

	// Order processing start date and time
	ProcessedAtFrom *core.TimeFormat `json:"processed_at_from,omitempty"`

	// Time when the order appeared in your personal account
	ProcessedAtTo *core.TimeFormat `json:"processed_at_to,omitempty"`

	// Product identifier in the Ozon system, SKU
	SKU []int64 `json:"sku"`

	// Status text
	StatusAlias []string `json:"status_alias"`

	// Numerical status
	Statuses []int64 `json:"statused"`

	// Product name
	Title string `json:"title"`
}

type GetShipmentReportResponse struct {
	core.CommonResponse

	// Method result
	Result GetShipmentReportResult `json:"result"`
}

type GetShipmentReportResult struct {
	// Unique report identifier
	Code string `json:"code"`
}

// Shipment report with orders details:
//   - order statuses
//   - processing start date
//   - order numbers
//   - shipment numbers
//   - shipment costs
//   - shipments contents
func (c Reports) GetShipment(ctx context.Context, params *GetShipmentReportParams) (*GetShipmentReportResponse, error) {
	url := "/v1/report/postings/create"

	resp := &GetShipmentReportResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type IssueOnDiscountedProductsResponse struct {
	core.CommonResponse

	// Unique report identifier
	Code string `json:"code"`
}

// Generates a report on discounted products in Ozon warehouses.
// For example, Ozon can discount a product due to damage when delivering.
//
// Returns report identifier. To get the report, send the identifier in the request body of a method `/v1/report/discounted/info`
func (c Reports) IssueOnDiscountedProducts(ctx context.Context) (*IssueOnDiscountedProductsResponse, error) {
	url := "/v1/report/discounted/create"

	resp := &IssueOnDiscountedProductsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetFBSStocksParams struct {
	// Response language
	Language string `json:"language"`

	// Warehouses identifiers
	WarehouseIds []int64 `json:"warehouse_id"`
}

type GetFBSStocksResponse struct {
	core.CommonResponse

	// Method result
	Result GetFBSStocksResult `json:"result"`
}

type GetFBSStocksResult struct {
	// Unique report identifier
	Code string `json:"code"`
}

// Report with information about the number of available and reserved products in stock.
//
// The method returns a report identifier.
// To get the report, send the identifier in the request of the `/v1/report/info` method.
func (c Reports) GetFBSStocks(ctx context.Context, params *GetFBSStocksParams) (*GetFBSStocksResponse, error) {
	url := "/v1/report/warehouse/stock"

	resp := &GetFBSStocksResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
