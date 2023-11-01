package main

import (
	"bytes"
	"compress/zlib"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
	"io"
	"log"
)

// TODO: organize code into functions

func main() {
	var srcFile string
	var srcFileTags string
	var mode string
	var dbFile string
	var helpFlag string
	var versionNo string = "0.0.1"

	flag.StringVar(&mode, "m", "input", "Use chst in 'input' or 'output' mode")
	flag.StringVar(&srcFile, "f", "sampledoc.pdf", "Specify input-/output file for chst to target")
	flag.StringVar(&srcFileTags, "t", "", "Tags associated with the document.\nSeparate multiple tags by comma ','")
	flag.StringVar(&dbFile, "db", "chst.db", "OPTIONAL\nDatabase file for chst to use")
	flag.StringVar(&helpFlag, "h", "", "This help menu")
	flag.StringVar(&versionNo, "v", "", "OPTIONAL\nDatabase file for chst to use")
	flag.Parse()

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
