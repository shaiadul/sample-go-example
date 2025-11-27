package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sample-go/database"
	"sample-go/util"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid JSON data", http.StatusBadRequest)
		return
	}

	createdUser := newUser.Store()

	util.SentData(w, createdUser, http.StatusCreated)

}
