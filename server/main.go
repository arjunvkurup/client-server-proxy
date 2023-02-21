package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"golang.org/x/sync/semaphore"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

var data PageData

// requestHandler handles the request
func requestHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		reqPath := req.URL.Path
		pathSlices := strings.Split(reqPath, "/")
		fileName := pathSlices[len(pathSlices)-1]
		if reqPath == "/" {
			http.ServeFile(res, req, req.URL.Path[1:])
		} else if reqPath == "/static/todolist.html" {
			var buf bytes.Buffer

			err := tmpl.Execute(&buf, data)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "text/html; charset=UTF-8")
			buf.WriteTo(res)
		} else {
			contentType := http.DetectContentType([]byte(fileName))
			req.Header = make(http.Header)
			if contentType == "text/html; charset=utf-8" || contentType == "text/plain; charset=utf-8" || contentType == "text/css; charset=utf-8" || contentType == "text/javascript; charset=utf-8" || contentType == "image/jpeg; charset=utf-8" || contentType == "image/gif; charset=utf-8" || contentType == "application/json; charset=utf-8" {
				fmt.Println("Content Type", contentType)
				req.Header.Set("Content-Type", contentType)
				http.ServeFile(res, req, req.URL.Path[1:])
			} else {
				res.WriteHeader(http.StatusBadRequest)
				res.Write([]byte("400 - Bad Request"))
				return
			}

		}
	} else if req.Method == "POST" {
		reqPath := req.URL.Path
		if reqPath == "/static/todolist.html" { // implementing todo list
			formData := Todo{Item: req.FormValue("todolist"), Done: false}
			data.Todos = append(data.Todos, formData)
			var buf bytes.Buffer

			err := tmpl.Execute(&buf, data)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "text/html; charset=UTF-8")
			buf.WriteTo(res)
		}
	} else {
		fmt.Print("501")
		http.Error(res, "(501): Not Implemented", http.StatusNotImplemented)
	}

}

func main() {
	// todo list implementation
	data = PageData{
		Title: "TODO List",
		Todos: []Todo{
			{Item: "Install GO", Done: true},
			{Item: "Learn GO", Done: false},
		},
	}
	tmpl = template.Must(template.ParseFiles("static/todolist.html"))

	sem := semaphore.NewWeighted(10) // 10 is the max number of concurrent requests, hence being used

	router := http.NewServeMux()

	router.HandleFunc("/", requestHandler)

	// create a new http server
	server := http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	PORT := ":" + arguments[1]

	// create a new listener
	listener, err := net.Listen("tcp4", PORT)
	if err != nil {
		log.Fatal(err)
		return
	}

	// start the server
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// acquire a new semaphore
		sem.Acquire(context.Background(), 1)

		// start a new go routine
		go func() {
			defer sem.Release(1)
			defer conn.Close()
			server.Serve(listener)
		}()
	}
}
