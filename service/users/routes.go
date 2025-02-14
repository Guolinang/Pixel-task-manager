package users

import (
	"fmt"
	"log"
	"net/http"
	"server/auf"
	"server/config"
	"server/types"
	"server/utils"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoute() {
	http.HandleFunc("/register", h.handleRegister)
	http.HandleFunc("/login", h.handleLogin)
}

func (j *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {

		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("method is not supported"))
		return

	}

	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Print("register ", payload)

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err.(validator.ValidationErrors)))
		return
	}

	_, err := j.store.GetUserBylogin(payload.Login)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with login %s already exists", payload.Login))
		return
	}

	hashed, err := auf.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = j.store.CreateUser(types.User{
		Login:    payload.Login,
		Password: hashed,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return

	}
	utils.WriteJSON(w, http.StatusCreated, "Good")

}

func (j *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {

		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("method is not supported"))
		return

	}

	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err.(validator.ValidationErrors)))
		return
	}

	u, err := j.store.GetUserBylogin(payload.Login)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("wrong login or password"))
		return
	}

	if !auf.ComparePassword(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("wrong login or password"))
		return
	}

	secret := []byte(config.Env.JWTSecret)
	token, err := auf.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})

}
