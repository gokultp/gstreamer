package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gokultp/gstreamer/internal/dbmodels"
	"github.com/gokultp/gstreamer/internal/helpers"
)

func StreamHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetStream(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetStream(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		HandleError(err, w)
		return
	}
	user, err := dbmodels.GetUserByID(*userID)
	if err != nil {
		HandleError(err, w)
		return
	}
	stream, err := helpers.GetStreamByUser(*user.FavStreamer, *user.AccessToken)
	if err != nil {
		HandleError(err, w)
		return
	}
	var data map[string]interface{}
	fmt.Println(string(stream))
	json.Unmarshal(stream, &data)
	// if data["stream"] != nil {
	// 	stream := data["stream"].(map[string]interface{})
	// 	if stream["channel"] != nil {
	// 		channel := stream["channel"].(map[string]interface{})
	// 		chId := uint64(channel["_id"].(float64))

	// 	}

	// }
	chat, err := helpers.GetChatRooms(*user.ID, *user.AccessToken)
	if err != nil {
		HandleError(err, w)
		return
	}
	var chatData map[string]interface{}
	json.Unmarshal(chat, &chatData)

	data["chat"] = chatData

	jsonResponse(data, w)
}
