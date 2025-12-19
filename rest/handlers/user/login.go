package user

import (
	"encoding/json"
	"net/http"
	"sample-go/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		util.SentError(w, http.StatusBadRequest, "Please provide valid JSON data")
		return
	}

	usr, err := h.userRepo.Find(req.Email, req.Password)
	if err != nil {
		util.SentError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	accessToken, err := util.CreateJwt(h.cnf.JWTSecret, util.Payload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})
	if err != nil {
		util.SentError(w, http.StatusInternalServerError, "Error while creating JWT token")
		return
	}

	util.SentData(w, http.StatusCreated, accessToken)

}
