package action

import (
	"awesomeProject/storage"
	"fmt"
	"time"
)


func sendTwitte(username string, content string) {
	storage.DBstart()

	t1 := time.Now().Year()
	t2 := time.Now().Month()
	t3 := time.Now().Day()
	t4 := time.Now().Hour()
	t5 := time.Now().Minute()
	t6 := time.Now().Second()
	t7 := time.Now().Nanosecond()
	t := time.Date(t1, t2, t3, t4, t5, t6, t7, time.Local)
	fmt.Println(t)
	storage.DB.Exec("insert into twitte (username, content, time) values (?, ?, ?)", username, content, t)
	//if err != nil { fmt.Println(err) } else { fmt.Println("insert twitte succcess !") }
	storage.DB.Close();
}

func followUser(username string, follow string) bool{ //good
	storage.DBstart()
	_,err := storage.DB.Query("insert into follow (fromU, toU) values (?, ?)", username, follow)
	storage.DB.Close()
	if err == nil {
		return true
	}
	return false
}

func unfollowUser(username string, unfollowname string) bool{ //good
	storage.DBstart()
	_,err := storage.DB.Query("delete from follow where fromU = ? and toU = ?", username, unfollowname)
	storage.DB.Close()
	if err == nil {
		return true
	}
	return false
}
