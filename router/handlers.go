package router

import (
	"net/http"
	"awesomeProject/config"

)

func Login(w http.ResponseWriter, r *http.Request) {
	 if r.Method == http.MethodPost {
		err := validUser(w, r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/secret", http.StatusSeeOther)
		return
	}
	config.Tpl.ExecuteTemplate(w, "login.html", nil)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := insertDb(r)
		if err != nil {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	config.Tpl.ExecuteTemplate(w, "signup.html", nil)
}

func Secret(w http.ResponseWriter, r *http.Request) {

	response, err := checkCookie(r, "session_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther	)
	}

	config.Tpl.ExecuteTemplate(w, "makeorder.html", response)
}
