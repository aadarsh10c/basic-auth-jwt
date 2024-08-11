package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	// http.HandleFunc("/refresh", Refresh)
	// http.HandleFunc("/logout", Logout)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
