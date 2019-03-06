package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gokultp/gstreamer/internal/serviceerrors"
	"github.com/gokultp/gstreamer/pkg/errors"
)

func SubscribeEvent(accessToken, callback, topic string) errors.IError {
	data := map[string]interface{}{
		"hub.callback":      callback,
		"hub.topic":         topic,
		"hub.lease_seconds": 6000,
		"hub.mode":          "subscribe",
	}

	jsonPayload, err := json.Marshal(data)
	if err != nil {
		return serviceerrors.TwitchRequestError(err.Error())
	}
	req, _ := http.NewRequest(http.MethodPost, "https://api.twitch.tv/helix/webhooks/hub", bytes.NewBuffer(jsonPayload))
	clientID := os.Getenv(envTwitchClientID)
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", accessToken))
	req.Header.Add("Client-ID", clientID)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return serviceerrors.TwitchRequestError(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println("asdf", resp.StatusCode, string(bodyBytes))
	if resp.StatusCode > 299 {
		if err != nil {
			return serviceerrors.TwitchRequestError(err.Error())
		}

		return serviceerrors.TwitchRequestError("status: %d, message:%s", resp.StatusCode, string(bodyBytes))
	}
	return nil
}
