package api

import (
	"database/sql"
	"log"
	"net/http"
	"server/character"
	"server/service/tasks"
	"server/service/users"

	_ "github.com/lib/pq"
)

type Server struct {
	address string
}

func NewServer(address string) *Server {
	return &Server{address: address}
}

func NewDB(connString string) (*sql.DB, error) {
	Db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	return Db, nil

}

func (s *Server) Run() error {
	http.Handle("/", http.FileServer(http.Dir("frontend/dist")))

	db, err := NewDB("user=postgres dbname=taskmanager sslmode=disable password=123")

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB connected")
	profileStore := character.NewStore(db)

	userstore := users.NewStore(db)
	userHandler := users.NewHandler(userstore, profileStore)
	userHandler.RegisterRoute()

	taskStore := tasks.NewStore(db)
	tasksHandler := tasks.NewHandler(taskStore)
	tasksHandler.RegisterRoute()

	profileHandler := character.NewHandler(profileStore)
	profileHandler.RegisterRoute()

	log.Print("Listening on port ", s.address)
	return http.ListenAndServe(s.address, nil)
}
