package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Server is running on port:", port)

	h1 := func(w http.ResponseWriter, r *http.Request) {
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppolasso"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}

		temp := template.Must(template.ParseFiles("index.html"))

		err := temp.Execute(w, films)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)

		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		temp := template.Must(template.ParseFiles("index.html"))

		err := temp.ExecuteTemplate(w, "film-list-element", Film{
			Title:    title,
			Director: director,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	http.ListenAndServe("127.0.0.1:"+port, nil)
}
