package types

import (
	"server/utils"
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `jsob:"password"`
}

type LoginUserPayload struct {
	Login    string `json:"login" validate:"required"`
	Password string `jsob:"password" validate:"required"`
}

type RegisterUserPayload struct {
	Login    string `json:"login" validate:"required"`
	Password string `jsob:"password" validate:"required"`
}

type UserStore interface {
	GetUserBylogin(string) (*User, error)
	CreateUser(User) error
}

type Task struct {
	UserID       int       `json:"userID"`
	TaskID       int       `json:"id"`
	IsImportant  bool      `json:"important"`
	TaskName     string    `json:"name"`
	Difficulty   int       `json:"difficulty"`
	SDescription string    `json:"s_desc"`
	Type         string    `json:"type"`
	Stats        string    `json:"stat"`
	Deadline     time.Time `json:"urgency"`
	Repeat       string    `json:"repeat"`
	Subtask      string    `json:"subtask"`
	FDescription string    `json:"l_desc"`
	Done         bool      `json:"done"`
}

type GetTasksResponse struct {
	Unfinished []Task `json:"unfinished"`
	Important  []Task `json:"important"`
	Today      []Task `json:"today"`
}

type TaskStore interface {
	GetUserTasks(int) ([]Task, error)
	CreateTask(*Task) error
	GetSortedTasks(int, utils.JsonDate) (*GetTasksResponse, error)
	DeleteTask(*Task) error
	UpdateTask(*Task) error
}

type ManageTaskPayload struct {
	TaskID       int       `json:"id"`
	UserID       int       `json:"userID"`
	IsImportant  bool      `json:"important"`
	TaskName     string    `json:"name"`
	Difficulty   int       `json:"difficulty"`
	SDescription string    `json:"s_desc"`
	Type         string    `json:"type"`
	Stats        string    `json:"stat"`
	Deadline     time.Time `json:"urgency"`
	Repeat       string    `json:"repeat"`
	Subtask      string    `json:"subtask"`
	FDescription string    `json:"l_desc"`
	Done         bool      `json:"done"`
}

type UserClaims struct {
	UserId int `json:"userid"`
}
