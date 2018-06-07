package router

import (
	"net/http"
	"awesomeProject/config"
)



func Login(w http.ResponseWriter, r *http.Request) {
	config.Tpl.ExecuteTemplate(w, "login.html", nil)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		_, err := insertDb(r)
		if err != nil {
			http.Error(w, http.StatusText(406), http.StatusMethodNotAllowed)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}



	config.Tpl.ExecuteTemplate(w, "signup.html", nil)
}