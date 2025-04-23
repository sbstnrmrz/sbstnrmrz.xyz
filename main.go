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

func serveScript(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println("method not allowed")
		return
	}

	scriptPath := r.URL.Path[len("/script/"):]

	http.ServeFile(w, r, scriptPath)
}

func serveAsset(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println("method not allowed")
		return
	}

	scriptPath := r.URL.Path[len("assets/"):]

	http.ServeFile(w, r, scriptPath)
}

func main() {
    homePageTmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
		return
	}

    zigtrisPageTmpl, err := template.ParseFiles("zigtris.html")
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

	http.HandleFunc("/zigtris", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
            err := zigtrisPageTmpl.Execute(w, nil)
            if err != nil {
                log.Fatal(err)
            }
            return
		}	
	})

	http.HandleFunc("/image/", serveImage)
	http.HandleFunc("/script/", serveScript)
	http.HandleFunc("assets/", serveAsset)

	const port = 8080
	fmt.Printf("Server listening on http://localhost:%d\n", port)
	http.ListenAndServe(":" + strconv.Itoa(port), nil)
}
