package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

// Handler function for the proxy server
func handlerFunction(res http.ResponseWriter, req *http.Request) {

	// need to set the ip of the destination server
	dst_rt_url, err := url.Parse("http://localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	// create a reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(dst_rt_url)
	if req.Method == http.MethodGet {
		proxy.ServeHTTP(res, req)
		return
	} else {
		res.WriteHeader(http.StatusNotImplemented)
		res.Write([]byte("501 - Not Implemented"))
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port!")
		return
	}

	PORT := ":" + arguments[1]

	router := http.NewServeMux()

	router.HandleFunc("/", handlerFunction) // establishing the route and the handler function

	proxy_server := http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// only handles tcp ipv4 protocols
	listener, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	for {
		_, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		proxy_server.Serve(listener)
	}
}
