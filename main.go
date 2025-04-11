package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

func serveImage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println("method not allowed")
		return
	}

	imagePath := r.URL.Path[len("/image/"):]
	fullPath := filepath.Join("res", imagePath)

	http.ServeFile(w, r, fullPath)
}


func main() {
    homePageTmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
		return
	}


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
            err := homePageTmpl.Execute(w, nil)
            if err != nil {
                log.Fatal(err)
            }
            return
		}	
	})

	http.HandleFunc("/image/", serveImage)

	const port = 8080
	fmt.Printf("Server listening on http://localhost:%d\n", port)
	http.ListenAndServe(":" + strconv.Itoa(port), nil)
}


