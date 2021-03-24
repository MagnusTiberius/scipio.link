package main

import (
  "fmt"
  "html/template"
  "net/http"
  

  "log"

  "github.com/gorilla/pat"
  "github.com/markbates/goth"
  "github.com/markbates/goth/gothic"
  "github.com/markbates/goth/providers/google"
  "github.com/gorilla/sessions"
)

// https://www.loginradius.com/blog/async/google-authentication-with-golang-and-goth/


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

key := "Secret-session-key"  // Replace with your SESSION_SECRET or similar
  maxAge := 86400 * 30  // 30 days
  isProd := false       // Set to true when serving over https

  store := sessions.NewCookieStore([]byte(key))
  store.MaxAge(maxAge)
  store.Options.Path = "/"
  store.Options.HttpOnly = true   // HttpOnly should always be enabled
  store.Options.Secure = isProd

  gothic.Store = store

  goth.UseProviders(
    google.New("our-google-client-id", "our-google-client-secret", "http://localhost:8080/auth/google/callback", "email", "profile"),
  )

  p := pat.New()
  p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

    user, err := gothic.CompleteUserAuth(res, req)
    if err != nil {
      fmt.Fprintln(res, err)
      return
    }
    t, _ := template.ParseFiles("templates/success.html")
    t.Execute(res, user)
  })

  p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
    gothic.BeginAuthHandler(res, req)
  })

  p.Get("/privacy/policy", func(res http.ResponseWriter, req *http.Request) {
    t, _ := template.ParseFiles("templates/policy.html")
    t.Execute(res, false)
  })

  p.Get("/", func(res http.ResponseWriter, req *http.Request) {
    t, _ := template.ParseFiles("templates/index.html")
    t.Execute(res, false)
  })


    //http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}