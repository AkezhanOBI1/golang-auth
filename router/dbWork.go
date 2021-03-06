package router

import (
	"net/http"
	"errors"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"awesomeProject/config"
	"fmt"
	"time"
	"github.com/go-redis/redis"
)

type User struct {
	Email string `json:"email"`
	Name string `json:"name"`
	LastName string `json:"lastName"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}


func insertDb(r *http.Request) error {
	user := User{}
	user.Email = r.FormValue("email")
	user.Name = r.FormValue("name")
	user.LastName = r.FormValue("surname")
	user.Password = r.FormValue("password")
	user.ConfirmPassword = r.FormValue("confirmPassword")

	if user.Password != user.ConfirmPassword {
		return errors.New("405 Password do not match")
	}

	bytePass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 15)
	if err != nil {
		panic("Encrypt error")
	}
	passDb := hex.EncodeToString(bytePass)
	if err  := config.Db.Ping(); err != nil {
		panic("Pingin error")
	}

	fmt.Println("DATABASE is ok before inserting")
	go func() {
		sqlStatement := "INSERT INTO user_info (user_email, user_name, user_last_name, user_password) VALUES ($1, $2, $3, $4)"
		_, err = config.Db.Query(sqlStatement, user.Email, user.Name, user.LastName, passDb)
		if err != nil {
			panic("Error inserting")
		}
		fmt.Println("Inserted into db")
	}()

	pong, err := config.Cache.Ping().Result()
	fmt.Println(pong, err)

	go func() {
		 err := config.Cache.Set(user.Email, passDb, time.Hour).Err()
		if err != nil {
			panic("REDIS ERROR")
		}
		fmt.Println("Inserted into Redis")
	}()

	return nil
}


func validUser(w http.ResponseWriter, r *http.Request) error {
	userEmail := r.FormValue("email")
	userPassword := r.FormValue("password")

  	// From Reddis
	dbPass, err := config.Cache.Get(userEmail).Result()
	if err == redis.Nil {
		if err := config.Db.Ping(); err != nil {
			panic(err)
		}
		fmt.Println("You connected to your database.")

		rows, err := config.Db.Query("SELECT user_password FROM user_info WHERE user_email = $1", userEmail)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&dbPass)
		}
		unhash, err := hex.DecodeString(dbPass)
		if err != nil {
			panic("Decoing error")
		}

		err = bcrypt.CompareHashAndPassword(unhash, []byte(userPassword))
		if err != nil {
			panic("Passwords do not match")
		}

	}else if err != nil {
		panic("Redis getting error")
	}

	unhash, err := hex.DecodeString(dbPass)
	if err != nil {
		panic("Decoing error")
	}

	err = bcrypt.CompareHashAndPassword(unhash, []byte(userPassword))
	if err != nil {
		return errors.New("Password do not mathc 401")
	}

	// set Cookies

	err = setCookie(w, userEmail)
	if err != nil {
		return err
	}
	return nil
	/*
	// In PostGress

*/
}

