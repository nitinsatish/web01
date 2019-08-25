package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseGlob("*.gohtml"))

func aboutme(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is Nitin")
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func you(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your details:\n %s", r.Header["User-Agent"])
}

func welcome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "index.gohtml", nil)
	} else {
		r.ParseForm()
		fmt.Fprintf(w, "Welcome %s", r.Form.Get("firstname"))
	}
}
func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about/", aboutme)
	http.HandleFunc("/you", you)
	http.HandleFunc("/welcome", welcome)

	http.ListenAndServe(":8000", nil)
}
