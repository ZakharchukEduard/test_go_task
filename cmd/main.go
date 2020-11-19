package main

import (
	"log"

	todo "github.com/ZakharchukEduard/test_go_task"
	"github.com/ZakharchukEduard/test_go_task/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
