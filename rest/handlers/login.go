package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sample-go/config"
	"sample-go/database"
	"sample-go/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid JSON data", http.StatusBadRequest)
		return
	}

	usr := database.Find(reqLogin.Email, reqLogin.Password)
	if usr == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	cnf := config.GetConfig()

	accessToken, err := util.CreateJwt(cnf.JWTSecret, util.Payload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})
	if err != nil {
		http.Error(w, "Error while creating JWT token", http.StatusInternalServerError)
		return
	}

	util.SentData(w, accessToken, http.StatusCreated)

}
