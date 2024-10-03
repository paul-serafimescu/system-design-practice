package api

import (
	"encoding/json"
	"http-server/models"
	"http-server/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
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

	_, validated := service.ValidateSessionRequest(&sessionRequest)

	chatServer, err := service.GetChatWebsocketServer()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if !validated {
		http.Error(w, "Invalid Credentials", http.StatusNotFound)
	} else {
		resp, _ := json.Marshal(models.SessionResponse{
			WsHostname: chatServer.Hostname,
			WsPort:     chatServer.Port,
			Token:      uuid.NewString(), // TODO: jwt on user
		})

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
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
		log.Info().Msgf("created user: %v", newUser)
		w.Write([]byte("success!")) // change this
	}
}
