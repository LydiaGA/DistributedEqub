package main

import (
	"bufio"
	db2 "equb2/DistributedEqub/db"
	"fmt"
	"strconv"

	// "net"
	"net/http"
	//"strconv"
	"io/ioutil"
)

type Data struct {
	Title string
}

func MainHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		var body, _ = LoadFile("main.html")
		fmt.Fprintf(w, body)
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
		serverIp := r.FormValue("equbname")
		serverPort := r.FormValue("month")
		port := r.FormValue("myPort")
		name := r.FormValue("name")
		amount, err := strconv.Atoi(r.FormValue("amount"))
		StartClient(serverIp, serverPort, port, name, amount)
	}
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		var body, _ = LoadFile("create.html")
		fmt.Fprintf(w, body)
		//tmpl, err := template.ParseFiles("create.html")
		//if err != nil{
		//	fmt.Println(err)
		//}
		//data := Data{
		//	Title: "HIIII",
		//}
		//tmpl.Execute(w, data)

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
	http.HandleFunc("/main", MainHandler)
	http.HandleFunc("/create", CreateHandler)
	http.HandleFunc("/join", JoinHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}

}
