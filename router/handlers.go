package router

import (
	"net/http"
	"awesomeProject/config"
)



func Login(w http.ResponseWriter, r *http.Request) {
	/*if r.Method == http.MethodPost {
		err := validUser(r)
		if err != nil {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			http.Redirect(w, r, "/signup", http.StatusSeeOther)
			return
		}

	}*/
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

