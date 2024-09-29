package api

import (
	"encoding/json"
	"fmt"
	"http-server/models"
	"http-server/service"
	"net/http"
)

func CreateSession(w http.ResponseWriter, r *http.Request) {
	var sessionRequest models.SessionRequest
	if err := json.NewDecoder(r.Body).Decode(&sessionRequest); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	user, validated := service.ValidateSessionRequest(&sessionRequest)
	if !validated {
		http.Error(w, "Invalid Credentials", http.StatusNotFound)
	}

	fmt.Printf("%+v\n", user)
	w.Write([]byte("success!"))
}
