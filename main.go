package main

import (
	"net/http"
	"awesomeProject/router"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", router.Login)
	http.HandleFunc("/signup", router.Signup)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "/signup", http.StatusSeeOther)
}