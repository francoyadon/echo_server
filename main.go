package main

import (
    "net/http"
    "log"
)

const DefaultPort = "8089";



// EchoHandler echos back the request as a response
func EchoHandler(writer http.ResponseWriter, request *http.Request) {

    log.Println("Echoing back request made to " + request.URL.Path + " to client (" + request.RemoteAddr + ")")

    writer.Header().Set("Access-Control-Allow-Origin", "*")

    // allow pre-flight headers
    writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")

    request.Write(writer)
}

func main() {

    log.Println("starting server, listening on port " + DefaultPort)

    http.HandleFunc("/", EchoHandler)
    http.ListenAndServe(":" + DefaultPort, nil)
}
