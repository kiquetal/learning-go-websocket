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
	log.Println("Authenticating user", u.FirstName)
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
	log.Println("Successful authentication")
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
func (repo *DBRepo) TestPusher(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["message"] = "Hello from Vigilate"
	log.Println("Sending message to pusher")
	error := repo.App.WsClient.Trigger("public-channel", "test-event", data)
	if error != nil {
		log.Println(error)
		return
	}

}
