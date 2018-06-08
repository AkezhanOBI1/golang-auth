package router

import (
	"net/http"
	"github.com/satori/go.uuid"
	"awesomeProject/config"
	"time"
	"errors"
	"github.com/go-redis/redis"
)

func setCookie(w http.ResponseWriter, id string) error{
	sessionToken, err := uuid.NewV4()
	if err != nil {
		panic("Session Token")
	}
	err = config.Cache.Set(sessionToken.String(), id, time.Hour).Err()
	if err != nil {
		// If there is an error in setting the cache, return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return errors.New("Error Setting cookie")
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken.String(),
		Expires: time.Now().Add(time.Hour),
	})

	return  nil
}



func checkCookie(r *http.Request, id string) (string, error){
	c, err := r.Cookie(id)
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return  "No Cookie No Party", errors.New("No Cookie")
		}
		// For any other type of error, return a bad request status
		return "Error", errors.New("Error occures")
	}
	sessionToken := c.Value
	response, err := config.Cache.Get(sessionToken).Result()
	if err == redis.Nil {
		return "No Cookie No Party", errors.New("No Cookie")
	}else if err != nil {
		return "Error", errors.New("Error occures")
	}
	return response, nil
}





