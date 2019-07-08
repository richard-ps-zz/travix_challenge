package main
import (
	"log"
	"fmt"
    "net/http"
    _ "github.com/lib/pq"
    "database/sql"
    "os"
    "encoding/json"
)

var (
  host     = os.Getenv("DATABASE_POSTGRESQL_SERVICE_HOST")
  port     = os.Getenv("DATABASE_POSTGRESQL_SERVICE_PORT")
  user     = "postgres"
  password = "abc"
  dbname   = "articlesdb"
)

type article struct {
    Article string
    Date  string
    Heading string
    Type string
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	fmt.Printf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)	
	db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    fmt.Println("# Querying")
    rows, err := db.Query("SELECT * FROM articles")
    checkErr(err)

    var articles = []article{}

    for rows.Next() {
        var content string
        var date string
        var heading string
        var news_type string
        err = rows.Scan(&content, &date, &heading, &news_type)
        checkErr(err)
        article_struct := article{Article: content,
                            Date: date,
                            Heading: heading,
                            Type: news_type}

        articles = append(articles, article_struct)
    }

    b, err := json.Marshal(articles)
    checkErr(err)
    fmt.Fprintf(w, string(b))

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