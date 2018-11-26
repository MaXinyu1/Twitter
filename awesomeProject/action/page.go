package action

import (
	"awesomeProject/cookie"
	"awesomeProject/storage"
	"fmt"
	"net/http"
	"html/template"
)

func PersonalPage(w http.ResponseWriter, r *http.Request){
	username := cookie.GetUserName(r)
	cookie.SetSession(username, w)
	//fmt.Println(username)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("show/personalPage.html")
		pagecontent := storage.GetTwitterPage(username)
		err := t.Execute(w, pagecontent)
		if err != nil {
			fmt.Println(err)
		}
		//t.Execute(w, nil)
	} else {
		r.ParseForm()
		logout := r.Form.Get("logout")
		if logout == "logout" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		method := r.Form.Get("pg")
		switch method {
		case "Send Twitte":
			content := r.Form.Get("twitte")
			sendTwitte(username, content)
		case "follow":
			follow := r.Form.Get("follow")
			followUser(username, follow)
		case "unfollow":
			unfollow := r.Form.Get("unfollow")
			unfollowUser(username, unfollow)
		}
		http.Redirect(w, r, "/personalPage", http.StatusFound)
	}
}