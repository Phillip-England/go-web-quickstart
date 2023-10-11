package main

import (
	"fmt"
	"net/http"
	"web-quickstart/pkg/PageBuilder"

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

    // homepage
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html")
        b := PageBuilder.New("Home Page")
        b.Add(`<p>Yooo</p>`)
        w.Write(b.GetBasePageAsBytes())
    })

    // serving
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }

}

