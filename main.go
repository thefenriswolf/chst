package main

import (
	"bytes"
	"compress/zlib"
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
	"log"
)

func main() {
	// connect
	db, err := sql.Open("sqlite", ":memory:") // test with temp db
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	content := "hello world"
	fmt.Println(content)

	var buffer bytes.Buffer

	writer := zlib.NewWriter(&buffer)
	defer writer.Close()

	writer.Write([]byte(content))
	fmt.Println(buffer.String())
}
