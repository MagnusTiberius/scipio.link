package main

import (
    "fmt"
    "log"
    "net/http"
	"html/template"
)

type Student struct {
	Name       string
	College    string
	RollNumber int
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
    student := Student{
        Name:       "GB",
        College:    "GolangBlogs",
        RollNumber: 1,
    }
    parsedTemplate, _ := template.ParseFiles("templates/demo1.html")
    err := parsedTemplate.Execute(w, student)
    if err != nil {
        log.Println("Error executing template :", err)
        return
    }
}

func main() {
    //http.HandleFunc("/", handler)
	http.HandleFunc("/", renderTemplate)
    log.Fatal(http.ListenAndServe(":80", nil))
}