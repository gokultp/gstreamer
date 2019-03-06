package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/gokultp/gstreamer/internal/contracts"
	"github.com/gokultp/gstreamer/internal/serviceerrors"
	"github.com/gokultp/gstreamer/pkg/errors"
)

const (
	twitchScopes          = "channel:read:subscriptions user_read"
	twitchAuthURL         = "https://id.twitch.tv/oauth2/authorize"
	twitchTokenURL        = "https://id.twitch.tv/oauth2/token"
	twitchResponseType    = "code"
	twitchRedirectURL     = "/auth/cb"
	twitchGrantType       = "authorization_code"
	envTwitchClientID     = "TWITCH_CLIENT_ID"
	envTwitchClientSecret = "TWITCH_CLIENT_SECRET"
	envHost               = "HOST"
	authFieldRedirectURI  = "redirect_uri"
	authFieldClientID     = "client_id"
	authFieldClientSecret = "client_secret"
	authFieldResponseType = "response_type"
	authFieldScope        = "scope"
	authFieldGrantType    = "grant_type"
	authFieldCode         = "code"
)

func GetAuthRedirectURL() string {
	params := url.Values{}
	authURL, _ := url.Parse(twitchAuthURL)
	clientID := os.Getenv(envTwitchClientID)
	params.Set(authFieldScope, twitchScopes)
	params.Set(authFieldResponseType, twitchResponseType)
	params.Set(authFieldRedirectURI, os.Getenv(envHost)+twitchRedirectURL)
	params.Set(authFieldClientID, clientID)
	authURL.RawQuery = params.Encode()
	return authURL.String()
}

func GetAuthToken(code string) (*contracts.AuthToken, errors.IError) {
	params := url.Values{}
	clientID := os.Getenv(envTwitchClientID)
	clientSecret := os.Getenv(envTwitchClientSecret)
	params.Set(authFieldGrantType, twitchGrantType)
	params.Set(authFieldRedirectURI, twitchRedirectURL)
	params.Set(authFieldClientID, clientID)
	params.Set(authFieldClientSecret, clientSecret)
	params.Set(authFieldCode, code)

	resp, err := http.PostForm(twitchTokenURL, params)
	if err != nil {
		return nil, serviceerrors.TwitchRequestError(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, serviceerrors.TwitchRequestError(err.Error())
		}
		var authData contracts.AuthToken
		err = json.Unmarshal(bodyBytes, &authData)
		if err != nil {
			return nil, serviceerrors.TwitchRequestError(err.Error())
		}
		return &authData, nil
	}
	return nil, serviceerrors.TwitchRequestError("status: %d", resp.StatusCode)

}
