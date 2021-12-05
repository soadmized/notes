package main

import (
	"log"
	"net/http"
)

func main() {
	s := &service{repo: newStorage()}
	r := router(s)
	log.Fatal(http.ListenAndServe(":8000", r))
}
