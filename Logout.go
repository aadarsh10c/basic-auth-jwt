package main

import (
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now().Add(2 * time.Second),
	})
}
