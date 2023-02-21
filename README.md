# Client-Server-Proxy using Go

This repository contains a web server built in Go that can accept HTTP requests and return response data from locally stored files to clients. The server is designed to handle concurrent requests by creating a Go routine for each new client request.

The server is easy to set up and use, with clear documentation and a simple API for serving files. It is built with efficiency and performance in mind, making it an ideal choice for high-traffic web applications.

### go.mod

This files contains all the module paths and references

### main.go

The main.go is the program which runs the primary web server, to which all the clients and the proxy will be connecting. In the main.go, we have the
main():- This function first, creates and a new multiplexer in the server, which is used to handle routes for the connections. Then we define the custom properties of the server. Then we start listening on server, started on the port, accepted from CLI arguments. Then we accept connections to the server, and serve the connections on the route, based on the request. We are also using semaphores here to limit the number of concurrent connections to server.
func `requesthandler()`:- In this function, we check for the different request methods. We consider, two method, **GET** and **POST**. For every other method we are returning an error. Now inside the section for the **GET** method, we extract the URL and we parse the URL to obtain the file name/filepath, as requested. If the path is the root, we serve the http FileServer. If the request specifies a particular file, then we first check if the file is available. If the file is available, then we check the content type of the file, from the list of predefined content types. If all the conditions are satisfied, then we subsequently server the files. In the event the content type is not available, we return and error. Also if the file is not available , we give an error for the same. Similar process is followed for the POST process. The requesthandler function also host our todolist. It is handled in the following way. When the connection reaches the todolist.html, it will parse all the data, passed through the connection, then we serve the html in to the web browser. In todohtml, we have implemented an option to add a new todo. In the browser, there is a add button, for adding new todo to the list. On clicking the add button, we sent a POST request to the server, in the server we will add a new todo to the todo list and serve back the html to the browser.

### proxy.go

In the proxy.go, we implement the proxy server, which takes a connection from the client and connects to server. In the main function, we define the multiplexer for the proxy server for define the router for serving the incoming connections. We defined the properties of the proxy server. Then we accept the connections from the client and pass the connections to the web server.
In the handler function, we first connect to the web server, through the defined url. Then in a variable, we use the NewStringHostReverseProxy function map the incoming connection to the web server and subsequently send the responses back from the web server to the client.

### Instructions on how to run the server and proxy

- First go into each individual folder (server and proxy)
- Then run the command “make run” for each file separately(we have to run the server first then the proxy, obviously!!). The **Makefile** contains the instructions/command to build and run the docker container.

### Hosted on AWS EC2 instance

Addresses:

- Main Server: [http://52.200.235.170:8000/](http://52.200.235.170:8000/) (Deprecated)
- Proxy Server: [http://52.45.199.86:8080/](http://52.45.199.86:8080/) (Deprecated)

> Note: Assigned elastic IP to both the AWS EC2 instances, however, since we are unsure whether billing will be added, we decided to temporarly disable the instances until demonstration.
