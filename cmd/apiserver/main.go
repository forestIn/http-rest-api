package main

import (
	"log"

	"github.com/forestin/http-rest-api/internal/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start; err != nil {
		log.Fatal()
	}

}
