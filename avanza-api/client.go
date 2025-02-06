package avanza_api

import (
	"fmt"
	"time"

	api_models "github.com/Melle101/mom-bot-v3/avanza-api/api-models"
	"github.com/pquerna/otp/totp"
)

func CreateClient(authInfo InitAuthInfo) (*ApiClient, error) {
	credInfo, err := getCredentials(authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get credentials: %w", err)
	}

	client := &ApiClient{
		CredInfo: credInfo,
		CredHeaders: map[string][]string{
			"Cookie":          {"csid=" + credInfo.Auth_session, "cstoken=" + credInfo.Cs_token, "AZACSRF=" + credInfo.Secret_token},
			"X-SecurityToken": {credInfo.Secret_token},
		},
	}

	return client, nil
}

func getCredentials(authInfo InitAuthInfo) (CredInfo, error) {

	url := api_models.BASE_URL + api_models.AUTH_URL
	payload := map[string]string{
		"username": authInfo.Username,
		"password": authInfo.Password,
	}

	resp, headers, err := HTTPPostHeaders[api_models.Auth_response](url, payload, nil)
	if err != nil {
		return CredInfo{}, fmt.Errorf("failed to make auth request: %w", err)
	}

	if (resp.Tfa_response == api_models.Tfa_response{}) {
		return CredInfo{
			Secret_token: headers.Get("X-SecurityToken"),
			Auth_session: resp.AuthenticationSession,
			Push_sub_id:  resp.PushSubscriptionId,
			Cs_token:     resp.CustomerSessionToken,
		}, nil
	} else {
		// Handle TFA
		return handleTFA(resp.Tfa_response.Transaction_id, authInfo.TOTP)
	}
}

func handleTFA(trID, totpToken string) (CredInfo, error) {
	url := api_models.BASE_URL + api_models.TOTP_URL

	totp_code, err := totp.GenerateCode(totpToken, time.Now())
	if err != nil {
		return CredInfo{}, fmt.Errorf("failed to generate TOTP code: %w", err)
	}

	payload := map[string]string{
		"method":   "TOTP",
		"totpCode": totp_code,
	}

	reqHeaders := map[string][]string{
		"Cookie": {"AZAMFATRANSACTION=" + trID},
	}

	resp, headers, err := HTTPPostHeaders[api_models.Auth_response](url, payload, reqHeaders)
	if err != nil {
		return CredInfo{}, fmt.Errorf("failed to make 2fa_auth request: %w", err)
	}

	if resp.AuthenticationSession != "" {
		return CredInfo{
			Secret_token: headers.Get("X-SecurityToken"),
			Auth_session: resp.AuthenticationSession,
			Push_sub_id:  resp.PushSubscriptionId,
			Cs_token:     resp.CustomerSessionToken,
		}, nil
	} else {
		return CredInfo{}, fmt.Errorf("unexpected TFA response: %v", resp.Tfa_response)
	}
}

func (c *ApiClient) Disconnect() (*api_models.SessionClose, error) {
	url := api_models.BASE_URL + api_models.DISC_URL

	s, err := HTTPDelete[api_models.SessionClose](url, c.CredHeaders)
	if err != nil {
		return s, fmt.Errorf("failed to close session: %w", err)
	}

	c.CredInfo = CredInfo{}
	c.CredHeaders = map[string][]string{}

	return s, nil
}
