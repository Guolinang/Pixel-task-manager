package tasks

import (
	"fmt"
	"net/http"
	"server/types"
	"server/utils"
)

type Handler struct {
	store types.TaskStore
}

func NewHandler(store types.TaskStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoute() {
	http.HandleFunc("/tasks", h.handleGetTasks)
}

func (h *Handler) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("method is not supported"))
		return
	}
	// var payload types.GetTasksPayload
	// if err := utils.ParseJSON(r, &payload); err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, err)
	// 	return
	// }
	// if err := utils.Validate.Struct(payload); err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err.(validator.ValidationErrors)))
	// 	return
	// }

	token := r.Header.Get("Authorization")
	if token == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("token not found"))
		return
	}

	// userid, err := auf.ParseJWT([]byte(config.Env.JWTSecret), token)

	// if err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("authentication failed: %v", err))
	// 	return
	// }
	userid := 2
	tasks, err := h.store.GetUserTasks(int(userid))

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, tasks)

}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkQXQiOjE3NDE5ODAzMDAsInVzZXJpZCI6IjIifQ.IrPF0zFc_KQwTooRSZhZvBpD17z9oikI_aQfLMrDYTA
