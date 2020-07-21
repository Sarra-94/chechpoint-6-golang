package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Email struct {
	Title    string
	Subject  string
	To       string
	Message  string
	Greeting string
}

func createMessage() Email {
	message := Email{"Golang Tutorial", "Template Packages", "Radhouen Assakra", "How Are you ? \n I start learn golang and I discovered golang structure , types and functions and I'm no need to take a look about go template? Can you please send me a useful link to quick start .", "Thanks,best regards"}
	return message
}

func homeFunc(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	var fm = template.FuncMap{
		"message": createMessage,
	}
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("email.html"))
	if err := tpl.ExecuteTemplate(w, "email.html", createMessage()); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", homeFunc)
	http.ListenAndServe("localhost:1230", r)
}
