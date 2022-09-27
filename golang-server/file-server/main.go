// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func validateFileType(filename string, fileType string) bool {
	var validFilename = regexp.MustCompile("^\\w+\\." + fileType)
	m := validFilename.FindStringSubmatch(filename)
	if m == nil {
		return false
	}
	return true
}

func getTxtFiles() []fs.FileInfo {
	files, err := ioutil.ReadDir(".")
	var txtFiles []string
	if err != nil {
		return nil
	}

	for _, file := range files {
		isText := validateFileType(file.Name(), "txt")
		if !isText {
			continue
		}
		txtFiles = append(txtFiles, file.Name())
	}
	return files
}

func getHtml(file fs.FileInfo) string {
	sz := len(file.Name())
	name := file.Name()[:sz-4]
	oneRow := fmt.Sprintf("<li><a href=\"/view/%s\">%s</a></li>", name, name)
	return oneRow
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	files := getTxtFiles()
	all := ""
	for _, file := range files {
		all += getHtml(file)
	}
	all = fmt.Sprintf("<div><ul>%s</ul></div>", all)
	_, err := fmt.Fprintln(w, all)
	if err != nil {
		return
	}
}

func main() {
	getTxtFiles()
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/list/", listHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
