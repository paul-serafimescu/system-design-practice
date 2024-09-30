package api

import (
	"encoding/json"
	"fmt"
	"http-server/models"
	"http-server/service"
	"net/http"
)

// this function currently matches login credentials to a user
//
// next steps: create a session (whatever that is, thinking token) and direct client to websocket server
// so overall flow is login -> get wss -> establish conn to wss gateway -> update redis # of clients connected to wss instance
// so the wss should "register" on boot and regularly send a heartbeat to our API (could use ICMP? idk if we get RTT info)
// can use redis to check heartbeat status as well and set timeout since PG is pretty expensive
func CreateSession(w http.ResponseWriter, r *http.Request) {
	var sessionRequest models.SessionRequest
	if err := json.NewDecoder(r.Body).Decode(&sessionRequest); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, validated := service.ValidateSessionRequest(&sessionRequest)

	if !validated {
		http.Error(w, "Invalid Credentials", http.StatusNotFound)
	} else {
		fmt.Printf("%+v\n", user)
		w.Write([]byte("success!")) // change this
	}
}

func CreateNewAccount(w http.ResponseWriter, r *http.Request) {
	var signupRequest models.SignupRequest

	if err := json.NewDecoder(r.Body).Decode(&signupRequest); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	newUser, err := service.CreateNewAccount(&signupRequest)

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
	} else {
		fmt.Printf("created user: %v\n", newUser)
		w.Write([]byte("success!")) // change this
	}
}
