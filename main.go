package main

import (
	f "fmt"
	"log"
	"net/http"
	"net/url"
)

func main()  {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	f.Println("Server is listening on port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	u, err := url.Parse(r.RequestURI)
	if err != nil {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if u.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}


	// if r.Url.Path != "/hello" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	f.Fprintf(w, "Hello World!")
}


func formHandler(w http.ResponseWriter, r *http.Request)  {
	if err := r.ParseForm(); err != nil {
		f.Fprintf(w, "ParseForm() err: %v", err)
		return	
	}

	f.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	name := r.FormValue("name")
	address := r.FormValue("address")
	phone := r.FormValue("phone")
	f.Fprintf(w, "Name = %s\n", name)
	f.Fprintf(w, "Address = %s\n", address)
	f.Fprintf(w, "Phone = %s\n", phone)



	// if r.Url.Path != "/form" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "form.html")
}