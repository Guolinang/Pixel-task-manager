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
	store          types.UserStore
	storeCharacter types.CharacterStore
}

func NewHandler(store types.UserStore, storeC types.CharacterStore) *Handler {
	return &Handler{store: store, storeCharacter: storeC}
}

func (h *Handler) RegisterRoute() {
	http.HandleFunc("/register", h.handleRegister)
	http.HandleFunc("/login", h.handleLogin)
}

func (j *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		j.handleNewAccount(w, r)
	case "PUT":
		j.handleUpdateAccount(w, r)
	default:
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("method is not supportedaaaa"))
	}

}

func (j *Handler) handleUpdateAccount(w http.ResponseWriter, r *http.Request) {

	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err.(validator.ValidationErrors)))
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("token not found"))
		return
	}
	userid, err := auf.ParseJWT([]byte(config.Env.JWTSecret), token)

	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("authentication failed: %v", err))
		return
	}

	u, err := j.store.GetUserById(userid)

	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	if u.Login != payload.Login {
		_, err := j.store.GetUserBylogin(payload.Login)
		if err == nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with login %s already exists", payload.Login))
			return
		}
	}

	hashed, err := auf.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = j.store.UpdateUser(types.User{
		Login:    payload.Login,
		Password: hashed,
		ID:       userid,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, "Good")

}

func (j *Handler) handleNewAccount(w http.ResponseWriter, r *http.Request) {

	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

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

	user, err := j.store.GetUserBylogin(payload.Login)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = j.storeCharacter.CreateCharacter(&types.Character{
		UserID: user.ID,
		Level:  1,
		Exp:    0,
		MaxExp: 100,
		Hp:     100,
		MaxHp:  100,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
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
	log.Print(token)
	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})

}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkQXQiOjE3NDI0MDY0MDEsInVzZXJpZCI6IjcifQ.I-CgV5vaIBMyo0ceH11Ep5CQQwtsz1avjMMTKY3x7oM
