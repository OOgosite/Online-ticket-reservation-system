package main

import (
    "appengine"
    "appengine/user"
    "fmt"
    "http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
        url, err := user.LoginURL(c, r.URL.String())
        if err != nil {
            http.Error(w, err.String(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Location", url)
        w.WriteHeader(http.StatusFound)
        return
    }
	fmt.Fprintf(w, "<center>")
    fmt.Fprintf(w, "<h1>Hello, %v!</h1>", u)
	fmt.Fprintf(w, "<br><br><br><center>This is Test web application on Go + App Engine</center>")
	fmt.Fprintf(w, "<br><center>Develop by Chatchai</center>")
	fmt.Fprintf(w, "</center>")
}
