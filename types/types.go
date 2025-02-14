package types

import "time"

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
	// 	GetUserById(int) (*User, error)
	CreateUser(User) error
}

type Task struct {
	UserID       int       `json:"userID"`
	TaskID       int       `json:"taskID"`
	IsImportant  bool      `json:"isImportant"`
	TaskName     string    `json:"taskName"`
	Difficulty   int       `json:"difficulty"`
	SDescription string    `json:"sDescription"`
	Type         string    `json:"type"`
	Stats        string    `json:"stats"`
	Deadline     time.Time `json:"deadline"`
	Repeat       string    `json:"repeat"`
	Subtask      string    `json:"subtask"`
	FDescription string    `json:"fDescription"`
	Done         bool      `json:"done"`
}

type TaskStore interface {
	GetUserTasks(int) ([]Task, error)
	CreateTask(*Task) error
}

type GetTasksPayload struct {
	UserID       int       `json:"userID" validate:"required"`
	TaskID       int       `json:"taskID" validate:"required"`
	IsImportant  bool      `json:"isImportant" validate:"required"`
	TaskName     string    `json:"taskName" validate:"required"`
	Difficulty   int       `json:"difficulty" validate:"required"`
	SDescription string    `json:"sDescription" validate:"required"`
	Type         string    `json:"type" validate:"required"`
	Stats        string    `json:"stats" validate:"required"`
	Deadline     time.Time `json:"deadline" validate:"required"`
	Repeat       string    `json:"repeat" validate:"required"`
	Subtask      string    `json:"subtask" validate:"required"`
	FDescription string    `json:"fDescription" validate:"required"`
	Done         bool      `json:"done" validate:"required"`
}

type UserClaims struct {
	UserId int `json:"userid"`
}
