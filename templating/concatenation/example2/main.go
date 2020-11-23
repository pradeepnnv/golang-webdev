package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(`Required Argument "Name" is missing.`)
	}
	name := "Pradeep"
	tpl := `<html>
	<body>
	<h1>` + name + `</h1></body></html>`

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file ", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))

}
