package character

import (
	"fmt"
	"net/http"
	"server/auf"
	"server/config"
	"server/types"
	"server/utils"
)

type Handeler struct {
	store types.CharacterStore
}

func NewHandler(store types.CharacterStore) *Handeler {
	return &Handeler{store: store}
}

func (h *Handeler) RegisterRoute() {
	http.HandleFunc("/character", h.handleProfile)
}

func (h *Handeler) handleProfile(w http.ResponseWriter, r *http.Request) {

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

	switch r.Method {
	case "GET":
		h.handleGetCharacter(w, userid)
	case "PUT":
		h.handleUpdateCharacter(w, r, userid)
	default:
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("method is not supported"))
	}

}

func (h *Handeler) handleGetCharacter(w http.ResponseWriter, userid int) {

	charecter, err := h.store.GetCharacter(userid)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, charecter)
}

func (h *Handeler) handleUpdateCharacter(w http.ResponseWriter, r *http.Request, userid int) {

	var payload types.ManageCharacterPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json %v", err))
	}
	err := h.store.UpdateCharacter(&types.Character{
		UserID: userid,
		Level:  payload.Level,
		Exp:    payload.Exp,
		MaxExp: payload.MaxExp,
		Hp:     payload.Hp,
		MaxHp:  payload.MaxHp,
		Str:    payload.Str,
		Int:    payload.Int,
		Char:   payload.Char,
		Wis:    payload.Wis,
		Cnst:   payload.Cnst,
		Head:   payload.Head,
		Face:   payload.Face,
		Body:   payload.Body,
		Dress:  payload.Dress,
		Other:  payload.Other,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)

}
