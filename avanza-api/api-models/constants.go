package api_models

import "net/http"

var DEF_HEADERS = http.Header{
	"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0"},
	"Content-Type": {"application/json"},
	"Connection":   {"keep-alive"},
}

const (
	BASE_URL              = "https://www.avanza.se"
	AUTH_URL              = "/_api/authentication/sessions/usercredentials"         //POST
	TOTP_URL              = "/_api/authentication/sessions/totp"                    //POST
	DISC_URL              = "/_api/authentication/sessions/webtoken"                //DELETE
	SESSION_INFO_URL      = "/_api/authentication/session/info/session"             //GET
	ACCOUNTS_URL          = "/_api/trading-critical/rest/alllightweightaccounts"    //GET
	POSITIONS_URL         = "/_api/position-data/positions"                         //GET
	PRICE_INFO_URL        = "/_mobile/market/index/"                                //GET
	HISTORICAL_PRICES_URL = "/_api/price-chart/stock/"                              //GET
	WARRANT_INFO_URL      = "/_api/market-guide/warrant/"                           //GET
	WARRANT_LIST_URL      = "/_api/market-warrant-filter/"                          //POST
	PLACE_ORDER_URL       = "/_api/trading-critical/rest/order/new"                 //POST
	MATCHING_PRICE_URL    = "/_api/trading/rest/matchingprice/"                     //POST
	CHECK_ORDER_URL       = "/_api/trading-critical/rest/order/find"                //GET
	MODIFY_ORDER_URL      = "/_api/trading-critical/rest/order/modify/"             //POST
	REQ_ID_URL            = "/_api/trading/rest/order/requestid"                    //GET
	VALIDATE_URL          = "/_api/trading-critical/rest/order/validation/validate" //POST
)
