package router

import (
	"net/http"
	"awesomeProject/config"
	"github.com/go-redis/redis"
)



func Login(w http.ResponseWriter, r *http.Request) {
	 if r.Method == http.MethodPost {
		err := validUser(w, r)
		if err != nil {
		//	http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
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

	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value


	response, err := config.Cache.Get(sessionToken).Result()
	if err == redis.Nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	config.Tpl.ExecuteTemplate(w, "makeorder.html", response)
}