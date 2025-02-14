package main

import (
	"log"
	"server/api"
)

func main() {
	s := api.NewServer(":8000")
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}

}
