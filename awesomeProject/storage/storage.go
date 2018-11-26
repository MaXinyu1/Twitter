package storage

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"
)

const (
	userName = "root"
	password = ""
	ip = "127.0.0.1"
	port = "3306"
	dbName = "twitter"
)

var DB = &sql.DB{}

type User struct {
	UserName  string
	passWord  string
	Posts     Twitlist
	Following []string
}

type Twitte struct {
	username, time, content string
}

type Twitlist []Twitte

type TwitterPage struct {
	username   string
	UnFollowed []string
	Following  []string
	Posts      []string
}

// Sort Function needed these three Function
func (I Twitlist) Len() int {
	return len(I)
}
func (I Twitlist) Less(i, j int) bool {
	return I[i].time > I[j].time
}
func (I Twitlist) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}

func DBstart() {
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB,_ = sql.Open("mysql", path);
	_,err := sql.Open("mysql", path);
	if err != nil {
		fmt.Printf("connect mysql failed! [%s]", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		panic(err.Error())
		return
	}
}


func showFollow(username string) []string{ //good
	rows,_ := DB.Query("select toU from follow where fromU = ?", username)
	var following []string
	for rows.Next(){
		var who string
		rows.Scan(&who)
		following = append(following, who)
	}
	return following
}

func showUnfollow(username string) []string{ //good
	rows,_ := DB.Query("select username from user where username not in (select toU from follow where fromU = ?)", username)
	//db.Query("select username from user where username not in ?", follows)
	var unfollowing []string
	for rows.Next(){
		var unfollow string
		rows.Scan(&unfollow)
		//fmt.Println(unfollow)
		unfollowing = append(unfollowing, unfollow)
	}
	return unfollowing
}

func showPost(follows []string) Twitlist{ //good
	var twitte Twitte
	var twitlist Twitlist
	for i := range follows {
		name := follows[i]
		rows,_ := DB.Query("select * from twitte where username = ?", name)
		for rows.Next(){
			rows.Scan(&twitte.username, &twitte.content, &twitte.time)
			//fmt.Println(twitte.content)
			twitlist = append(twitlist, twitte)
		}
	}
	return twitlist
}

func Deletes(Following []string, username string) []string{
	var res []string
	for _,n := range Following{
		if n != username {
			res = append(res, n)
		}
	}
	return res
}

func getContent (Posts Twitlist) []string {
	var res []string
	var temp string
	for _,n := range Posts{
		temp = n.username + ":     >" + n.content
		//fmt.Println(temp)
		res = append(res, temp)
	}
	return res
}

func GetTwitterPage(username string) TwitterPage {
	DBstart()

	Following := showFollow(username)
	UnFollowed := showUnfollow(username)
	Posts := showPost(Following)

	sort.Sort(Posts)

	//transfer Posts to string form
	newPosts := getContent(Posts)

	// Remove the user itself from following list (just not shown in screen but in memory)
	Following = Deletes(Following, username)

	pg := TwitterPage{username: username, Following: Following, UnFollowed: UnFollowed, Posts: newPosts}
	//fmt.Println(pg.UnFollowed[0])
	DB.Close()
	return pg
}


