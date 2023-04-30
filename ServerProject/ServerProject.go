package main

import (
	Cnf "ServerProject/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var configuration Cnf.Configuration

func init() {
	println()
	file, err := os.Open("cfg/ServerProject.cfg")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	if err != nil {
		log.Print(err.Error())
	}
	decoder := json.NewDecoder(file)
	errd := decoder.Decode(&configuration)
	if errd != nil {
		fmt.Println("error:", errd)
	}
	fmt.Println("config loaded")
}

func main() {

	Process := func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		//case "GET":
		//	fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		case "POST":
			var p Cnf.Person

			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			contentType := r.Header.Get("Content-Type")
			if contentType != "application/json" {
				fmt.Println("Post not Jason")
			}

		default:
			fmt.Fprintf(w, "Loading!!! \n")
			fmt.Fprintf(w, "Sorry, POST methods are supported.")
		}
	}
	Denied := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, "Access denid!")
		if err != nil {
			return
		}
		println(err.Error())

	}
	http.HandleFunc("/", Denied)
	http.HandleFunc("/"+configuration.Served, Process)
	if err := http.ListenAndServe(":"+configuration.Port, nil); err != nil {
		log.Fatal(err)
	}
}
