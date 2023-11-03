package main

import (
	"bytes"
	"compress/zlib"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
	"io"
	"log"
)

func archive(input []byte, mode byte) ([]byte, error) {
	var buffer bytes.Buffer
	if mode == 0 {
		writer := zlib.NewWriter(&buffer)
		writer.Write(input)
		writer.Close()
		return []byte(buffer.Bytes()), nil

	} else if mode == 1 {
		reader, err := zlib.NewReader(bytes.NewBuffer(input))
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(&buffer, reader)
		out := []byte(buffer.Bytes())
		return out, nil
	} else {
		msg := fmt.Sprintf("invalid operating mode specified: %d", mode)
		return input, errors.New(msg)
	}
}

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
	content := "hello world hello world"
	compress, err := archive([]byte(content), 0)
	if err != nil {
		log.Fatal(err)
	}
	deflate, err := archive(compress, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("input: ", []byte(content))
	fmt.Println("compressed: ", compress)
	fmt.Println("uncompressed", deflate)

}
