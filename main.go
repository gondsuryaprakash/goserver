package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParseForm err : %v", err)
		return
	}
	fmt.Fprint(w, "Post Form is Successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprint(w, "Name = %s\n", name)
	fmt.Fprint(w, "Address = %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("here 24")
	if r.URL.Path != "/hello" {
		fmt.Print("here 25")
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		fmt.Print("here 30	")
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server Started ")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
