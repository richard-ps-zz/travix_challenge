package main
import (
	"fmt"
    _ "github.com/lib/pq"
    "database/sql"
    "os"
    "io/ioutil"
    "strings"
)

var (
  host     = os.Getenv("DATABASE_POSTGRESQL_SERVICE_HOST")
  port     = os.Getenv("DATABASE_POSTGRESQL_SERVICE_PORT")
  user     = "postgres"
  password = "abc"
  dbname   = "articlesdb"
)

func main() {
    fmt.Println("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)
    dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)  
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    file, err := ioutil.ReadFile("create_table.sql")
    checkErr(err)

    stmt := string(file)
    _, err2 := db.Exec(stmt)
    checkErr(err2)

    file2, err := ioutil.ReadFile("articles.sql")
    checkErr(err)

    requests := strings.Split(string(file2), "*")

    for _, stmt2 := range requests {
        fmt.Println(stmt2)
        _, err3 := db.Exec(stmt2)
        checkErr(err3)
    }
}

func checkErr(err error) {
    if err != nil {
        fmt.Println(err.Error())
    }
}