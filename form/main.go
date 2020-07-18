package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Email string
	Subject string
	Message string
}

func main(){
	tmpl := template.Must(template.ParseFiles("views/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			err := tmpl.Execute(w, nil)
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		fmt.Println(r.Header)

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		log.Println(details)

		tmpl.Execute(w, struct {
			Success bool
		}{true})
	})
	address := ":8080"
	log.Printf("Listening on port %s\n", address)

	http.ListenAndServe(address, nil)
}