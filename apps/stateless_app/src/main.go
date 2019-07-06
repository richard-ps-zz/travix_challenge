package main
import (
	"log"
	"fmt"
    "net/http"
    _ "github.com/lib/pq"
    "database/sql"
    "os"
)

var (
  host     = os.Getenv("DATABASE_POSTGRESQL_SERVICE_HOST")
  port     = os.Getenv("DATABASE_POSTGRESQL_SERVICE_PORT")
  user     = "postgres"
  password = "abc"
  dbname   = "articlesdb"
)

func GetArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)	
	db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    fmt.Println("# Querying")
    rows, err := db.Query("SELECT * FROM articles")
    checkErr(err)

    for rows.Next() {
        var article string
        var data string
        var heading string
        var news_type string
        err = rows.Scan(&article, &data, &heading, &news_type)
        checkErr(err)
        fmt.Fprintf(w, "%3v | %8v | %6v | %6v\n", article, data, heading, news_type)
    }

}

func main() {
    http.HandleFunc("/articles", GetArticles)
    fmt.Printf("Listening on localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}