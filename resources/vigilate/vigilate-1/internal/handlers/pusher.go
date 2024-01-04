package handlers

import (
	"github.com/pusher/pusher-http-go"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (repo *DBRepo) PusherAuth(w http.ResponseWriter, r *http.Request) {
	userId := repo.App.Session.GetInt(r.Context(), "userID")
	u, _ := repo.DB.GetUserById(userId)

	params, _ := ioutil.ReadAll(r.Body)
	presenceData := pusher.MemberData{
		UserID: strconv.Itoa(userId),
		UserInfo: map[string]string{
			"name": u.FirstName,
			"id":   strconv.Itoa(userId),
		},
	}
	response, err := app.WsClient.AuthenticatePresenceChannel(params, presenceData)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
