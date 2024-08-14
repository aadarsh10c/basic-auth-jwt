package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/logout", Logout)
	fmt.Println("Listining at port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
