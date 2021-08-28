package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Baraha/csv_converter.git/api"
	"github.com/ClickHouse/clickhouse-go"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var database *sql.DB

func main() {
	connect, err := sql.Open("clickhouse", "tcp://127.0.0.1:9000?debug=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}

	}
	r := router.New()
	r.POST("/load", api.LoadFile)
	fmt.Println("server is start!")
	log.Fatal(fasthttp.ListenAndServe(":8081", r.Handler))

}
