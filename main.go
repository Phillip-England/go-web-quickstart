package main

import (
	"fmt"
	"net/http"
	components "web-quickstart/pkg/components/helloWorld"
	"web-quickstart/pkg/pageBuilder"

	"github.com/joho/godotenv"
)

//===============================================================
// MAIN
//===============================================================

func main() {

    // loading .env
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Println("Error loading .env file")
    }

    // Create a file server for the ./public directory and server static files from /public
    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/public/", http.StripPrefix("/public/", fs))

    // homepage
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html")
        b := pageBuilder.New("Home Page")
        b.Add(components.HelloWorld())
        w.Write(b.GetBasePageAsBytes())
    })

    // serving
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }

}

