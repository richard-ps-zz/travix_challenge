package main
import (
	"log"
	"fmt"
    "net/http"
)

func GetArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}

func main() {
    http.HandleFunc("/articles", GetArticles)
    fmt.Printf("Listening on localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}