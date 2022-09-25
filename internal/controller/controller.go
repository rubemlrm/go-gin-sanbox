package api

import (
	"log"
	"os"
)

func testing() {
	log.Println(os.Getenv("TESTING_FLAG"))
}
