package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gokultp/gstreamer/internal/contracts"
	"github.com/gokultp/gstreamer/internal/serviceerrors"
	"github.com/gokultp/gstreamer/pkg/errors"
)

const (
	twitchBaseURL = "https://api.twitch.tv/kraken"
)

func GetUserInfo(accessToken string) (*contracts.User, errors.IError) {
	clientID := os.Getenv(envTwitchClientID)
	req, _ := http.NewRequest(http.MethodGet, twitchBaseURL+"/user", nil)
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", accessToken))
	req.Header.Add("Client-ID", clientID)
	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, serviceerrors.TwitchRequestError(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, serviceerrors.TwitchRequestError(err.Error())
		}
		var user contracts.Streamer
		err = json.Unmarshal(bodyBytes, &user)
		if err != nil {
			return nil, serviceerrors.TwitchRequestError(err.Error())
		}
		return contracts.ConvertStreamer(user), nil
	}
	return nil, serviceerrors.TwitchRequestError("endpoint: /user, status: %d", resp.StatusCode)

}

func GetTwitchUserByName(name, accessToken string) (*contracts.User, errors.IError) {
	clientID := os.Getenv(envTwitchClientID)
	req, _ := http.NewRequest(http.MethodGet, twitchBaseURL+"/users?login="+name, nil)
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", accessToken))
	req.Header.Add("Client-ID", clientID)
	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, serviceerrors.TwitchRequestError(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
	if err != nil {
		return nil, serviceerrors.TwitchRequestError(err.Error())
	}
	if resp.StatusCode == http.StatusOK {

		var users struct {
			Users []contracts.Streamer `json:"users"`
		}
		err = json.Unmarshal(bodyBytes, &users)
		if err != nil {
			return nil, serviceerrors.TwitchRequestError(err.Error())
		}
		if len(users.Users) == 0 {
			return nil, serviceerrors.ResourceNotFoundError("%s", "No such users found")
		}
		return contracts.ConvertStreamer(users.Users[0]), nil
	}
	return nil, serviceerrors.TwitchRequestError("endpoint: /users, status: %d", resp.StatusCode)

}

func GetStreamByUser(userId uint64, accessToken string) ([]byte, errors.IError) {
	clientID := os.Getenv(envTwitchClientID)
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/streams/%d", twitchBaseURL, userId), nil)
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", accessToken))
	req.Header.Add("Client-ID", clientID)
	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, serviceerrors.TwitchRequestError(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, serviceerrors.TwitchRequestError(err.Error())
	}
	return bodyBytes, nil
}

func GetChatRooms(channel uint64, accessToken string) ([]byte, errors.IError) {
	clientID := os.Getenv(envTwitchClientID)
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/chat/80121757/rooms", twitchBaseURL), nil)
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", accessToken))
	req.Header.Add("Client-ID", clientID)
	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, serviceerrors.TwitchRequestError(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, serviceerrors.TwitchRequestError(err.Error())
	}
	return bodyBytes, nil
}
