package main

import (
	"time"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

const (
	baseURL     = "http://localhost:8080/"
	shortURLLen = 10
)

var (
	shortURLMap   = make(map[string]string)
	inverseURLMap = make(map[string]string)
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		longURL := r.FormValue("longurl")
		if len(longURL) > 1000 { // limit the length of the long URL
			fmt.Fprintln(w, "Error: long URL is too long")
			return
		}
		shortURL := generateShortURL(longURL)
		shortURLMap[shortURL] = longURL
		inverseURLMap[longURL] = shortURL
		fmt.Fprintf(w, "Short URL: %s", baseURL+shortURL)
		return
	} else if r.Method == "GET" {
		fmt.Fprintf(w, `<html><body><form method="POST">
		Long URL: <input type="text" name="longurl">
		<input type="submit" value="Shorten">
		</form></body></html>`)
	} else {
		fmt.Println(errors.New("Wrong Request. Only supports POST and GET").Error())
		return
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := strings.TrimPrefix(r.URL.Path, "/")
	longURL, ok := shortURLMap[shortURL]
	fmt.Println(longURL)
	if !ok {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, longURL, http.StatusSeeOther)
}

func generateShortURL(longURL string) string {
	shortURL, ok := inverseURLMap[longURL]
	if ok {
		return shortURL
	}
	rand.Seed(time.Now().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var hashStr string
	for len(hashStr) < shortURLLen {
		hashStr += string(chars[rand.Intn(len(chars))])
	}
	return hashStr
}

func main() {
	http.HandleFunc("/home/", homeHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Listening on http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
