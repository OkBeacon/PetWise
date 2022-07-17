package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main()  {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	fmt.Println("Listen & Serve http://localhost:9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil{
		log.Fatalf("ListenAndServe: %s", err)
	}
}

func index(w http.ResponseWriter, r *http.Request){
	//_, _ = io.WriteString(w, "Hello World!")
	_ = tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func processor(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fname := r.FormValue("firster")
	lname := r.FormValue("laster")

	d := struct {
		First string
		Last string
	}{
		First: fname,
		Last: lname,
	}

	_ = tpl.ExecuteTemplate(w, "processor.gohtml", d)
}
