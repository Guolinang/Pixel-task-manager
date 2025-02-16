package tasks

import (
	"fmt"
	"log"
	"net/http"
	"server/auf"
	"server/config"
	"server/types"
	"server/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store types.TaskStore
}

func NewHandler(store types.TaskStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoute() {
	http.HandleFunc("/tasks", h.handleTasks)
}

func (h *Handler) handleTasks(w http.ResponseWriter, r *http.Request) {

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
		h.handleGetTasks(w, r, userid)
	case "POST":
		h.handleCreateTask(w, r, userid)
	case "PUT":
		h.handleUpdateTask(w, r)
	case "DELETE":
		h.handleDeleteTask(w, r)
	default:
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("method is not supported"))
	}

}

func (h *Handler) handleUpdateTask(w http.ResponseWriter, r *http.Request) {

	var payload types.ManageTaskPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json %v", err))
		return
	}

	log.Print(payload)
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err.(validator.ValidationErrors)))
		return
	}
	err := h.store.UpdateTask(&types.Task{
		TaskID:       payload.TaskID,
		TaskName:     payload.TaskName,
		IsImportant:  payload.IsImportant,
		Difficulty:   payload.Difficulty,
		SDescription: payload.SDescription,
		Type:         payload.Type,
		Stats:        payload.Stats,
		Deadline:     payload.Deadline,
		Repeat:       payload.Repeat,
		Subtask:      payload.Subtask,
		FDescription: payload.FDescription,
		Done:         payload.Done,
	})

	if err != nil {

		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error updating task %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	var payload types.ManageTaskPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json %v", err))
		return
	}
	log.Print(payload)
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err.(validator.ValidationErrors)))
		return
	}
	err := h.store.DeleteTask(&types.Task{
		TaskID:       payload.TaskID,
		TaskName:     payload.TaskName,
		IsImportant:  payload.IsImportant,
		Difficulty:   payload.Difficulty,
		SDescription: payload.SDescription,
		Type:         payload.Type,
		Stats:        payload.Stats,
		Deadline:     payload.Deadline,
		Repeat:       payload.Repeat,
		Subtask:      payload.Subtask,
		FDescription: payload.FDescription,
		Done:         payload.Done,
	})

	if err != nil {

		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error delete task %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleGetTasks(w http.ResponseWriter, r *http.Request, userid int) {

	queryParams := r.URL.Query()
	dateStr := queryParams.Get("date")
	var date time.Time
	var err error
	if dateStr == "" {
		date = time.Now()
	} else {
		date, err = time.Parse(time.RFC3339, dateStr)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}

	tasks, err := h.store.GetSortedTasks(int(userid), utils.JsonDate(date))

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, tasks)

}

func (h *Handler) handleCreateTask(w http.ResponseWriter, r *http.Request, userid int) {

	var payload types.ManageTaskPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json %v", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err.(validator.ValidationErrors)))
		return
	}
	err := h.store.CreateTask(&types.Task{
		UserID:       int(userid),
		TaskName:     payload.TaskName,
		IsImportant:  payload.IsImportant,
		Difficulty:   payload.Difficulty,
		SDescription: payload.SDescription,
		Type:         payload.Type,
		Stats:        payload.Stats,
		Deadline:     time.Time(payload.Deadline),
		Repeat:       payload.Repeat,
		Subtask:      payload.Subtask,
		FDescription: payload.FDescription,
		Done:         payload.Done,
	})
	if err != nil {

		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error creating task %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)

}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkQXQiOjE3NDE5ODAzMDAsInVzZXJpZCI6IjIifQ.IrPF0zFc_KQwTooRSZhZvBpD17z9oikI_aQfLMrDYTA
