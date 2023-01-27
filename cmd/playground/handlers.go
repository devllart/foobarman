package main

import (
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	JsonRun()
}
