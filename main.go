package main

import (
	"bytes"
	"compress/zlib"
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
	"io"
	"log"
)

// TODO: organize code into functions

func main() {
	// connect
	// TODO: switch to on disk database when file loading is implemented
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

	var cbuffer bytes.Buffer
	var rbuffer bytes.Buffer

	writer := zlib.NewWriter(&cbuffer)
	defer writer.Close()

	reader, err := zlib.NewReader(&cbuffer)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	//TODO: research how to properly use reader
	writer.Write([]byte(content))
	fmt.Println(cbuffer.String())

	io.Copy(&rbuffer, reader)
	fmt.Println(rbuffer.String())

}
