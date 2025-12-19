package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sample-go/repo"
	"sample-go/util"
)

type ReqCreateUser struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateUser

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		util.SentError(w, http.StatusBadRequest, "Please provide valid JSON data")
		return
	}

	usr, err := h.userRepo.Create(repo.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		IsShopOwner: req.IsShopOwner,
	})
	if err != nil {
		util.SentError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	util.SentData(w, http.StatusCreated, usr)

}
