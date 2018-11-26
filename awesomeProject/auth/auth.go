package auth

import (
	"awesomeProject/cookie"
	"awesomeProject/storage"
	"fmt"
	"html/template"
	"net/http"
)

//var db = &sql.DB{}
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t,_ := template.ParseFiles("show/login1.html")
		t.Execute(w, nil)
	} else {
		redirectAddress := ""
		r.ParseForm()
		username := r.Form["username"][0]
		password := r.Form["password"][0]
		method := r.Form["lr"][0]
		if method == "login" {
			isRight := loginCheck(username, password)
			if isRight == true {
				cookie.SetSession(username, w)
				redirectAddress = "personalPage"
			} else {
				redirectAddress = "wrongPassword"
			}
			http.Redirect(w, r, redirectAddress, http.StatusFound)

		}else {
			isRight := registerCheck(username, password)
			if isRight == true {
				redirectAddress = "registerSuccess"
			} else {
				redirectAddress = "registerFail"
			}
			http.Redirect(w, r, redirectAddress, http.StatusFound)
		}
	}
}

func loginCheck(Username string, Password string) (isTrue bool) {
	storage.DBstart()
	rows, err := storage.DB.Query("SELECT password FROM user where username = ?", Username);
	if err != nil {
		fmt.Println(err)
	} else {
		for rows.Next() {
			password := ""
			rows.Scan(&password);
			if Password == password {
				isTrue = true
				//fmt.Println("Right password")
			}else {
				isTrue = false
				//fmt.Println("Wrong password")
			}
		}
	}
	storage.DB.Close()
	return isTrue
}


func WrongPassword (w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t,_ := template.ParseFiles("show/wrongPassword.html")
		t.Execute(w, nil)
	} else {
		http.Redirect(w, r, "login", http.StatusFound)
	}
}

func registerCheck(username string, password string) (isTrue bool) {
	storage.DBstart()
	rows,err := storage.DB.Query("select username from user")

	if err != nil {
		fmt.Println(err)
	} else {
		name := ""
		for rows.Next() {
			rows.Scan(&name)
			if name == username {
				isTrue = false
				return isTrue
			}
		}
	}

	//if the user doesn't exit then add it to the database
	_,err = storage.DB.Exec("insert into user (username, password) values (?, ?)", username, password)
	//add him/herself to friend list
	_,err = storage.DB.Exec("insert into follow (fromU, toU) values (?, ?)", username, username)
	//if err != nil {fmt.Println(err)} else {fmt.Println("insert succcess !")}
	storage.DB.Close()
	isTrue = true
	return isTrue
}

func RegisterSuccess (w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t,_ := template.ParseFiles("show/registerSuccess.html")
		t.Execute(w, nil)
	} else {
		http.Redirect(w, r, "login", http.StatusFound)
	}
}

func RegisterFail (w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t,_ := template.ParseFiles("show/registerFail.html")
		t.Execute(w, nil)
	} else {
		http.Redirect(w, r, "login", http.StatusFound)
	}
}
