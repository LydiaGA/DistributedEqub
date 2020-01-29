package main

import (
	"equb1/DistributedEqub/config"
	db2 "equb1/DistributedEqub/db"
	"fmt"
	"html/template"
	"strconv"

	// "net"
	"net/http"
	//"strconv"
	"io/ioutil"
)

type Data struct {
	Equb     db2.Equb
	Month    string
	Me       db2.Member
	MyId     uint
	ServerId uint
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	db2.Migrate()
	db := db2.GetDatabase()
	defer db.Close()
	myId := uint(0)
	if len(db2.FindMe(db)) != 0 {
		myId = db2.FindMe(db)[0].MyId
	}

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			fmt.Println(err)
		}
		data := Data{
			MyId: myId,
		}
		fmt.Println(config.Me.ID)
		err2 := tmpl.Execute(w, data)
		if err2 != nil {
			fmt.Println(err2)
		}
	} else if r.Method == "POST" {
		r.ParseForm()
	}
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	equb := GetEqub()
	if r.Method == "GET" {
		//var body, _ = LoadFile("main.html")
		//fmt.Fprintf(w, body)
		tmpl, err := template.ParseFiles("main.html")
		if err != nil {
			fmt.Println(err)
		}
		data := Data{
			Equb:     equb,
			Month:    months[equb.CurrentMonth-1],
			Me:       config.Me,
			ServerId: equb.NextServerID - 1,
		}
		fmt.Println(config.Me.ID)
		err2 := tmpl.Execute(w, data)
		if err2 != nil {
			fmt.Println(err2)
		}
	} else if r.Method == "POST" {
		r.ParseForm()
	}

}

func JoinHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		var body, _ = LoadFile("join.html")
		fmt.Fprintf(w, body)

	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
		serverIp := r.FormValue("serverIp")
		serverPort := r.FormValue("serverPort")
		port := r.FormValue("myPort")
		name := r.FormValue("name")
		amount, err := strconv.Atoi(r.FormValue("amount"))
		StartClient(serverIp, serverPort, port, name, amount)
		http.Redirect(w, r, "/main", http.StatusSeeOther)
	}
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		var body, _ = LoadFile("create.html")
		fmt.Fprintf(w, body)
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
		equbName := r.FormValue("equbname")
		month, err := strconv.Atoi(r.FormValue("month"))
		port := r.FormValue("myPort")
		name := r.FormValue("name")
		amount, err := strconv.Atoi(r.FormValue("amount"))
		go StartServer(equbName, month, port, name, amount)
		http.Redirect(w, r, "/main", http.StatusSeeOther)
	}
}

func ResumeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		var body, _ = LoadFile("resume.html")
		fmt.Fprintf(w, body)
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
		serverIp := r.FormValue("serverIp")
		serverPort := r.FormValue("serverPort")
		Resume(serverIp, serverPort)
		http.Redirect(w, r, "/main", http.StatusSeeOther)
	}
}

func PayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		MakePayment()
		http.Redirect(w, r, "/main", http.StatusSeeOther)
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/main", http.StatusSeeOther)
	}
}

func ChangeMonthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ChangeMonth()
		http.Redirect(w, r, "/main", http.StatusSeeOther)
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/main", http.StatusSeeOther)
	}
}

func LoadFile(filename string) (string, error) {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil

}

func main() {
	db2.Migrate()
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/index", IndexHandler)
	http.HandleFunc("/main", MainHandler)
	http.HandleFunc("/create", CreateHandler)
	http.HandleFunc("/join", JoinHandler)
	http.HandleFunc("/resume", ResumeHandler)
	http.HandleFunc("/pay", PayHandler)
	http.HandleFunc("/change", ChangeMonthHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}

}
