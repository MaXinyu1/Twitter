package main

import (
	"awesomeProject/action"
	"awesomeProject/auth"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/login",auth.Login)
	http.HandleFunc("/wrongPassword", auth.WrongPassword)
	http.HandleFunc("/registerSuccess", auth.RegisterSuccess)
	http.HandleFunc("/registerFail", auth.RegisterFail)
	http.HandleFunc("/personalPage", action.PersonalPage)


	er := http.ListenAndServe(":9090",nil)

	if er != nil {
		log.Fatal("ListenAndServer: ", er)
	}
}